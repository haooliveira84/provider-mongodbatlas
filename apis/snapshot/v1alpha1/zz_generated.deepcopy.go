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
func (in *BackupSnapshot) DeepCopyInto(out *BackupSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSnapshot.
func (in *BackupSnapshot) DeepCopy() *BackupSnapshot {
	if in == nil {
		return nil
	}
	out := new(BackupSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSnapshotList) DeepCopyInto(out *BackupSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSnapshotList.
func (in *BackupSnapshotList) DeepCopy() *BackupSnapshotList {
	if in == nil {
		return nil
	}
	out := new(BackupSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSnapshotObservation) DeepCopyInto(out *BackupSnapshotObservation) {
	*out = *in
	if in.CloudProvider != nil {
		in, out := &in.CloudProvider, &out.CloudProvider
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MasterKeyUUID != nil {
		in, out := &in.MasterKeyUUID, &out.MasterKeyUUID
		*out = new(string)
		**out = **in
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]MembersObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MongodVersion != nil {
		in, out := &in.MongodVersion, &out.MongodVersion
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ReplicaSetName != nil {
		in, out := &in.ReplicaSetName, &out.ReplicaSetName
		*out = new(string)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(float64)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
		*out = new(string)
		**out = **in
	}
	if in.SnapshotIds != nil {
		in, out := &in.SnapshotIds, &out.SnapshotIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SnapshotType != nil {
		in, out := &in.SnapshotType, &out.SnapshotType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StorageSizeBytes != nil {
		in, out := &in.StorageSizeBytes, &out.StorageSizeBytes
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSnapshotObservation.
func (in *BackupSnapshotObservation) DeepCopy() *BackupSnapshotObservation {
	if in == nil {
		return nil
	}
	out := new(BackupSnapshotObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSnapshotParameters) DeepCopyInto(out *BackupSnapshotParameters) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.RetentionInDays != nil {
		in, out := &in.RetentionInDays, &out.RetentionInDays
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSnapshotParameters.
func (in *BackupSnapshotParameters) DeepCopy() *BackupSnapshotParameters {
	if in == nil {
		return nil
	}
	out := new(BackupSnapshotParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSnapshotSpec) DeepCopyInto(out *BackupSnapshotSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSnapshotSpec.
func (in *BackupSnapshotSpec) DeepCopy() *BackupSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSnapshotStatus) DeepCopyInto(out *BackupSnapshotStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSnapshotStatus.
func (in *BackupSnapshotStatus) DeepCopy() *BackupSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(BackupSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembersObservation) DeepCopyInto(out *MembersObservation) {
	*out = *in
	if in.CloudProvider != nil {
		in, out := &in.CloudProvider, &out.CloudProvider
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ReplicaSetName != nil {
		in, out := &in.ReplicaSetName, &out.ReplicaSetName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembersObservation.
func (in *MembersObservation) DeepCopy() *MembersObservation {
	if in == nil {
		return nil
	}
	out := new(MembersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembersParameters) DeepCopyInto(out *MembersParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembersParameters.
func (in *MembersParameters) DeepCopy() *MembersParameters {
	if in == nil {
		return nil
	}
	out := new(MembersParameters)
	in.DeepCopyInto(out)
	return out
}
