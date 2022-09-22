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

type BackendObservation struct {
	Accessor *string `json:"accessor,omitempty" tf:"accessor,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackendParameters struct {

	// The description of the auth backend
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the auth method is local only
	// +kubebuilder:validation:Optional
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// path to mount the backend. This defaults to the type.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Tune []TuneParameters `json:"tune,omitempty" tf:"tune,omitempty"`

	// Name of the auth backend
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type TuneObservation struct {
}

type TuneParameters struct {

	// +kubebuilder:validation:Optional
	AllowedResponseHeaders []*string `json:"allowedResponseHeaders,omitempty" tf:"allowed_response_headers"`

	// +kubebuilder:validation:Optional
	AuditNonHMACRequestKeys []*string `json:"auditNonHmacRequestKeys,omitempty" tf:"audit_non_hmac_request_keys"`

	// +kubebuilder:validation:Optional
	AuditNonHMACResponseKeys []*string `json:"auditNonHmacResponseKeys,omitempty" tf:"audit_non_hmac_response_keys"`

	// +kubebuilder:validation:Optional
	DefaultLeaseTTL *string `json:"defaultLeaseTtl,omitempty" tf:"default_lease_ttl"`

	// +kubebuilder:validation:Optional
	ListingVisibility *string `json:"listingVisibility,omitempty" tf:"listing_visibility"`

	// +kubebuilder:validation:Optional
	MaxLeaseTTL *string `json:"maxLeaseTtl,omitempty" tf:"max_lease_ttl"`

	// +kubebuilder:validation:Optional
	PassthroughRequestHeaders []*string `json:"passthroughRequestHeaders,omitempty" tf:"passthrough_request_headers"`

	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type"`
}

// BackendSpec defines the desired state of Backend
type BackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendParameters `json:"forProvider"`
}

// BackendStatus defines the observed state of Backend.
type BackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Backend is the Schema for the Backends API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vaultjet}
type Backend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackendSpec   `json:"spec"`
	Status            BackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendList contains a list of Backends
type BackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backend `json:"items"`
}

// Repository type metadata.
var (
	Backend_Kind             = "Backend"
	Backend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Backend_Kind}.String()
	Backend_KindAPIVersion   = Backend_Kind + "." + CRDGroupVersion.String()
	Backend_GroupVersionKind = CRDGroupVersion.WithKind(Backend_Kind)
)

func init() {
	SchemeBuilder.Register(&Backend{}, &BackendList{})
}
