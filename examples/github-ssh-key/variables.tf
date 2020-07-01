variable "repo_url" {
  type        = string
  description = "The repository URL. Example: git@github.com:octocat/Hello-World.git"
}

variable "private_key" {
  type        = string
  description = "A PEM-encoded private key."
}
