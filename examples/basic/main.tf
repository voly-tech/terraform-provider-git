terraform {
  required_providers {
    git = {
      source  = "innovationnorway/git"
      version = ">= 0.1.3"
    }
  }
}

data "git_repository" "example" {
  url = "https://github.com/innovationnorway/terraform-provider-git.git"
}

output "repository" {
  value = {
    branch = data.git_repository.example.branch
    commit = substr(data.git_repository.example.commit_sha, 0, 7)
    tag    = data.git_repository.example.tag
  }
}
