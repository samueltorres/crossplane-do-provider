package vpc

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("digitalocean_vpc", func(r *config.Resource) {
		r.ShortGroup = "vpc"

		r.ExternalName = config.IdentifierFromProvider
	})
}
