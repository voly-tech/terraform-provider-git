provider "git" {
  private_key = var.private_key
}

data "git_repository" "example" {
  url = var.repo_url
}

output "repository" {
  value = data.git_repository.example
}
