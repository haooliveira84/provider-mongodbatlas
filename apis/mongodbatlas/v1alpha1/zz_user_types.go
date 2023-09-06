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

type LabelsObservation struct {

	// The key that you want to write.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value that you want to write.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type LabelsParameters struct {

	// The key that you want to write.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value that you want to write.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RolesObservation struct {

	// Collection for which the role applies. You can specify a collection for the read and readWrite roles. If you do not specify a collection for read and readWrite, the role applies to all collections in the database (excluding some collections in the system. database).
	CollectionName *string `json:"collectionName,omitempty" tf:"collection_name,omitempty"`

	// Database on which the user has the specified role. A role on the admin database can include privileges that apply to the other databases.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Name of the role to grant. See Create a Database User roles.roleName for valid values and restrictions.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`
}

type RolesParameters struct {

	// Collection for which the role applies. You can specify a collection for the read and readWrite roles. If you do not specify a collection for read and readWrite, the role applies to all collections in the database (excluding some collections in the system. database).
	// +kubebuilder:validation:Optional
	CollectionName *string `json:"collectionName,omitempty" tf:"collection_name,omitempty"`

	// Database on which the user has the specified role. A role on the admin database can include privileges that apply to the other databases.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Name of the role to grant. See Create a Database User roles.roleName for valid values and restrictions.
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`
}

type ScopesObservation struct {

	// Name of the cluster or Atlas Data Lake that the user has access to.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of resource that the user has access to. Valid values are: CLUSTER and DATA_LAKE
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ScopesParameters struct {

	// Name of the cluster or Atlas Data Lake that the user has access to.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of resource that the user has access to. Valid values are: CLUSTER and DATA_LAKE
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type UserObservation struct {

	// Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB.
	// Accepted values include:
	AuthDatabaseName *string `json:"authDatabaseName,omitempty" tf:"auth_database_name,omitempty"`

	// If this value is set, the new database user authenticates with AWS IAM credentials. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	AwsIAMType *string `json:"awsIamType,omitempty" tf:"aws_iam_type,omitempty"`

	// Database on which the user has the specified role. A role on the admin database can include privileges that apply to the other databases.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// The database user's name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Labels []LabelsObservation `json:"labels,omitempty" tf:"labels,omitempty"`

	// Method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE.
	LdapAuthType *string `json:"ldapAuthType,omitempty" tf:"ldap_auth_type,omitempty"`

	// Human-readable label that indicates whether the new database user authenticates with OIDC (OpenID Connect) federated authentication. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	OidcAuthType *string `json:"oidcAuthType,omitempty" tf:"oidc_auth_type,omitempty"`

	// The unique ID for the project to create the database user.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
	Roles []RolesObservation `json:"roles,omitempty" tf:"roles,omitempty"`

	Scopes []ScopesObservation `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// Username for authenticating to MongoDB. USER_ARN or ROLE_ARN if aws_iam_type is USER or ROLE.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// X.509 method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	X509Type *string `json:"x509Type,omitempty" tf:"x509_type,omitempty"`
}

type UserParameters struct {

	// Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB.
	// Accepted values include:
	// +kubebuilder:validation:Optional
	AuthDatabaseName *string `json:"authDatabaseName,omitempty" tf:"auth_database_name,omitempty"`

	// If this value is set, the new database user authenticates with AWS IAM credentials. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	// +kubebuilder:validation:Optional
	AwsIAMType *string `json:"awsIamType,omitempty" tf:"aws_iam_type,omitempty"`

	// Database on which the user has the specified role. A role on the admin database can include privileges that apply to the other databases.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	Labels []LabelsParameters `json:"labels,omitempty" tf:"labels,omitempty"`

	// Method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE.
	// +kubebuilder:validation:Optional
	LdapAuthType *string `json:"ldapAuthType,omitempty" tf:"ldap_auth_type,omitempty"`

	// Human-readable label that indicates whether the new database user authenticates with OIDC (OpenID Connect) federated authentication. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	// +kubebuilder:validation:Optional
	OidcAuthType *string `json:"oidcAuthType,omitempty" tf:"oidc_auth_type,omitempty"`

	// User's initial password. Password can be changed after creation using your preferred method, e.g. via the MongoDB Atlas UI, to ensure security.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The unique ID for the project to create the database user.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
	// +kubebuilder:validation:Optional
	Roles []RolesParameters `json:"roles,omitempty" tf:"roles,omitempty"`

	// +kubebuilder:validation:Optional
	Scopes []ScopesParameters `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// Username for authenticating to MongoDB. USER_ARN or ROLE_ARN if aws_iam_type is USER or ROLE.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// X.509 method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	// +kubebuilder:validation:Optional
	X509Type *string `json:"x509Type,omitempty" tf:"x509_type,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API. Provides a Database User resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.projectId)",message="projectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.roles)",message="roles is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.username)",message="username is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}