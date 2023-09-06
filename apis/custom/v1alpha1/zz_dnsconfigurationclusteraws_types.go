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

type DNSConfigurationClusterAwsObservation struct {

	// Indicates whether the project's clusters deployed to AWS use custom DNS. If true, the Get All Clusters and Get One Cluster endpoints return the connectionStrings.private and connectionStrings.privateSrv fields for clusters deployed to AWS .
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Required 	Unique identifier for the project.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type DNSConfigurationClusterAwsParameters struct {

	// Indicates whether the project's clusters deployed to AWS use custom DNS. If true, the Get All Clusters and Get One Cluster endpoints return the connectionStrings.private and connectionStrings.privateSrv fields for clusters deployed to AWS .
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Required 	Unique identifier for the project.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

// DNSConfigurationClusterAwsSpec defines the desired state of DNSConfigurationClusterAws
type DNSConfigurationClusterAwsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DNSConfigurationClusterAwsParameters `json:"forProvider"`
}

// DNSConfigurationClusterAwsStatus defines the observed state of DNSConfigurationClusterAws.
type DNSConfigurationClusterAwsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DNSConfigurationClusterAwsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DNSConfigurationClusterAws is the Schema for the DNSConfigurationClusterAwss API. Provides a Custom DNS Configuration for Atlas Clusters on AWS resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type DNSConfigurationClusterAws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.enabled)",message="enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.projectId)",message="projectId is a required parameter"
	Spec   DNSConfigurationClusterAwsSpec   `json:"spec"`
	Status DNSConfigurationClusterAwsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSConfigurationClusterAwsList contains a list of DNSConfigurationClusterAwss
type DNSConfigurationClusterAwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSConfigurationClusterAws `json:"items"`
}

// Repository type metadata.
var (
	DNSConfigurationClusterAws_Kind             = "DNSConfigurationClusterAws"
	DNSConfigurationClusterAws_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DNSConfigurationClusterAws_Kind}.String()
	DNSConfigurationClusterAws_KindAPIVersion   = DNSConfigurationClusterAws_Kind + "." + CRDGroupVersion.String()
	DNSConfigurationClusterAws_GroupVersionKind = CRDGroupVersion.WithKind(DNSConfigurationClusterAws_Kind)
)

func init() {
	SchemeBuilder.Register(&DNSConfigurationClusterAws{}, &DNSConfigurationClusterAwsList{})
}