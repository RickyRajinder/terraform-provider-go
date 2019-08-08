package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Schema:             map[string]*schema.Schema {
			"address": &schema.Schema {
				Type:             schema.TypeString,
				Required:         true,
			},
		},

		Create:             resourceServerCreate,
		Read:               resourceServerRead,
		Update:             resourceServerUpdate,
		Delete:             resourceServerDelete,
	}
}