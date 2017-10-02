package main

import (
	"github.com/datacentred/terraform-provider-datacentred/datacentred"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

// func main() {
// 	plugin.Serve(&plugin.ServeOpts{
// 		ProviderFunc: template.Provider})
// }

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return datacentred.Provider()
		},
	})
}
