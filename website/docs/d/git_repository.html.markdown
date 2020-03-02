---
layout: "git"
page_title: "Git: git_repository"
sidebar_current: "docs-git-datasource-repository"
description: |-
  Get details about a git repository.
---

# git_repository

Use this data source to retrieve information about a git repository.

## Example Usage

```hcl
data "git_repository" "example" {
  path = path.cwd
}
```

## Argument Reference

The following arguments are supported:

* `path` - (Optional) The repository path. It can also be sourced from the `GIT_DIR` environment variable. Default is `.git`.

* `url` - (Optional) The repository URL.

* `branch` - (Optional) The name of the branch.

* `tag` - (Optional) The name of the tag.

## Attributes Reference

* `commit_sha` - The SHA-1 hash of the current commit.
