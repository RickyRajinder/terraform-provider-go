package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"time"
)

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(time.Now().UTC().String())
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	resolver := d.Get("resolver").(string)

	ip, err := getIP(resolver)

	if err == nil {
		d.Set("ip_address", string(ip))
		d.SetId(time.Now().UTC().String())
	} else {
		return fmt.Errorf("error requesting external IP: %d", err)
	}
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	if d.HasChange("resolver") {
		if err := resourceServerCreate(d, m); err != nil {
			return err
		}
		d.SetPartial("resolver")
	}
	d.Partial(false)

	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Schema:             map[string]*schema.Schema {
			"resolver": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "https://checkip.amazonaws.com",
				Description: "URL used to resolve external IP address",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External IP address",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},

		Create:             resourceServerCreate,
		Read:               resourceServerRead,
		Update:             resourceServerUpdate,
		Delete:             resourceServerDelete,
	}
}