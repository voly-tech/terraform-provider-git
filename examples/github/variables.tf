variable "repo_url" {
  type        = string
  description = "The repository URL. Example: https://github.com/octocat/Hello-World.git"
}

variable "github_token" {
  type        = string
  description = "A personal access token with the `repo` scope."
}
