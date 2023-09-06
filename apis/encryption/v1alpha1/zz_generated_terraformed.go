/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this AtRest
func (mg *AtRest) GetTerraformResourceType() string {
	return "mongodbatlas_encryption_at_rest"
}

// GetConnectionDetailsMapping for this AtRest
func (tr *AtRest) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"aws_kms": "spec.forProvider.awsKmsSecretRef", "aws_kms_config[*].access_key_id": "spec.forProvider.awsKmsConfig[*].accessKeyIdSecretRef", "aws_kms_config[*].customer_master_key_id": "spec.forProvider.awsKmsConfig[*].customerMasterKeyIdSecretRef", "aws_kms_config[*].secret_access_key": "spec.forProvider.awsKmsConfig[*].secretAccessKeySecretRef", "azure_key_vault": "spec.forProvider.azureKeyVaultSecretRef", "azure_key_vault_config[*].client_id": "spec.forProvider.azureKeyVaultConfig[*].clientIdSecretRef", "azure_key_vault_config[*].key_identifier": "spec.forProvider.azureKeyVaultConfig[*].keyIdentifierSecretRef", "azure_key_vault_config[*].secret": "spec.forProvider.azureKeyVaultConfig[*].secretSecretRef", "azure_key_vault_config[*].subscription_id": "spec.forProvider.azureKeyVaultConfig[*].subscriptionIdSecretRef", "azure_key_vault_config[*].tenant_id": "spec.forProvider.azureKeyVaultConfig[*].tenantIdSecretRef", "google_cloud_kms": "spec.forProvider.googleCloudKmsSecretRef", "google_cloud_kms_config[*].key_version_resource_id": "spec.forProvider.googleCloudKmsConfig[*].keyVersionResourceIdSecretRef", "google_cloud_kms_config[*].service_account_key": "spec.forProvider.googleCloudKmsConfig[*].serviceAccountKeySecretRef"}
}

// GetObservation of this AtRest
func (tr *AtRest) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AtRest
func (tr *AtRest) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AtRest
func (tr *AtRest) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AtRest
func (tr *AtRest) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AtRest
func (tr *AtRest) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AtRest using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AtRest) LateInitialize(attrs []byte) (bool, error) {
	params := &AtRestParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AtRest) GetTerraformSchemaVersion() int {
	return 0
}