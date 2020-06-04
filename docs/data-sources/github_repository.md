# `git_repository` Data Source

Use this data source to retrieve information about a Git repository.

## Example Usage

```hcl
data "git_repository" "example" {
  path = path.cwd
}
```

## Argument Reference

The following arguments are supported:

* `path` - (Optional) The repository path. It can also be sourced from the `GIT_DIR` environment variable.

* `url` - (Optional) The repository URL.

-> **NOTE:** Either a `path` or `url` must be specified - but not both.

* `branch` - (Optional) The name of the branch.

* `tag` - (Optional) The name of the tag.

## Attributes Reference

* `commit_sha` - The SHA-1 hash of the current commit.
