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

type QueryLimitObservation struct {
	CurrentUsage *float64 `json:"currentUsage,omitempty" tf:"current_usage,omitempty"`

	DefaultLimit *float64 `json:"defaultLimit,omitempty" tf:"default_limit,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastModifiedDate *string `json:"lastModifiedDate,omitempty" tf:"last_modified_date,omitempty"`

	LimitName *string `json:"limitName,omitempty" tf:"limit_name,omitempty"`

	MaximumLimit *float64 `json:"maximumLimit,omitempty" tf:"maximum_limit,omitempty"`

	OverrunPolicy *string `json:"overrunPolicy,omitempty" tf:"overrun_policy,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	TenantName *string `json:"tenantName,omitempty" tf:"tenant_name,omitempty"`

	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type QueryLimitParameters struct {

	// +kubebuilder:validation:Optional
	DefaultLimit *float64 `json:"defaultLimit,omitempty" tf:"default_limit,omitempty"`

	// +kubebuilder:validation:Optional
	LimitName *string `json:"limitName,omitempty" tf:"limit_name,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumLimit *float64 `json:"maximumLimit,omitempty" tf:"maximum_limit,omitempty"`

	// +kubebuilder:validation:Optional
	OverrunPolicy *string `json:"overrunPolicy,omitempty" tf:"overrun_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	TenantName *string `json:"tenantName,omitempty" tf:"tenant_name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

// QueryLimitSpec defines the desired state of QueryLimit
type QueryLimitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueryLimitParameters `json:"forProvider"`
}

// QueryLimitStatus defines the observed state of QueryLimit.
type QueryLimitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueryLimitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QueryLimit is the Schema for the QueryLimits API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type QueryLimit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.limitName)",message="limitName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.overrunPolicy)",message="overrunPolicy is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.projectId)",message="projectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.tenantName)",message="tenantName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.value)",message="value is a required parameter"
	Spec   QueryLimitSpec   `json:"spec"`
	Status QueryLimitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueryLimitList contains a list of QueryLimits
type QueryLimitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QueryLimit `json:"items"`
}

// Repository type metadata.
var (
	QueryLimit_Kind             = "QueryLimit"
	QueryLimit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QueryLimit_Kind}.String()
	QueryLimit_KindAPIVersion   = QueryLimit_Kind + "." + CRDGroupVersion.String()
	QueryLimit_GroupVersionKind = CRDGroupVersion.WithKind(QueryLimit_Kind)
)

func init() {
	SchemeBuilder.Register(&QueryLimit{}, &QueryLimitList{})
}
