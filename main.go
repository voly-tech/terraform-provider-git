package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/innovationnorway/terraform-provider-git/internal/git"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: git.Provider})
}
