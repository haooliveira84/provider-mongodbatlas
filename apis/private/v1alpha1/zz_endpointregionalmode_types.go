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

type EndpointRegionalModeObservation struct {

	// Flag that indicates whether the regionalized private endpoint setting is enabled for the project.   Set this value to true to create more than one private endpoint in a cloud provider region to connect to multi-region and global Atlas sharded clusters. You can enable this setting only if your Atlas project contains no replica sets. You can't disable this setting if you have:
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique identifier for the project.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type EndpointRegionalModeParameters struct {

	// Flag that indicates whether the regionalized private endpoint setting is enabled for the project.   Set this value to true to create more than one private endpoint in a cloud provider region to connect to multi-region and global Atlas sharded clusters. You can enable this setting only if your Atlas project contains no replica sets. You can't disable this setting if you have:
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Unique identifier for the project.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

// EndpointRegionalModeSpec defines the desired state of EndpointRegionalMode
type EndpointRegionalModeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointRegionalModeParameters `json:"forProvider"`
}

// EndpointRegionalModeStatus defines the observed state of EndpointRegionalMode.
type EndpointRegionalModeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointRegionalModeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointRegionalMode is the Schema for the EndpointRegionalModes API. Provides a Private Endpoint Regional Mode resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type EndpointRegionalMode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.projectId)",message="projectId is a required parameter"
	Spec   EndpointRegionalModeSpec   `json:"spec"`
	Status EndpointRegionalModeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointRegionalModeList contains a list of EndpointRegionalModes
type EndpointRegionalModeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointRegionalMode `json:"items"`
}

// Repository type metadata.
var (
	EndpointRegionalMode_Kind             = "EndpointRegionalMode"
	EndpointRegionalMode_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointRegionalMode_Kind}.String()
	EndpointRegionalMode_KindAPIVersion   = EndpointRegionalMode_Kind + "." + CRDGroupVersion.String()
	EndpointRegionalMode_GroupVersionKind = CRDGroupVersion.WithKind(EndpointRegionalMode_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointRegionalMode{}, &EndpointRegionalModeList{})
}
