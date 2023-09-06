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

type IngestionSchedulesObservation struct {

	// Number that indicates the frequency interval for a set of snapshots.
	FrequencyInterval *float64 `json:"frequencyInterval,omitempty" tf:"frequency_interval,omitempty"`

	// Human-readable label that identifies how often this snapshot triggers.
	FrequencyType *string `json:"frequencyType,omitempty" tf:"frequency_type,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the snapshot.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unit of time in which MongoDB Atlas measures snapshot retention.
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`

	// Duration in days, weeks, or months that MongoDB Atlas retains the snapshot.
	RetentionValue *float64 `json:"retentionValue,omitempty" tf:"retention_value,omitempty"`
}

type IngestionSchedulesParameters struct {
}

type LakePipelineObservation struct {

	// Timestamp that indicates when the Data Lake Pipeline was created.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the Data Lake Pipeline.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of backup schedule policy items that you can use as a Data Lake Pipeline source.
	IngestionSchedules []IngestionSchedulesObservation `json:"ingestionSchedules,omitempty" tf:"ingestion_schedules,omitempty"`

	// Timestamp that indicates the last time that the Data Lake Pipeline was updated.
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`

	// The unique ID for the project to create a data lake pipeline.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	Sink []SinkObservation `json:"sink,omitempty" tf:"sink,omitempty"`

	// List of backup snapshots that you can use to trigger an on demand pipeline run.
	Snapshots []SnapshotsObservation `json:"snapshots,omitempty" tf:"snapshots,omitempty"`

	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`

	// State of this Data Lake Pipeline.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Fields to be excluded for this Data Lake Pipeline.
	Transformations []TransformationsObservation `json:"transformations,omitempty" tf:"transformations,omitempty"`
}

type LakePipelineParameters struct {

	// The unique ID for the project to create a data lake pipeline.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	Sink []SinkParameters `json:"sink,omitempty" tf:"sink,omitempty"`

	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// Fields to be excluded for this Data Lake Pipeline.
	// +kubebuilder:validation:Optional
	Transformations []TransformationsParameters `json:"transformations,omitempty" tf:"transformations,omitempty"`
}

type PartitionFieldsObservation struct {

	// Human-readable label that identifies the field name used to partition data.
	FieldName *string `json:"fieldName,omitempty" tf:"field_name,omitempty"`

	// Sequence in which MongoDB Atlas slices the collection data to create partitions. The resource expresses this sequence starting with zero.
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type PartitionFieldsParameters struct {

	// Human-readable label that identifies the field name used to partition data.
	// +kubebuilder:validation:Required
	FieldName *string `json:"fieldName" tf:"field_name,omitempty"`

	// Sequence in which MongoDB Atlas slices the collection data to create partitions. The resource expresses this sequence starting with zero.
	// +kubebuilder:validation:Required
	Order *float64 `json:"order" tf:"order,omitempty"`
}

type SinkObservation struct {

	// Ordered fields used to physically organize data in the destination.
	PartitionFields []PartitionFieldsObservation `json:"partitionFields,omitempty" tf:"partition_fields,omitempty"`

	// Human-readable label that identifies the cloud provider that stores this snapshot.
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`

	// Target cloud provider region for this Data Lake Pipeline. Supported cloud provider regions.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Type of ingestion destination of this Data Lake Pipeline.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SinkParameters struct {

	// Ordered fields used to physically organize data in the destination.
	// +kubebuilder:validation:Optional
	PartitionFields []PartitionFieldsParameters `json:"partitionFields,omitempty" tf:"partition_fields,omitempty"`

	// Human-readable label that identifies the cloud provider that stores this snapshot.
	// +kubebuilder:validation:Optional
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`

	// Target cloud provider region for this Data Lake Pipeline. Supported cloud provider regions.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Type of ingestion destination of this Data Lake Pipeline.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SnapshotsObservation struct {

	// List that identifies the regions to which MongoDB Atlas copies the snapshot.
	CopyRegion *string `json:"copyRegion,omitempty" tf:"copy_region,omitempty"`

	// Date and time when MongoDB Atlas took the snapshot.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Date and time when MongoDB Atlas deletes the snapshot.
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	FrequencyYype *string `json:"frequencyYype,omitempty" tf:"frequency_yype,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the snapshot.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique string that identifies the Amazon Web Services (AWS) Key Management Service (KMS) Customer Master Key (CMK) used to encrypt the snapshot.
	MasterKey *string `json:"masterKey,omitempty" tf:"master_key,omitempty"`

	// Version of the MongoDB host that this snapshot backs up.
	MongodVersion *string `json:"mongodVersion,omitempty" tf:"mongod_version,omitempty"`

	// List that contains unique identifiers for the policy items.
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// Human-readable label that identifies the cloud provider that stores this snapshot.
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`

	// Human-readable label that identifies the replica set from which MongoDB Atlas took this snapshot.
	ReplicaSetName *string `json:"replicaSetName,omitempty" tf:"replica_set_name,omitempty"`

	// List of backup snapshots that you can use to trigger an on demand pipeline run.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Human-readable label that identifies when this snapshot triggers.
	SnapshotType *string `json:"snapshotType,omitempty" tf:"snapshot_type,omitempty"`

	// Human-readable label that indicates the stage of the backup process for this snapshot.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Type of ingestion destination of this Data Lake Pipeline.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SnapshotsParameters struct {
}

type SourceObservation struct {

	// Human-readable name that identifies the cluster.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Human-readable name that identifies the collection.
	CollectionName *string `json:"collectionName,omitempty" tf:"collection_name,omitempty"`

	// Human-readable name that identifies the database.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the Data Lake Pipeline.
	PolicyItemID *string `json:"policyItemId,omitempty" tf:"policy_item_id,omitempty"`

	// The unique ID for the project to create a data lake pipeline.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Type of ingestion destination of this Data Lake Pipeline.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SourceParameters struct {

	// Human-readable name that identifies the cluster.
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Human-readable name that identifies the collection.
	// +kubebuilder:validation:Optional
	CollectionName *string `json:"collectionName,omitempty" tf:"collection_name,omitempty"`

	// Human-readable name that identifies the database.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the Data Lake Pipeline.
	// +kubebuilder:validation:Optional
	PolicyItemID *string `json:"policyItemId,omitempty" tf:"policy_item_id,omitempty"`

	// The unique ID for the project to create a data lake pipeline.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Type of ingestion destination of this Data Lake Pipeline.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TransformationsObservation struct {

	// Key in the document.
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// Type of ingestion destination of this Data Lake Pipeline.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TransformationsParameters struct {

	// Key in the document.
	// +kubebuilder:validation:Optional
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// Type of ingestion destination of this Data Lake Pipeline.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// LakePipelineSpec defines the desired state of LakePipeline
type LakePipelineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LakePipelineParameters `json:"forProvider"`
}

// LakePipelineStatus defines the observed state of LakePipeline.
type LakePipelineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LakePipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LakePipeline is the Schema for the LakePipelines API. Provides a Data Lake Pipeline resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type LakePipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.projectId)",message="projectId is a required parameter"
	Spec   LakePipelineSpec   `json:"spec"`
	Status LakePipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LakePipelineList contains a list of LakePipelines
type LakePipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LakePipeline `json:"items"`
}

// Repository type metadata.
var (
	LakePipeline_Kind             = "LakePipeline"
	LakePipeline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LakePipeline_Kind}.String()
	LakePipeline_KindAPIVersion   = LakePipeline_Kind + "." + CRDGroupVersion.String()
	LakePipeline_GroupVersionKind = CRDGroupVersion.WithKind(LakePipeline_Kind)
)

func init() {
	SchemeBuilder.Register(&LakePipeline{}, &LakePipelineList{})
}