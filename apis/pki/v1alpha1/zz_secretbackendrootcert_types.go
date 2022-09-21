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

type SecretBackendRootCertObservation struct {
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IssuingCA *string `json:"issuingCa,omitempty" tf:"issuing_ca,omitempty"`

	Serial *string `json:"serial,omitempty" tf:"serial,omitempty"`

	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`
}

type SecretBackendRootCertParameters struct {

	// List of alternative names.
	// +kubebuilder:validation:Optional
	AltNames []*string `json:"altNames,omitempty" tf:"alt_names,omitempty"`

	// The PKI secret backend the resource belongs to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-vault/apis/mount/v1alpha1.Mount
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// CN of root to create.
	// +kubebuilder:validation:Required
	CommonName *string `json:"commonName" tf:"common_name,omitempty"`

	// The country.
	// +kubebuilder:validation:Optional
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// Flag to exclude CN from SANs.
	// +kubebuilder:validation:Optional
	ExcludeCnFromSans *bool `json:"excludeCnFromSans,omitempty" tf:"exclude_cn_from_sans,omitempty"`

	// The format of data.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// List of alternative IPs.
	// +kubebuilder:validation:Optional
	IPSans []*string `json:"ipSans,omitempty" tf:"ip_sans,omitempty"`

	// The number of bits to use.
	// +kubebuilder:validation:Optional
	KeyBits *float64 `json:"keyBits,omitempty" tf:"key_bits,omitempty"`

	// The desired key type.
	// +kubebuilder:validation:Optional
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// The locality.
	// +kubebuilder:validation:Optional
	Locality *string `json:"locality,omitempty" tf:"locality,omitempty"`

	// The maximum path length to encode in the generated certificate.
	// +kubebuilder:validation:Optional
	MaxPathLength *float64 `json:"maxPathLength,omitempty" tf:"max_path_length,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The organization.
	// +kubebuilder:validation:Optional
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// List of other SANs.
	// +kubebuilder:validation:Optional
	OtherSans []*string `json:"otherSans,omitempty" tf:"other_sans,omitempty"`

	// The organization unit.
	// +kubebuilder:validation:Optional
	Ou *string `json:"ou,omitempty" tf:"ou,omitempty"`

	// List of domains for which certificates are allowed to be issued.
	// +kubebuilder:validation:Optional
	PermittedDNSDomains []*string `json:"permittedDnsDomains,omitempty" tf:"permitted_dns_domains,omitempty"`

	// The postal code.
	// +kubebuilder:validation:Optional
	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	// The private key format.
	// +kubebuilder:validation:Optional
	PrivateKeyFormat *string `json:"privateKeyFormat,omitempty" tf:"private_key_format,omitempty"`

	// The province.
	// +kubebuilder:validation:Optional
	Province *string `json:"province,omitempty" tf:"province,omitempty"`

	// The street address.
	// +kubebuilder:validation:Optional
	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`

	// Time to live.
	// +kubebuilder:validation:Optional
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Type of root to create. Must be either "exported" or "internal".
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// List of alternative URIs.
	// +kubebuilder:validation:Optional
	URISans []*string `json:"uriSans,omitempty" tf:"uri_sans,omitempty"`
}

// SecretBackendRootCertSpec defines the desired state of SecretBackendRootCert
type SecretBackendRootCertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendRootCertParameters `json:"forProvider"`
}

// SecretBackendRootCertStatus defines the observed state of SecretBackendRootCert.
type SecretBackendRootCertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendRootCertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendRootCert is the Schema for the SecretBackendRootCerts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vaultjet}
type SecretBackendRootCert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretBackendRootCertSpec   `json:"spec"`
	Status            SecretBackendRootCertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendRootCertList contains a list of SecretBackendRootCerts
type SecretBackendRootCertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendRootCert `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendRootCert_Kind             = "SecretBackendRootCert"
	SecretBackendRootCert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendRootCert_Kind}.String()
	SecretBackendRootCert_KindAPIVersion   = SecretBackendRootCert_Kind + "." + CRDGroupVersion.String()
	SecretBackendRootCert_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendRootCert_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendRootCert{}, &SecretBackendRootCertList{})
}
