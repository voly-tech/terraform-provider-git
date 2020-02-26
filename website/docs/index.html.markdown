---
layout: "git"
page_title: "Provider: Git"
sidebar_current: "docs-git-index"
description: |-
  The Git provider is used to interact with git repositories.
---

# Git Provider

The Git provider is used to interact with git repositories.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "git" {}

data "git_repository" "example" {
  path = ".git"
}
```
