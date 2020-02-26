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
  path = ".git"
}
```

## Argument Reference

The following arguments are supported:

* `path` - (Optional) The repository path. Default is `.git`.

## Attributes Reference

* `branch` - The name of the current branch.

* `commit_sha` - The SHA-1 hash of the current commit.

* `tag` - The name of the most recent tag that points to the current commit.
