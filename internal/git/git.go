package git

import (
	"sort"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	version "github.com/hashicorp/go-version"
)

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
