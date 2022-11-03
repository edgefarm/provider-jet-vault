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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	backend "github.com/crossplane-contrib/provider-jet-vault/internal/controller/auth/backend"
	secret "github.com/crossplane-contrib/provider-jet-vault/internal/controller/generic/secret"
	group "github.com/crossplane-contrib/provider-jet-vault/internal/controller/identity/group"
	groupalias "github.com/crossplane-contrib/provider-jet-vault/internal/controller/identity/groupalias"
	authbackendconfig "github.com/crossplane-contrib/provider-jet-vault/internal/controller/kubernetes/authbackendconfig"
	mount "github.com/crossplane-contrib/provider-jet-vault/internal/controller/mount/mount"
	secretbackendconfigurls "github.com/crossplane-contrib/provider-jet-vault/internal/controller/pki/secretbackendconfigurls"
	secretbackendrole "github.com/crossplane-contrib/provider-jet-vault/internal/controller/pki/secretbackendrole"
	secretbackendrootcert "github.com/crossplane-contrib/provider-jet-vault/internal/controller/pki/secretbackendrootcert"
	policy "github.com/crossplane-contrib/provider-jet-vault/internal/controller/policy/policy"
	providerconfig "github.com/crossplane-contrib/provider-jet-vault/internal/controller/providerconfig"
	authbackendrole "github.com/crossplane-contrib/provider-jet-vault/internal/controller/token/authbackendrole"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backend.Setup,
		secret.Setup,
		group.Setup,
		groupalias.Setup,
		authbackendconfig.Setup,
		mount.Setup,
		secretbackendconfigurls.Setup,
		secretbackendrole.Setup,
		secretbackendrootcert.Setup,
		policy.Setup,
		providerconfig.Setup,
		authbackendrole.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
