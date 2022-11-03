package identity

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// IdentityPackagePath is the golang path for this package.
	IdentityPackagePath = "github.com/crossplane-contrib/provider-jet-vault/config/identity"
)

var (
	// PathSelfLinkExtractor is the golang path to SelfLinkExtractor function
	// in this package.
	PathAccessorExtractor = IdentityPackagePath + ".AccessorExtractor()"
)

// AccessorExtractor extracts accessor of the resources from
// "status.atProvider.accessor".
func AccessorExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.accessor")
		if err != nil {
			return ""
		}
		return r
	}
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("vault_identity_group", func(r *config.Resource) {

		// we need to map data_json properly
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("vault_identity_group_alias", func(r *config.Resource) {

		// we need to map data_json properly
		r.ExternalName = config.IdentifierFromProvider

		r.References["mount_accessor"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-jet-vault/apis/auth/v1alpha1.Backend",
			Extractor: PathAccessorExtractor,
		}

		r.References["canonical_id"] = config.Reference{
			Type: "Group",
		}
	})
}
