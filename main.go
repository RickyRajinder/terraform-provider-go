package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/rickyrajinder/terraform-provider-go/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}