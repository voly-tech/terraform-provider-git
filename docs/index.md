# Git Provider

The Git provider is used to interact with [Git](https://git-scm.com/) repositories.

## Example Usage

```hcl
provider "git" {}

data "git_repository" "example" {
  path = path.root
}

resource "azurerm_resource_group" "example" {
  ...

  tags = {
    branch = data.git_repository.example.branch
    commit = substr(data.git_repository.example.commit_sha, 0, 7)
    tag    = data.git_repository.example.tag
  }
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `username` - (Optional) A Git username. This is used to access a remote repository over HTTP. It can also be sourced from the `GIT_USERNAME` environment variable.

* `password` - (Optional) A Git password. This is used to access a remote repository over HTTP. It can also be sourced from the `GIT_PASSWORD` environment variable.

* `private_key` - (Optional) A PEM-encoded private key. This is used to access a remote repository over SSH. It can also be sourced from the `GIT_PRIVATE_KEY` environment variable.

* `private_key_file` - (Optional) A path to a PEM-encoded private key. This is used to access a remote repository over SSH. It can also be sourced from the `GIT_PRIVATE_KEY_FILE` environment variable. Either this or `private_key` can be specified, but not both.

* `ignore_host_key` - (Optional) Set this to `true` to disable SSH host key verification. This will accept any host key and is strongly discouraged. It can also be sourced from the `GIT_IGNORE_HOST_KEY` environment variable. Default is `false`.

* `skip_tls_verify` - (Optional) Set this to `true` to disable verification of the server's TLS certificate chain. This is strongly discouraged. It can also be sourced from the `GIT_SKIP_TLS_VERIFY` environment variable. Default is `false`.
