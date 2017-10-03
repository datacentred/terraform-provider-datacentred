package datacentred

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATACENTRED_ACCESS", nil),
				Description: "Datacentred API access key",
			},
			"secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATACENTRED_SECRET", nil),
				Description: "Datacentred API secret key",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"datacentred_user": resourceDatacentredUser(),
		},
	}
}
