package git

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GIT_USERNAME", nil),
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("GIT_PASSWORD", nil),
			},

			"private_key_file": {
				Type:          schema.TypeString,
				Optional:      true,
				DefaultFunc:   schema.EnvDefaultFunc("GIT_PRIVATE_KEY_FILE", nil),
				ConflictsWith: []string{"username", "password"},
			},

			"ignore_host_key": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GIT_IGNORE_HOST_KEY", false),
			},

			"skip_tls_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GIT_SKIP_TLS_VERIFY", false),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"git_repository": dataSourceGitRepository(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username:              d.Get("username").(string),
		Password:              d.Get("password").(string),
		PrivateKeyFile:        d.Get("private_key_file").(string),
		InsecureIgnoreHostKey: d.Get("ignore_host_key").(bool),
		InsecureSkipTLSVerify: d.Get("skip_tls_verify").(bool),
	}

	return config.Client()
}
