package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChompLeft() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceChompLeftRead,
		Schema: map[string]*schema.Schema{
			"lookup": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"separator": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"src": &schema.Schema{
				Type:        schema.TypeMap,
				Description: "Input goes here",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignore_error": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "",
			},
			"key": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
			"found": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "",
			},
		},
	}
}

func dataSourceChompLeftRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// input params
	src := d.Get("src").(map[string]interface{})
	lookup := d.Get("lookup").(string)
	separator := d.Get("separator").(string)
	ignore := d.Get("ignore_error").(bool)

	v, found := lchomp(lookup, separator, src)
	if !found && !ignore {
		return diag.FromErr(fmt.Errorf("no key found in scope: %v", lookup))
	}
	if err := d.Set("found", found); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("key", v); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func lchomp(key, sep string, m map[string]interface{}) (string, bool) {
	a := make([]string, 2) // always loop at least once
	for len(a) > 1 {
		if _, ok := m[key]; ok {
			return key, ok
		}
		a = strings.Split(key, sep)
		key = strings.Join(a[1:len(a)], sep)
	}
	return "", false
}
