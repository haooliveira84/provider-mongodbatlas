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

type TeamsObservation_2 struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the team you want to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The unique identifier for the organization you want to associate the team with.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The unique identifier for the team.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// The Atlas usernames (email address). You can only add Atlas users who are part of the organization. Users who have not accepted an invitation to join the organization cannot be added as team members. There is a maximum of 250 Atlas users per team.
	Usernames []*string `json:"usernames,omitempty" tf:"usernames,omitempty"`
}

type TeamsParameters_2 struct {

	// The name of the team you want to create.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The unique identifier for the organization you want to associate the team with.
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The Atlas usernames (email address). You can only add Atlas users who are part of the organization. Users who have not accepted an invitation to join the organization cannot be added as team members. There is a maximum of 250 Atlas users per team.
	// +kubebuilder:validation:Optional
	Usernames []*string `json:"usernames,omitempty" tf:"usernames,omitempty"`
}

// TeamsSpec defines the desired state of Teams
type TeamsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TeamsParameters_2 `json:"forProvider"`
}

// TeamsStatus defines the observed state of Teams.
type TeamsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TeamsObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Teams is the Schema for the Teamss API. Provides a Team resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type Teams struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.orgId)",message="orgId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.usernames)",message="usernames is a required parameter"
	Spec   TeamsSpec   `json:"spec"`
	Status TeamsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamsList contains a list of Teamss
type TeamsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Teams `json:"items"`
}

// Repository type metadata.
var (
	Teams_Kind             = "Teams"
	Teams_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Teams_Kind}.String()
	Teams_KindAPIVersion   = Teams_Kind + "." + CRDGroupVersion.String()
	Teams_GroupVersionKind = CRDGroupVersion.WithKind(Teams_Kind)
)

func init() {
	SchemeBuilder.Register(&Teams{}, &TeamsList{})
}
