package kubernetes

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_kubernetes_auth_backend_config", func(r *config.Resource) {

		// we need to map data_json properly
		r.ExternalName = config.IdentifierFromProvider

		r.References["backend"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-vault/apis/auth/v1alpha1.Backend",
		}
	})
}
