package identity

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_identity_entity", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "vault"
		r.ShortGroup = "identity"

        // Identifier for this resource is assigned by the provider. In other
        // words it is not simply the name of the resource
		r.ExternalName = config.IdentifierFromProvider

	})
}
