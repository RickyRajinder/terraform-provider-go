package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap:     map[string]*schema.Resource{
				"resource_server": resourceServer(),
		},
	}
}