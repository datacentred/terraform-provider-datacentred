package datacentred

import (
	"fmt"
	"github.com/datacentred/datacentred-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDatacentredUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	user, err := datacentred.CreateUser(map[string]string{
		"email":      d.Get("email").(string),
		"password":   d.Get("password").(string),
		"first_name": d.Get("first_name").(string),
		"last_name":  d.Get("last_name").(string),
	})

	if err != nil {
		return fmt.Errorf("Error creating Datacentred user: %s", err)
	}

	d.SetId(user.Id)

	return resourceUserRead(d, meta)
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	user, err := datacentred.FindUser(d.Id())
	if err != nil {
		return fmt.Errorf("Error finding Datacentred user: %s", err)
	}

	d.Set("email", user.Email)
	d.Set("first_name", user.FirstName)
	d.Set("last_name", user.LastName)

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	user, err := datacentred.FindUser(d.Id())
	if err != nil {
		return fmt.Errorf("Error finding Datacentred user: %s", err)
	}

	var hasChange bool

	if d.HasChange("first_name") {
		hasChange = true
		user.FirstName = d.Get("first_name").(string)
	}

	if d.HasChange("last_name") {
		hasChange = true
		user.LastName = d.Get("last_name").(string)
	}

	if d.HasChange("email") {
		hasChange = true
		user.Email = d.Get("email").(string)
	}

	if d.HasChange("password") {
		user.ChangePassword(d.Get("password").(string))
	}

	if hasChange {
		_, err := user.Save()
		if err != nil {
			return fmt.Errorf("Error updating Datacentred user: %s", err)
		}
	}

	return resourceUserRead(d, meta)
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	user, err := datacentred.FindUser(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting Datacentred user: %s", err)
	}
	_, err = user.Destroy()
	if err != nil {
		return fmt.Errorf("Error deleting Datacentred user: %s", err)
	}

	return nil
}
