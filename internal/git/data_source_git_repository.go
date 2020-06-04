package git

import (
	"fmt"
	"sort"

	version "github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"gopkg.in/src-d/go-billy.v4/memfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func dataSourceGitRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGitRepositoryRead,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				DefaultFunc:  schema.EnvDefaultFunc("GIT_DIR", ".git"),
			},

			"url": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				ExactlyOneOf: []string{"path", "url"},
			},

			"branch": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"tag": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validation.StringIsNotEmpty,
				ConflictsWith: []string{"branch"},
			},

			"commit_sha": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGitRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	auth := m.Auth

	path := d.Get("path").(string)
	repoURL := d.Get("url").(string)
	branch := d.Get("branch").(string)
	tag := d.Get("tag").(string)

	var repo *git.Repository
	var ref *plumbing.Reference
	var err error

	if repoURL != "" {
		cloneOptions := git.CloneOptions{
			URL:          repoURL,
			Auth:         auth,
			SingleBranch: true,
			Depth:        1,
		}

		if branch != "" {
			cloneOptions.ReferenceName = plumbing.NewBranchReferenceName(branch)
		} else if tag != "" {
			cloneOptions.ReferenceName = plumbing.NewTagReferenceName(tag)
		}

		repo, err = git.Clone(memory.NewStorage(), memfs.New(), &cloneOptions)
		if err != nil {
			return fmt.Errorf("unable to clone repository: %s", err)
		}
	} else {
		repo, err = git.PlainOpen(path)
		if err != nil {
			return fmt.Errorf("unable to open repository: %s", err)
		}

		remote, _ := repo.Remote(git.DefaultRemoteName)
		if remote != nil {
			repoURL = remote.Config().URLs[0]
		}
	}

	if branch != "" {
		ref, err = repo.Reference(plumbing.NewBranchReferenceName(branch), false)
		if err != nil {
			return fmt.Errorf("unable to get branch: %s", err)
		}
	} else {
		ref, err = repo.Head()
		if err != nil {
			return fmt.Errorf("unable to get HEAD ref: %s", err)
		}

		if ref.Name().IsBranch() {
			branch = ref.Name().Short()
		}
	}

	if tag != "" {
		ref, err = repo.Reference(plumbing.NewTagReferenceName(tag), false)
		if err != nil {
			return fmt.Errorf("unable to get tag: %s", err)
		}
	} else {
		tags, err := getTags(repo, ref)
		if err != nil {
			return fmt.Errorf("unable to get tags: %s", err)
		}

		if tags != nil {
			tag = getLatestTag(tags)
		}
	}

	d.Set("branch", branch)
	d.Set("commit_sha", ref.Hash().String())
	d.Set("tag", tag)
	d.Set("url", repoURL)

	d.SetId(ref.Name().String())

	return nil
}

func getTags(repo *git.Repository, ref *plumbing.Reference) ([]string, error) {
	var tags []string

	tagRefs, err := repo.Tags()
	if err != nil {
		return nil, err
	}

	err = tagRefs.ForEach(func(tag *plumbing.Reference) error {
		if tag.Hash().String() == ref.Hash().String() {
			tags = append(tags, tag.Name().Short())
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func getLatestTag(tags []string) string {
	var versions []*version.Version

	for _, t := range tags {
		v, _ := version.NewVersion(t)
		if v != nil {
			versions = append(versions, v)
		}
	}

	sort.Sort(sort.Reverse(version.Collection(versions)))
	if len(versions) > 0 {
		return versions[0].Original()
	}

	sort.Strings(tags)
	return tags[0]
}
