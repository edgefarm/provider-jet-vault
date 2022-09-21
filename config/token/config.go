package token

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_token_auth_backend_role", func(r *config.Resource) {

		// we need to map data_json properly
		r.ExternalName = config.IdentifierFromProvider

	})
}
