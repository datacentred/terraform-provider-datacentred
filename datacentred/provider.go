package datacentred

import (
	"github.com/datacentred/datacentred-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"datacentred_access_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATACENTRED_ACCESS", nil),
				Description: "Datacentred API access key",
			},
			"datacentred_secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATACENTRED_SECRET", nil),
				Description: "Datacentred API secret key",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"datacentred_user": resourceDatacentredUser(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	datacentred.Config.AccessKey = d.Get("datacentred_access_key").(string)
	datacentred.Config.SecretKey = d.Get("datacentred_secret_key").(string)

	return &datacentred.Config, nil
}
