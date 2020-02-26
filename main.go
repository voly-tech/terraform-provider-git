package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/volcano-coffee-company/terraform-provider-git/git"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: git.Provider})
}
