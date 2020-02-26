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

func TestDataSourceGitRepository(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(cwd, "..", ".git")

	branch := execGit(t, "rev-parse", "--abbrev-ref", "HEAD")
	commit := execGit(t, "rev-parse", "HEAD")

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceGitRepositoryConfig(path),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.git_repository.test", "branch", strings.TrimSpace(branch)),
					resource.TestCheckResourceAttr("data.git_repository.test", "commit_sha", strings.TrimSpace(commit)),
				),
			},
		},
	})
}

func testDataSourceGitRepositoryConfig(path string) string {
	return fmt.Sprintf(`
data git_repository "test" {
  path = "%s"
}
`, path)
}
