package git

import (
	"fmt"
	"sort"

	version "github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func dataSourceGitRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRepositoryRead,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  ".git",
			},

			"branch": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"commit_sha": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"tag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	path := d.Get("path").(string)

	repo, err := git.PlainOpen(path)
	if err != nil {
		return fmt.Errorf("Error opening repository in %s: %s", path, err)
	}

	ref, err := repo.Head()
	if err != nil {
		return fmt.Errorf("Error getting HEAD reference: %s", err)
	}

	refName := ref.Name()

	branch := ""
	if refName.IsBranch() {
		branch = refName.Short()
	}

	tag := ""
	tags, err := getTags(repo, ref)
	if err != nil {
		return fmt.Errorf("Error getting tags: %s", err)
	}

	if tags != nil {
		tag = getLatestTag(tags)
	}

	d.Set("branch", branch)
	d.Set("commit_sha", ref.Hash().String())
	d.Set("tag", tag)

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
