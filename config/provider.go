/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-jet-vault/config/auth"
	"github.com/crossplane-contrib/provider-jet-vault/config/generic"
	"github.com/crossplane-contrib/provider-jet-vault/config/kubernetes"
	"github.com/crossplane-contrib/provider-jet-vault/config/mount"
	"github.com/crossplane-contrib/provider-jet-vault/config/pki"
	"github.com/crossplane-contrib/provider-jet-vault/config/policy"
	"github.com/crossplane-contrib/provider-jet-vault/config/token"
)

const (
	resourcePrefix = "vault"
	modulePath     = "github.com/crossplane-contrib/provider-jet-vault"
)

//go:embed schema.json
var providerSchema string

// IncludedResources lists all resource patterns included in small set release.
var IncludedResources = []string{
	// vault_auth*
	"vault_auth_backend$",

	// vault_generic*
	"vault_generic_secret$",

	// vault_kubernetes*
	"vault_kubernetes_auth_backend_config",

	// vault_mount*
	"vault_mount$",

	// vault_pki*
	"vault_pki_secret_backend_config_urls$",
	"vault_pki_secret_backend_root_cert$",
	"vault_pki_secret_backend_role$",

	// vault_policy*
	"vault_policy$",

	// vault_token*
	"vault_token_auth_backend_role$",
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList(IncludedResources))

	for _, configure := range []func(provider *tjconfig.Provider){
		auth.Configure,
		generic.Configure,
		kubernetes.Configure,
		mount.Configure,
		pki.Configure,
		policy.Configure,
		token.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
