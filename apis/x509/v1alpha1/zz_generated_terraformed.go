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

// GetTerraformResourceType returns Terraform resource type for this AuthenticationDatabaseUser
func (mg *AuthenticationDatabaseUser) GetTerraformResourceType() string {
	return "mongodbatlas_x509_authentication_database_user"
}

// GetConnectionDetailsMapping for this AuthenticationDatabaseUser
func (tr *AuthenticationDatabaseUser) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"current_certificate": "status.atProvider.currentCertificate", "customer_x509_cas": "spec.forProvider.customerX509CasSecretRef"}
}

// GetObservation of this AuthenticationDatabaseUser
func (tr *AuthenticationDatabaseUser) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AuthenticationDatabaseUser
func (tr *AuthenticationDatabaseUser) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AuthenticationDatabaseUser
func (tr *AuthenticationDatabaseUser) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AuthenticationDatabaseUser
func (tr *AuthenticationDatabaseUser) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AuthenticationDatabaseUser
func (tr *AuthenticationDatabaseUser) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AuthenticationDatabaseUser using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AuthenticationDatabaseUser) LateInitialize(attrs []byte) (bool, error) {
	params := &AuthenticationDatabaseUserParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AuthenticationDatabaseUser) GetTerraformSchemaVersion() int {
	return 0
}
