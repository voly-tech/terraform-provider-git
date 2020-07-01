provider "git" {
  username = "x-access-token"
  password = var.github_token
}

data "git_repository" "example" {
  url = var.repo_url
}

output "repository" {
  value = data.git_repository.example
}
