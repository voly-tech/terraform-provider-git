package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func execGit(t *testing.T, arg ...string) string {
	t.Helper()

	output, err := exec.Command("git", arg...).Output()
	if err != nil {
		t.Fatal(err)
	}

	return strings.TrimSpace(string(output))
}

func TestAccDataSourceGitRepository_path(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.ToSlash(filepath.Join(cwd, "testdata", "terraform-git-module-acctest"))

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitRepositoryPathConfig(path),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.git_repository.test", "branch", "master"),
					resource.TestCheckResourceAttr("data.git_repository.test", "commit_sha", "750e948880ebe167eba524dded790e8b9a79d01d"),
				),
			},
		},
	})
}

func TestAccDataSourceGitRepository_branch(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.ToSlash(filepath.Join(cwd, "testdata", "terraform-git-module-acctest"))

	branch := "master"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitRepositoryBranchConfig(path, branch),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.git_repository.test", "branch", "master"),
					resource.TestCheckResourceAttr("data.git_repository.test", "commit_sha", "750e948880ebe167eba524dded790e8b9a79d01d"),
				),
			},
		},
	})
}

func TestAccDataSourceGitRepository_tag(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.ToSlash(filepath.Join(cwd, "testdata", "terraform-git-module-acctest"))

	tag := "v0.1.0"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitRepositoryTagConfig(path, tag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.git_repository.test", "tag", "v0.1.0"),
					resource.TestCheckResourceAttr("data.git_repository.test", "commit_sha", "750e948880ebe167eba524dded790e8b9a79d01d"),
				),
			},
		},
	})
}

func TestAccDataSourceGitRepository_HTTPURL(t *testing.T) {
	url := "https://github.com/innovationnorway/terraform-git-module-acctest.git"
	tag := "v0.1.0"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitRepositoryURLConfig(url, tag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.git_repository.test", "url", url),
					resource.TestCheckResourceAttr("data.git_repository.test", "tag", tag),
					resource.TestCheckResourceAttr("data.git_repository.test", "commit_sha", "750e948880ebe167eba524dded790e8b9a79d01d"),
				),
			},
		},
	})
}

func testAccDataSourceGitRepositoryPathConfig(path string) string {
	return fmt.Sprintf(`
data git_repository "test" {
  path = "%s"
}
`, path)
}

func testAccDataSourceGitRepositoryBranchConfig(path string, branch string) string {
	return fmt.Sprintf(`
data git_repository "test" {
  path   = "%s"
  branch = "%s" 
}
`, path, branch)
}

func testAccDataSourceGitRepositoryTagConfig(path string, tag string) string {
	return fmt.Sprintf(`
data git_repository "test" {
  path = "%s"
  tag  = "%s" 
}
`, path, tag)
}

func testAccDataSourceGitRepositoryURLConfig(url string, tag string) string {
	return fmt.Sprintf(`
data git_repository "test" {
  url = "%s"
  tag  = "%s" 
}
`, url, tag)
}
