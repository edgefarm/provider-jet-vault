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

type EntityObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EntityParameters struct {

	// Whether the entity is disabled. Disabled entities' associated tokens cannot be used, but are not revoked.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Manage policies externally through `vault_identity_entity_policies`.
	// +kubebuilder:validation:Optional
	ExternalPolicies *bool `json:"externalPolicies,omitempty" tf:"external_policies,omitempty"`

	// Metadata to be associated with the entity.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Name of the entity.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Policies to be tied to the entity.
	// +kubebuilder:validation:Optional
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

// EntitySpec defines the desired state of Entity
type EntitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EntityParameters `json:"forProvider"`
}

// EntityStatus defines the observed state of Entity.
type EntityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EntityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Entity is the Schema for the Entitys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vaultjet}
type Entity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EntitySpec   `json:"spec"`
	Status            EntityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EntityList contains a list of Entitys
type EntityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Entity `json:"items"`
}

// Repository type metadata.
var (
	Entity_Kind             = "Entity"
	Entity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Entity_Kind}.String()
	Entity_KindAPIVersion   = Entity_Kind + "." + CRDGroupVersion.String()
	Entity_GroupVersionKind = CRDGroupVersion.WithKind(Entity_Kind)
)

func init() {
	SchemeBuilder.Register(&Entity{}, &EntityList{})
}
