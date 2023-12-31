/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OrganizationObservation struct {

	// Programmatic API Key description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the federation to link the newly created organization to. If specified, the proposed Organization Owner of the new organization must have the Organization Owner role in an organization associated with the federation.
	FederationSettingsID *string `json:"federationSettingsId,omitempty" tf:"federation_settings_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the organization you want to create. (Cannot be changed via this Provider after creation.)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization id.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the Atlas user that you want to assign the Organization Owner role. This user must be a member of the same organization as the calling API key.  This is only required when authenticating with Programmatic API Keys. MongoDB Atlas Admin API - Get User By Username
	OrgOwnerID *string `json:"orgOwnerId,omitempty" tf:"org_owner_id,omitempty"`

	// List of Organization roles that the Programmatic API key needs to have. Ensure that you provide at least one role and ensure all roles are valid for the Organization.  You must specify an array even if you are only associating a single role with the Programmatic API key. The MongoDB Documentation describes the roles that you can assign to a Programmatic API key.
	RoleNames []*string `json:"roleNames,omitempty" tf:"role_names,omitempty"`
}

type OrganizationParameters struct {

	// Programmatic API Key description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the federation to link the newly created organization to. If specified, the proposed Organization Owner of the new organization must have the Organization Owner role in an organization associated with the federation.
	// +kubebuilder:validation:Optional
	FederationSettingsID *string `json:"federationSettingsId,omitempty" tf:"federation_settings_id,omitempty"`

	// The name of the organization you want to create. (Cannot be changed via this Provider after creation.)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the Atlas user that you want to assign the Organization Owner role. This user must be a member of the same organization as the calling API key.  This is only required when authenticating with Programmatic API Keys. MongoDB Atlas Admin API - Get User By Username
	// +kubebuilder:validation:Optional
	OrgOwnerID *string `json:"orgOwnerId,omitempty" tf:"org_owner_id,omitempty"`

	// List of Organization roles that the Programmatic API key needs to have. Ensure that you provide at least one role and ensure all roles are valid for the Organization.  You must specify an array even if you are only associating a single role with the Programmatic API key. The MongoDB Documentation describes the roles that you can assign to a Programmatic API key.
	// +kubebuilder:validation:Optional
	RoleNames []*string `json:"roleNames,omitempty" tf:"role_names,omitempty"`
}

// OrganizationSpec defines the desired state of Organization
type OrganizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationParameters `json:"forProvider"`
}

// OrganizationStatus defines the observed state of Organization.
type OrganizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Organization is the Schema for the Organizations API. Provides a Organization resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type Organization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.description)",message="description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.orgOwnerId)",message="orgOwnerId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.roleNames)",message="roleNames is a required parameter"
	Spec   OrganizationSpec   `json:"spec"`
	Status OrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationList contains a list of Organizations
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Organization `json:"items"`
}

// Repository type metadata.
var (
	Organization_Kind             = "Organization"
	Organization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Organization_Kind}.String()
	Organization_KindAPIVersion   = Organization_Kind + "." + CRDGroupVersion.String()
	Organization_GroupVersionKind = CRDGroupVersion.WithKind(Organization_Kind)
)

func init() {
	SchemeBuilder.Register(&Organization{}, &OrganizationList{})
}
