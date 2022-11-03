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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthBackendConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AuthBackendConfigParameters struct {

	// Unique name of the kubernetes backend to configure.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-vault/apis/auth/v1alpha1.Backend
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// Optional disable JWT issuer validation. Allows to skip ISS validation.
	// +kubebuilder:validation:Optional
	DisableIssValidation *bool `json:"disableIssValidation,omitempty" tf:"disable_iss_validation,omitempty"`

	// Optional disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod.
	// +kubebuilder:validation:Optional
	DisableLocalCAJwt *bool `json:"disableLocalCaJwt,omitempty" tf:"disable_local_ca_jwt,omitempty"`

	// Optional JWT issuer. If no issuer is specified, kubernetes.io/serviceaccount will be used as the default issuer.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	// +kubebuilder:validation:Optional
	KubernetesCACert *string `json:"kubernetesCaCert,omitempty" tf:"kubernetes_ca_cert,omitempty"`

	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	// +kubebuilder:validation:Required
	KubernetesHost *string `json:"kubernetesHost" tf:"kubernetes_host,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	// +kubebuilder:validation:Optional
	PemKeys []*string `json:"pemKeys,omitempty" tf:"pem_keys,omitempty"`

	// A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
	// +kubebuilder:validation:Optional
	TokenReviewerJwtSecretRef *v1.SecretKeySelector `json:"tokenReviewerJwtSecretRef,omitempty" tf:"-"`
}

// AuthBackendConfigSpec defines the desired state of AuthBackendConfig
type AuthBackendConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendConfigParameters `json:"forProvider"`
}

// AuthBackendConfigStatus defines the observed state of AuthBackendConfig.
type AuthBackendConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendConfig is the Schema for the AuthBackendConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vaultjet}
type AuthBackendConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthBackendConfigSpec   `json:"spec"`
	Status            AuthBackendConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendConfigList contains a list of AuthBackendConfigs
type AuthBackendConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendConfig `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendConfig_Kind             = "AuthBackendConfig"
	AuthBackendConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendConfig_Kind}.String()
	AuthBackendConfig_KindAPIVersion   = AuthBackendConfig_Kind + "." + CRDGroupVersion.String()
	AuthBackendConfig_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendConfig{}, &AuthBackendConfigList{})
}