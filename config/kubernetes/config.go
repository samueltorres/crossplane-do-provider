package kubernetes

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("digitalocean_kubernetes_cluster", func(r *config.Resource) {
		r.ShortGroup = "kubernetes"
		r.ExternalName = config.IdentifierFromProvider

		r.References["vpcUuid"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-do/apis/vpc/v1alpha1.Vpc",
		}

		r.References["nodePool"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-do/apis/kubernetes/v1alpha1.NodePool",
		}
	})

	p.AddResourceConfigurator("digitalocean_kubernetes_node_pool", func(r *config.Resource) {
		r.ShortGroup = "kubernetes"
		r.ExternalName = config.IdentifierFromProvider
	})
}
