package pki

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_pki_secret_backend_config_urls", func(r *config.Resource) {

		// we need to map data_json properly
		r.ExternalName = config.IdentifierFromProvider

		r.References["backend"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-vault/apis/mount/v1alpha1.Mount",
		}
	})

	p.AddResourceConfigurator("vault_pki_secret_backend_root_cert", func(r *config.Resource) {

		// we need to map data_json properly
		r.ExternalName = config.IdentifierFromProvider

		r.References["backend"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-vault/apis/mount/v1alpha1.Mount",
		}
	})
}
