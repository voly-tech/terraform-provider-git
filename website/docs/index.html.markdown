---
layout: "git"
page_title: "Provider: Git"
sidebar_current: "docs-git-index"
description: |-
  The Git provider is used to interact with git repositories.
---

# Git Provider

The Git provider is used to interact with [Git](https://git-scm.com/) repositories.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "git" {}

data "git_repository" "example" {
  path = path.cwd
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `username` - (Optional) A Git username. This is used to access a remote repository over HTTP. It can also be sourced from the `GIT_USERNAME` environment variable.

* `password` - (Optional) A Git password. This is used to access a remote repository over HTTP. It can also be sourced from the `GIT_PASSWORD` environment variable.

* `private_key_file` - (Optional) A path to a PEM-encoded private key. This is used to access a remote repository over SSH. It can also be sourced from the `GIT_PRIVATE_KEY_FILE` environment variable.

* `ignore_host_key` - (Optional) Set this to `true` to disable SSH host key verification. This will accept any host key and is strongly discouraged. It can also be sourced from the `GIT_IGNORE_HOST_KEY` environment variable. Default is `false`.

* `skip_tls_verify` - (Optional) Set this to `true` to disable verification of the server's TLS certificate chain. This is strongly discouraged. It can also be sourced from the `GIT_SKIP_TLS_VERIFY` environment variable. Default is `false`.
