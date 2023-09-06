//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigList) DeepCopyInto(out *ClusterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigList.
func (in *ClusterConfigList) DeepCopy() *ClusterConfigList {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigObservation) DeepCopyInto(out *ClusterConfigObservation) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CustomZoneMapping != nil {
		in, out := &in.CustomZoneMapping, &out.CustomZoneMapping
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.CustomZoneMappings != nil {
		in, out := &in.CustomZoneMappings, &out.CustomZoneMappings
		*out = make([]CustomZoneMappingsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ManagedNamespaces != nil {
		in, out := &in.ManagedNamespaces, &out.ManagedNamespaces
		*out = make([]ManagedNamespacesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigObservation.
func (in *ClusterConfigObservation) DeepCopy() *ClusterConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigParameters) DeepCopyInto(out *ClusterConfigParameters) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CustomZoneMappings != nil {
		in, out := &in.CustomZoneMappings, &out.CustomZoneMappings
		*out = make([]CustomZoneMappingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManagedNamespaces != nil {
		in, out := &in.ManagedNamespaces, &out.ManagedNamespaces
		*out = make([]ManagedNamespacesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigParameters.
func (in *ClusterConfigParameters) DeepCopy() *ClusterConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigSpec) DeepCopyInto(out *ClusterConfigSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigSpec.
func (in *ClusterConfigSpec) DeepCopy() *ClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigStatus) DeepCopyInto(out *ClusterConfigStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigStatus.
func (in *ClusterConfigStatus) DeepCopy() *ClusterConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomZoneMappingsObservation) DeepCopyInto(out *CustomZoneMappingsObservation) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomZoneMappingsObservation.
func (in *CustomZoneMappingsObservation) DeepCopy() *CustomZoneMappingsObservation {
	if in == nil {
		return nil
	}
	out := new(CustomZoneMappingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomZoneMappingsParameters) DeepCopyInto(out *CustomZoneMappingsParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomZoneMappingsParameters.
func (in *CustomZoneMappingsParameters) DeepCopy() *CustomZoneMappingsParameters {
	if in == nil {
		return nil
	}
	out := new(CustomZoneMappingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedNamespacesObservation) DeepCopyInto(out *ManagedNamespacesObservation) {
	*out = *in
	if in.Collection != nil {
		in, out := &in.Collection, &out.Collection
		*out = new(string)
		**out = **in
	}
	if in.CustomShardKey != nil {
		in, out := &in.CustomShardKey, &out.CustomShardKey
		*out = new(string)
		**out = **in
	}
	if in.DB != nil {
		in, out := &in.DB, &out.DB
		*out = new(string)
		**out = **in
	}
	if in.IsCustomShardKeyHashed != nil {
		in, out := &in.IsCustomShardKeyHashed, &out.IsCustomShardKeyHashed
		*out = new(bool)
		**out = **in
	}
	if in.IsShardKeyUnique != nil {
		in, out := &in.IsShardKeyUnique, &out.IsShardKeyUnique
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedNamespacesObservation.
func (in *ManagedNamespacesObservation) DeepCopy() *ManagedNamespacesObservation {
	if in == nil {
		return nil
	}
	out := new(ManagedNamespacesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedNamespacesParameters) DeepCopyInto(out *ManagedNamespacesParameters) {
	*out = *in
	if in.Collection != nil {
		in, out := &in.Collection, &out.Collection
		*out = new(string)
		**out = **in
	}
	if in.CustomShardKey != nil {
		in, out := &in.CustomShardKey, &out.CustomShardKey
		*out = new(string)
		**out = **in
	}
	if in.DB != nil {
		in, out := &in.DB, &out.DB
		*out = new(string)
		**out = **in
	}
	if in.IsCustomShardKeyHashed != nil {
		in, out := &in.IsCustomShardKeyHashed, &out.IsCustomShardKeyHashed
		*out = new(bool)
		**out = **in
	}
	if in.IsShardKeyUnique != nil {
		in, out := &in.IsShardKeyUnique, &out.IsShardKeyUnique
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedNamespacesParameters.
func (in *ManagedNamespacesParameters) DeepCopy() *ManagedNamespacesParameters {
	if in == nil {
		return nil
	}
	out := new(ManagedNamespacesParameters)
	in.DeepCopyInto(out)
	return out
}
