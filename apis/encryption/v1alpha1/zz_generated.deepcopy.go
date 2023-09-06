//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtRest) DeepCopyInto(out *AtRest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtRest.
func (in *AtRest) DeepCopy() *AtRest {
	if in == nil {
		return nil
	}
	out := new(AtRest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AtRest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtRestList) DeepCopyInto(out *AtRestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AtRest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtRestList.
func (in *AtRestList) DeepCopy() *AtRestList {
	if in == nil {
		return nil
	}
	out := new(AtRestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AtRestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtRestObservation) DeepCopyInto(out *AtRestObservation) {
	*out = *in
	if in.AwsKMSConfig != nil {
		in, out := &in.AwsKMSConfig, &out.AwsKMSConfig
		*out = make([]AwsKMSConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AzureKeyVaultConfig != nil {
		in, out := &in.AzureKeyVaultConfig, &out.AzureKeyVaultConfig
		*out = make([]AzureKeyVaultConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GoogleCloudKMSConfig != nil {
		in, out := &in.GoogleCloudKMSConfig, &out.GoogleCloudKMSConfig
		*out = make([]GoogleCloudKMSConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtRestObservation.
func (in *AtRestObservation) DeepCopy() *AtRestObservation {
	if in == nil {
		return nil
	}
	out := new(AtRestObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtRestParameters) DeepCopyInto(out *AtRestParameters) {
	*out = *in
	if in.AwsKMSConfig != nil {
		in, out := &in.AwsKMSConfig, &out.AwsKMSConfig
		*out = make([]AwsKMSConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AwsKMSSecretRef != nil {
		in, out := &in.AwsKMSSecretRef, &out.AwsKMSSecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.AzureKeyVaultConfig != nil {
		in, out := &in.AzureKeyVaultConfig, &out.AzureKeyVaultConfig
		*out = make([]AzureKeyVaultConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AzureKeyVaultSecretRef != nil {
		in, out := &in.AzureKeyVaultSecretRef, &out.AzureKeyVaultSecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.GoogleCloudKMSConfig != nil {
		in, out := &in.GoogleCloudKMSConfig, &out.GoogleCloudKMSConfig
		*out = make([]GoogleCloudKMSConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GoogleCloudKMSSecretRef != nil {
		in, out := &in.GoogleCloudKMSSecretRef, &out.GoogleCloudKMSSecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtRestParameters.
func (in *AtRestParameters) DeepCopy() *AtRestParameters {
	if in == nil {
		return nil
	}
	out := new(AtRestParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtRestSpec) DeepCopyInto(out *AtRestSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtRestSpec.
func (in *AtRestSpec) DeepCopy() *AtRestSpec {
	if in == nil {
		return nil
	}
	out := new(AtRestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtRestStatus) DeepCopyInto(out *AtRestStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtRestStatus.
func (in *AtRestStatus) DeepCopy() *AtRestStatus {
	if in == nil {
		return nil
	}
	out := new(AtRestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsKMSConfigObservation) DeepCopyInto(out *AwsKMSConfigObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsKMSConfigObservation.
func (in *AwsKMSConfigObservation) DeepCopy() *AwsKMSConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AwsKMSConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsKMSConfigParameters) DeepCopyInto(out *AwsKMSConfigParameters) {
	*out = *in
	if in.AccessKeyIDSecretRef != nil {
		in, out := &in.AccessKeyIDSecretRef, &out.AccessKeyIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.CustomerMasterKeyIDSecretRef != nil {
		in, out := &in.CustomerMasterKeyIDSecretRef, &out.CustomerMasterKeyIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	if in.SecretAccessKeySecretRef != nil {
		in, out := &in.SecretAccessKeySecretRef, &out.SecretAccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsKMSConfigParameters.
func (in *AwsKMSConfigParameters) DeepCopy() *AwsKMSConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AwsKMSConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultConfigObservation) DeepCopyInto(out *AzureKeyVaultConfigObservation) {
	*out = *in
	if in.AzureEnvironment != nil {
		in, out := &in.AzureEnvironment, &out.AzureEnvironment
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.KeyVaultName != nil {
		in, out := &in.KeyVaultName, &out.KeyVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultConfigObservation.
func (in *AzureKeyVaultConfigObservation) DeepCopy() *AzureKeyVaultConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultConfigParameters) DeepCopyInto(out *AzureKeyVaultConfigParameters) {
	*out = *in
	if in.AzureEnvironment != nil {
		in, out := &in.AzureEnvironment, &out.AzureEnvironment
		*out = new(string)
		**out = **in
	}
	if in.ClientIDSecretRef != nil {
		in, out := &in.ClientIDSecretRef, &out.ClientIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.KeyIdentifierSecretRef != nil {
		in, out := &in.KeyIdentifierSecretRef, &out.KeyIdentifierSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.KeyVaultName != nil {
		in, out := &in.KeyVaultName, &out.KeyVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SecretSecretRef != nil {
		in, out := &in.SecretSecretRef, &out.SecretSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SubscriptionIDSecretRef != nil {
		in, out := &in.SubscriptionIDSecretRef, &out.SubscriptionIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.TenantIDSecretRef != nil {
		in, out := &in.TenantIDSecretRef, &out.TenantIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultConfigParameters.
func (in *AzureKeyVaultConfigParameters) DeepCopy() *AzureKeyVaultConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleCloudKMSConfigObservation) DeepCopyInto(out *GoogleCloudKMSConfigObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleCloudKMSConfigObservation.
func (in *GoogleCloudKMSConfigObservation) DeepCopy() *GoogleCloudKMSConfigObservation {
	if in == nil {
		return nil
	}
	out := new(GoogleCloudKMSConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleCloudKMSConfigParameters) DeepCopyInto(out *GoogleCloudKMSConfigParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.KeyVersionResourceIDSecretRef != nil {
		in, out := &in.KeyVersionResourceIDSecretRef, &out.KeyVersionResourceIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ServiceAccountKeySecretRef != nil {
		in, out := &in.ServiceAccountKeySecretRef, &out.ServiceAccountKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleCloudKMSConfigParameters.
func (in *GoogleCloudKMSConfigParameters) DeepCopy() *GoogleCloudKMSConfigParameters {
	if in == nil {
		return nil
	}
	out := new(GoogleCloudKMSConfigParameters)
	in.DeepCopyInto(out)
	return out
}
