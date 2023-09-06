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

type RoleAssignmentsObservation struct {

	// Unique identifier of the project to which you want the role mapping to apply.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Specifies the Roles that are attached to the Role Mapping. Available role IDs can be found on the User Roles
	// Reference.
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type RoleAssignmentsParameters struct {

	// Unique identifier of the project to which you want the role mapping to apply.
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Specifies the Roles that are attached to the Role Mapping. Available role IDs can be found on the User Roles
	// Reference.
	// +kubebuilder:validation:Optional
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type SettingsOrgRoleMappingObservation struct {

	// Unique human-readable label that identifies the identity provider group to which this role mapping applies.
	ExternalGroupName *string `json:"externalGroupName,omitempty" tf:"external_group_name,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsID *string `json:"federationSettingsId,omitempty" tf:"federation_settings_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies this role mapping.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Atlas roles and the unique identifiers of the groups and organizations associated with each role.
	RoleAssignments []RoleAssignmentsObservation `json:"roleAssignments,omitempty" tf:"role_assignments,omitempty"`
}

type SettingsOrgRoleMappingParameters struct {

	// Unique human-readable label that identifies the identity provider group to which this role mapping applies.
	// +kubebuilder:validation:Optional
	ExternalGroupName *string `json:"externalGroupName,omitempty" tf:"external_group_name,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	// +kubebuilder:validation:Optional
	FederationSettingsID *string `json:"federationSettingsId,omitempty" tf:"federation_settings_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Atlas roles and the unique identifiers of the groups and organizations associated with each role.
	// +kubebuilder:validation:Optional
	RoleAssignments []RoleAssignmentsParameters `json:"roleAssignments,omitempty" tf:"role_assignments,omitempty"`
}

// SettingsOrgRoleMappingSpec defines the desired state of SettingsOrgRoleMapping
type SettingsOrgRoleMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SettingsOrgRoleMappingParameters `json:"forProvider"`
}

// SettingsOrgRoleMappingStatus defines the observed state of SettingsOrgRoleMapping.
type SettingsOrgRoleMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SettingsOrgRoleMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SettingsOrgRoleMapping is the Schema for the SettingsOrgRoleMappings API. Provides a federated settings Role Mapping resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type SettingsOrgRoleMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.externalGroupName)",message="externalGroupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.federationSettingsId)",message="federationSettingsId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.orgId)",message="orgId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.roleAssignments)",message="roleAssignments is a required parameter"
	Spec   SettingsOrgRoleMappingSpec   `json:"spec"`
	Status SettingsOrgRoleMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SettingsOrgRoleMappingList contains a list of SettingsOrgRoleMappings
type SettingsOrgRoleMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SettingsOrgRoleMapping `json:"items"`
}

// Repository type metadata.
var (
	SettingsOrgRoleMapping_Kind             = "SettingsOrgRoleMapping"
	SettingsOrgRoleMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SettingsOrgRoleMapping_Kind}.String()
	SettingsOrgRoleMapping_KindAPIVersion   = SettingsOrgRoleMapping_Kind + "." + CRDGroupVersion.String()
	SettingsOrgRoleMapping_GroupVersionKind = CRDGroupVersion.WithKind(SettingsOrgRoleMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&SettingsOrgRoleMapping{}, &SettingsOrgRoleMappingList{})
}
