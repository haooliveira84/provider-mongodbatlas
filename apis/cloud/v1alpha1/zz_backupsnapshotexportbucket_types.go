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

type BackupSnapshotExportBucketObservation struct {

	// Name of the bucket that the provided role ID is authorized to access. You must also specify the iam_role_id.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports AWS.
	CloudProvider *string `json:"cloudProvider,omitempty" tf:"cloud_provider,omitempty"`

	// Unique identifier of the snapshot export bucket.
	ExportBucketID *string `json:"exportBucketId,omitempty" tf:"export_bucket_id,omitempty"`

	// Unique identifier of the role that Atlas can use to access the bucket. You must also specify the bucket_name.
	IAMRoleID *string `json:"iamRoleId,omitempty" tf:"iam_role_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier of the project for the Atlas cluster.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type BackupSnapshotExportBucketParameters struct {

	// Name of the bucket that the provided role ID is authorized to access. You must also specify the iam_role_id.
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports AWS.
	// +kubebuilder:validation:Optional
	CloudProvider *string `json:"cloudProvider,omitempty" tf:"cloud_provider,omitempty"`

	// Unique identifier of the role that Atlas can use to access the bucket. You must also specify the bucket_name.
	// +kubebuilder:validation:Optional
	IAMRoleID *string `json:"iamRoleId,omitempty" tf:"iam_role_id,omitempty"`

	// The unique identifier of the project for the Atlas cluster.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

// BackupSnapshotExportBucketSpec defines the desired state of BackupSnapshotExportBucket
type BackupSnapshotExportBucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupSnapshotExportBucketParameters `json:"forProvider"`
}

// BackupSnapshotExportBucketStatus defines the observed state of BackupSnapshotExportBucket.
type BackupSnapshotExportBucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupSnapshotExportBucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupSnapshotExportBucket is the Schema for the BackupSnapshotExportBuckets API. Provides a Cloud Backup Snapshot Export Bucket resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type BackupSnapshotExportBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.bucketName)",message="bucketName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.cloudProvider)",message="cloudProvider is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.iamRoleId)",message="iamRoleId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.projectId)",message="projectId is a required parameter"
	Spec   BackupSnapshotExportBucketSpec   `json:"spec"`
	Status BackupSnapshotExportBucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupSnapshotExportBucketList contains a list of BackupSnapshotExportBuckets
type BackupSnapshotExportBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupSnapshotExportBucket `json:"items"`
}

// Repository type metadata.
var (
	BackupSnapshotExportBucket_Kind             = "BackupSnapshotExportBucket"
	BackupSnapshotExportBucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupSnapshotExportBucket_Kind}.String()
	BackupSnapshotExportBucket_KindAPIVersion   = BackupSnapshotExportBucket_Kind + "." + CRDGroupVersion.String()
	BackupSnapshotExportBucket_GroupVersionKind = CRDGroupVersion.WithKind(BackupSnapshotExportBucket_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupSnapshotExportBucket{}, &BackupSnapshotExportBucketList{})
}
