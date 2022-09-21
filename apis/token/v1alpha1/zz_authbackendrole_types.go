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

type AuthBackendRoleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AuthBackendRoleParameters struct {

	// Set of allowed entity aliases for this role.
	// +kubebuilder:validation:Optional
	AllowedEntityAliases []*string `json:"allowedEntityAliases,omitempty" tf:"allowed_entity_aliases,omitempty"`

	// List of allowed policies for given role.
	// +kubebuilder:validation:Optional
	AllowedPolicies []*string `json:"allowedPolicies,omitempty" tf:"allowed_policies,omitempty"`

	// Set of allowed policies with glob match for given role.
	// +kubebuilder:validation:Optional
	AllowedPoliciesGlob []*string `json:"allowedPoliciesGlob,omitempty" tf:"allowed_policies_glob,omitempty"`

	// List of disallowed policies for given role.
	// +kubebuilder:validation:Optional
	DisallowedPolicies []*string `json:"disallowedPolicies,omitempty" tf:"disallowed_policies,omitempty"`

	// Set of disallowed policies with glob match for given role.
	// +kubebuilder:validation:Optional
	DisallowedPoliciesGlob []*string `json:"disallowedPoliciesGlob,omitempty" tf:"disallowed_policies_glob,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// If true, tokens created against this policy will be orphan tokens.
	// +kubebuilder:validation:Optional
	Orphan *bool `json:"orphan,omitempty" tf:"orphan,omitempty"`

	// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
	// +kubebuilder:validation:Optional
	PathSuffix *string `json:"pathSuffix,omitempty" tf:"path_suffix,omitempty"`

	// Whether to disable the ability of the token to be renewed past its initial TTL.
	// +kubebuilder:validation:Optional
	Renewable *bool `json:"renewable,omitempty" tf:"renewable,omitempty"`

	// Name of the role.
	// +kubebuilder:validation:Required
	RoleName *string `json:"roleName" tf:"role_name,omitempty"`

	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +kubebuilder:validation:Optional
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// Generated Token's Explicit Maximum TTL in seconds
	// +kubebuilder:validation:Optional
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	// +kubebuilder:validation:Optional
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	// +kubebuilder:validation:Optional
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	// +kubebuilder:validation:Optional
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// Generated Token's Period
	// +kubebuilder:validation:Optional
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// Generated Token's Policies
	// +kubebuilder:validation:Optional
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The initial ttl of the token to generate in seconds
	// +kubebuilder:validation:Optional
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

// AuthBackendRoleSpec defines the desired state of AuthBackendRole
type AuthBackendRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendRoleParameters `json:"forProvider"`
}

// AuthBackendRoleStatus defines the observed state of AuthBackendRole.
type AuthBackendRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRole is the Schema for the AuthBackendRoles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vaultjet}
type AuthBackendRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthBackendRoleSpec   `json:"spec"`
	Status            AuthBackendRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRoleList contains a list of AuthBackendRoles
type AuthBackendRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendRole `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendRole_Kind             = "AuthBackendRole"
	AuthBackendRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendRole_Kind}.String()
	AuthBackendRole_KindAPIVersion   = AuthBackendRole_Kind + "." + CRDGroupVersion.String()
	AuthBackendRole_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendRole_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendRole{}, &AuthBackendRoleList{})
}