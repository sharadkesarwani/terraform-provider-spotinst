package main

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/spotinst/terraform-provider-spotinst/spotinst"
	"log"
)

func main() {
	var debugMode bool
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return spotinst.Provider()
		},
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "terraform-spotinst/local/spotinst", opts)
		//err := plugin.Debug(context.Background(), "C:/Users/skesarwa/terraformfiles/spotinst/v2/resources/elastigroup_aws/general/terraform-spotinst/local/spotinst", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
