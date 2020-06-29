package git

import (
	"fmt"

	"github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
	var refName plumbing.ReferenceName
	var ref *plumbing.Reference
	var err error

	if branch != "" {
		refName = plumbing.NewBranchReferenceName(branch)
	} else if tag != "" {
		refName = plumbing.NewTagReferenceName(tag)
	}

	if repoURL != "" {
		cloneOptions := git.CloneOptions{
			URL:           repoURL,
			Auth:          auth,
			ReferenceName: refName,
			SingleBranch:  true,
			Depth:         1,
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
		ref, err = repo.Reference(refName, false)
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
		ref, err = repo.Reference(refName, false)
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
