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

type PeeringObservation struct {

	// Specifies the AWS region where the peer VPC resides. For complete lists of supported regions, see Amazon Web Services.
	AccepterRegionName *string `json:"accepterRegionName,omitempty" tf:"accepter_region_name,omitempty"`

	AtlasCidrBlock *string `json:"atlasCidrBlock,omitempty" tf:"atlas_cidr_block,omitempty"`

	// The Atlas GCP Project ID for the GCP VPC used by your atlas cluster that it is need to set up the reciprocal connection.
	AtlasGCPProjectID *string `json:"atlasGcpProjectId,omitempty" tf:"atlas_gcp_project_id,omitempty"`

	AtlasID *string `json:"atlasId,omitempty" tf:"atlas_id,omitempty"`

	AtlasVPCName *string `json:"atlasVpcName,omitempty" tf:"atlas_vpc_name,omitempty"`

	// AWS Account ID of the owner of the peer VPC.
	AwsAccountID *string `json:"awsAccountId,omitempty" tf:"aws_account_id,omitempty"`

	// Unique identifier for an Azure AD directory.
	AzureDirectoryID *string `json:"azureDirectoryId,omitempty" tf:"azure_directory_id,omitempty"`

	// Unique identifier of the Azure subscription in which the VNet resides.
	AzureSubscriptionID *string `json:"azureSubscriptionId,omitempty" tf:"azure_subscription_id,omitempty"`

	// Unique identifier of the Atlas network peering container.
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Unique identifier of the MongoDB Atlas container for the provider (GCP) or provider/region (AWS, AZURE). You can create an MongoDB Atlas container using the network_container resource or it can be obtained from the cluster returned values if a cluster has been created before the first container.
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// When "status" : "FAILED", Atlas provides a description of the error.
	ErrorMessage *string `json:"errorMessage,omitempty" tf:"error_message,omitempty"`

	// Description of the Atlas error when status is Failed, Otherwise, Atlas returns null.
	ErrorState *string `json:"errorState,omitempty" tf:"error_state,omitempty"`

	// Error state, if any. The VPC peering connection error state value can be one of the following: REJECTED, EXPIRED, INVALID_ARGUMENT.
	ErrorStateName *string `json:"errorStateName,omitempty" tf:"error_state_name,omitempty"`

	// GCP project ID of the owner of the network peer.
	GCPProjectID *string `json:"gcpProjectId,omitempty" tf:"gcp_project_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the network peer to which Atlas connects.
	NetworkName *string `json:"networkName,omitempty" tf:"network_name,omitempty"`

	// Unique identifier of the Atlas network peer.
	PeerID *string `json:"peerId,omitempty" tf:"peer_id,omitempty"`

	// The unique ID for the MongoDB Atlas project to create the database user.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Cloud provider to whom the peering connection is being made. (Possible Values AWS, AZURE, GCP).
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Name of your Azure resource group.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// AWS VPC CIDR block or subnet.
	RouteTableCidrBlock *string `json:"routeTableCidrBlock,omitempty" tf:"route_table_cidr_block,omitempty"`

	// Status of the Atlas network peering connection.  Azure/GCP: ADDING_PEER, AVAILABLE, FAILED, DELETING GCP Only:  WAITING_FOR_USER.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (AWS Only) The VPC peering connection status value can be one of the following: INITIATING, PENDING_ACCEPTANCE, FAILED, FINALIZING, AVAILABLE, TERMINATING.
	StatusName *string `json:"statusName,omitempty" tf:"status_name,omitempty"`

	// Unique identifier of the AWS peer VPC (Note: this is not the same as the Atlas AWS VPC that is returned by the network_container resource).
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Name of your Azure VNet.
	VnetName *string `json:"vnetName,omitempty" tf:"vnet_name,omitempty"`
}

type PeeringParameters struct {

	// Specifies the AWS region where the peer VPC resides. For complete lists of supported regions, see Amazon Web Services.
	// +kubebuilder:validation:Optional
	AccepterRegionName *string `json:"accepterRegionName,omitempty" tf:"accepter_region_name,omitempty"`

	// +kubebuilder:validation:Optional
	AtlasCidrBlock *string `json:"atlasCidrBlock,omitempty" tf:"atlas_cidr_block,omitempty"`

	// The Atlas GCP Project ID for the GCP VPC used by your atlas cluster that it is need to set up the reciprocal connection.
	// +kubebuilder:validation:Optional
	AtlasGCPProjectID *string `json:"atlasGcpProjectId,omitempty" tf:"atlas_gcp_project_id,omitempty"`

	// +kubebuilder:validation:Optional
	AtlasVPCName *string `json:"atlasVpcName,omitempty" tf:"atlas_vpc_name,omitempty"`

	// AWS Account ID of the owner of the peer VPC.
	// +kubebuilder:validation:Optional
	AwsAccountID *string `json:"awsAccountId,omitempty" tf:"aws_account_id,omitempty"`

	// Unique identifier for an Azure AD directory.
	// +kubebuilder:validation:Optional
	AzureDirectoryID *string `json:"azureDirectoryId,omitempty" tf:"azure_directory_id,omitempty"`

	// Unique identifier of the Azure subscription in which the VNet resides.
	// +kubebuilder:validation:Optional
	AzureSubscriptionID *string `json:"azureSubscriptionId,omitempty" tf:"azure_subscription_id,omitempty"`

	// Unique identifier of the MongoDB Atlas container for the provider (GCP) or provider/region (AWS, AZURE). You can create an MongoDB Atlas container using the network_container resource or it can be obtained from the cluster returned values if a cluster has been created before the first container.
	// +kubebuilder:validation:Optional
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// GCP project ID of the owner of the network peer.
	// +kubebuilder:validation:Optional
	GCPProjectID *string `json:"gcpProjectId,omitempty" tf:"gcp_project_id,omitempty"`

	// Name of the network peer to which Atlas connects.
	// +kubebuilder:validation:Optional
	NetworkName *string `json:"networkName,omitempty" tf:"network_name,omitempty"`

	// The unique ID for the MongoDB Atlas project to create the database user.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Cloud provider to whom the peering connection is being made. (Possible Values AWS, AZURE, GCP).
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Name of your Azure resource group.
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// AWS VPC CIDR block or subnet.
	// +kubebuilder:validation:Optional
	RouteTableCidrBlock *string `json:"routeTableCidrBlock,omitempty" tf:"route_table_cidr_block,omitempty"`

	// Unique identifier of the AWS peer VPC (Note: this is not the same as the Atlas AWS VPC that is returned by the network_container resource).
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Name of your Azure VNet.
	// +kubebuilder:validation:Optional
	VnetName *string `json:"vnetName,omitempty" tf:"vnet_name,omitempty"`
}

// PeeringSpec defines the desired state of Peering
type PeeringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PeeringParameters `json:"forProvider"`
}

// PeeringStatus defines the observed state of Peering.
type PeeringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PeeringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Peering is the Schema for the Peerings API. Provides a Network Peering resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type Peering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.containerId)",message="containerId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.projectId)",message="projectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.providerName)",message="providerName is a required parameter"
	Spec   PeeringSpec   `json:"spec"`
	Status PeeringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringList contains a list of Peerings
type PeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Peering `json:"items"`
}

// Repository type metadata.
var (
	Peering_Kind             = "Peering"
	Peering_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Peering_Kind}.String()
	Peering_KindAPIVersion   = Peering_Kind + "." + CRDGroupVersion.String()
	Peering_GroupVersionKind = CRDGroupVersion.WithKind(Peering_Kind)
)

func init() {
	SchemeBuilder.Register(&Peering{}, &PeeringList{})
}