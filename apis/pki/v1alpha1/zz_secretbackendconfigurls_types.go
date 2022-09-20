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

type SecretBackendConfigUrlsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecretBackendConfigUrlsParameters struct {

	// The path of the PKI secret backend the resource belongs to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-vault/apis/mount/v1alpha1.Mount
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// Specifies the URL values for the CRL Distribution Points field.
	// +kubebuilder:validation:Optional
	CrlDistributionPoints []*string `json:"crlDistributionPoints,omitempty" tf:"crl_distribution_points,omitempty"`

	// Specifies the URL values for the Issuing Certificate field.
	// +kubebuilder:validation:Optional
	IssuingCertificates []*string `json:"issuingCertificates,omitempty" tf:"issuing_certificates,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies the URL values for the OCSP Servers field.
	// +kubebuilder:validation:Optional
	OcspServers []*string `json:"ocspServers,omitempty" tf:"ocsp_servers,omitempty"`
}

// SecretBackendConfigUrlsSpec defines the desired state of SecretBackendConfigUrls
type SecretBackendConfigUrlsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendConfigUrlsParameters `json:"forProvider"`
}

// SecretBackendConfigUrlsStatus defines the observed state of SecretBackendConfigUrls.
type SecretBackendConfigUrlsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendConfigUrlsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendConfigUrls is the Schema for the SecretBackendConfigUrlss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vaultjet}
type SecretBackendConfigUrls struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretBackendConfigUrlsSpec   `json:"spec"`
	Status            SecretBackendConfigUrlsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendConfigUrlsList contains a list of SecretBackendConfigUrlss
type SecretBackendConfigUrlsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendConfigUrls `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendConfigUrls_Kind             = "SecretBackendConfigUrls"
	SecretBackendConfigUrls_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendConfigUrls_Kind}.String()
	SecretBackendConfigUrls_KindAPIVersion   = SecretBackendConfigUrls_Kind + "." + CRDGroupVersion.String()
	SecretBackendConfigUrls_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendConfigUrls_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendConfigUrls{}, &SecretBackendConfigUrlsList{})
}
