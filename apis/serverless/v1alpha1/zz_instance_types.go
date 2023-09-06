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

type InstanceObservation struct {
	ConnectionStringsPrivateEndpointSrv []*string `json:"connectionStringsPrivateEndpointSrv,omitempty" tf:"connection_strings_private_endpoint_srv,omitempty"`

	ConnectionStringsStandardSrv *string `json:"connectionStringsStandardSrv,omitempty" tf:"connection_strings_standard_srv,omitempty"`

	ContinuousBackupEnabled *bool `json:"continuousBackupEnabled,omitempty" tf:"continuous_backup_enabled,omitempty"`

	CreateDate *string `json:"createDate,omitempty" tf:"create_date,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Links []LinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	MongoDBVersion *string `json:"mongoDbVersion,omitempty" tf:"mongo_db_version,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	ProviderSettingsBackingProviderName *string `json:"providerSettingsBackingProviderName,omitempty" tf:"provider_settings_backing_provider_name,omitempty"`

	ProviderSettingsProviderName *string `json:"providerSettingsProviderName,omitempty" tf:"provider_settings_provider_name,omitempty"`

	ProviderSettingsRegionName *string `json:"providerSettingsRegionName,omitempty" tf:"provider_settings_region_name,omitempty"`

	StateName *string `json:"stateName,omitempty" tf:"state_name,omitempty"`

	TerminationProtectionEnabled *bool `json:"terminationProtectionEnabled,omitempty" tf:"termination_protection_enabled,omitempty"`
}

type InstanceParameters struct {

	// +kubebuilder:validation:Optional
	ContinuousBackupEnabled *bool `json:"continuousBackupEnabled,omitempty" tf:"continuous_backup_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Links []LinksParameters `json:"links,omitempty" tf:"links,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProviderSettingsBackingProviderName *string `json:"providerSettingsBackingProviderName,omitempty" tf:"provider_settings_backing_provider_name,omitempty"`

	// +kubebuilder:validation:Optional
	ProviderSettingsProviderName *string `json:"providerSettingsProviderName,omitempty" tf:"provider_settings_provider_name,omitempty"`

	// +kubebuilder:validation:Optional
	ProviderSettingsRegionName *string `json:"providerSettingsRegionName,omitempty" tf:"provider_settings_region_name,omitempty"`

	// +kubebuilder:validation:Optional
	StateName *string `json:"stateName,omitempty" tf:"state_name,omitempty"`

	// +kubebuilder:validation:Optional
	TerminationProtectionEnabled *bool `json:"terminationProtectionEnabled,omitempty" tf:"termination_protection_enabled,omitempty"`
}

type LinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type LinksParameters struct {
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.projectId)",message="projectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.providerSettingsBackingProviderName)",message="providerSettingsBackingProviderName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.providerSettingsProviderName)",message="providerSettingsProviderName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.providerSettingsRegionName)",message="providerSettingsRegionName is a required parameter"
	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}