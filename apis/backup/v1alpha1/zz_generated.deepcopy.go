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
func (in *CompliancePolicy) DeepCopyInto(out *CompliancePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompliancePolicy.
func (in *CompliancePolicy) DeepCopy() *CompliancePolicy {
	if in == nil {
		return nil
	}
	out := new(CompliancePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompliancePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompliancePolicyList) DeepCopyInto(out *CompliancePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CompliancePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompliancePolicyList.
func (in *CompliancePolicyList) DeepCopy() *CompliancePolicyList {
	if in == nil {
		return nil
	}
	out := new(CompliancePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompliancePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompliancePolicyObservation) DeepCopyInto(out *CompliancePolicyObservation) {
	*out = *in
	if in.AuthorizedEmail != nil {
		in, out := &in.AuthorizedEmail, &out.AuthorizedEmail
		*out = new(string)
		**out = **in
	}
	if in.CopyProtectionEnabled != nil {
		in, out := &in.CopyProtectionEnabled, &out.CopyProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionAtRestEnabled != nil {
		in, out := &in.EncryptionAtRestEnabled, &out.EncryptionAtRestEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.OnDemandPolicyItem != nil {
		in, out := &in.OnDemandPolicyItem, &out.OnDemandPolicyItem
		*out = make([]OnDemandPolicyItemObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PitEnabled != nil {
		in, out := &in.PitEnabled, &out.PitEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PolicyItemDaily != nil {
		in, out := &in.PolicyItemDaily, &out.PolicyItemDaily
		*out = make([]PolicyItemDailyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyItemHourly != nil {
		in, out := &in.PolicyItemHourly, &out.PolicyItemHourly
		*out = make([]PolicyItemHourlyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyItemMonthly != nil {
		in, out := &in.PolicyItemMonthly, &out.PolicyItemMonthly
		*out = make([]PolicyItemMonthlyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyItemWeekly != nil {
		in, out := &in.PolicyItemWeekly, &out.PolicyItemWeekly
		*out = make([]PolicyItemWeeklyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.RestoreWindowDays != nil {
		in, out := &in.RestoreWindowDays, &out.RestoreWindowDays
		*out = new(float64)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.UpdatedDate != nil {
		in, out := &in.UpdatedDate, &out.UpdatedDate
		*out = new(string)
		**out = **in
	}
	if in.UpdatedUser != nil {
		in, out := &in.UpdatedUser, &out.UpdatedUser
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompliancePolicyObservation.
func (in *CompliancePolicyObservation) DeepCopy() *CompliancePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(CompliancePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompliancePolicyParameters) DeepCopyInto(out *CompliancePolicyParameters) {
	*out = *in
	if in.AuthorizedEmail != nil {
		in, out := &in.AuthorizedEmail, &out.AuthorizedEmail
		*out = new(string)
		**out = **in
	}
	if in.CopyProtectionEnabled != nil {
		in, out := &in.CopyProtectionEnabled, &out.CopyProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionAtRestEnabled != nil {
		in, out := &in.EncryptionAtRestEnabled, &out.EncryptionAtRestEnabled
		*out = new(bool)
		**out = **in
	}
	if in.OnDemandPolicyItem != nil {
		in, out := &in.OnDemandPolicyItem, &out.OnDemandPolicyItem
		*out = make([]OnDemandPolicyItemParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PitEnabled != nil {
		in, out := &in.PitEnabled, &out.PitEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PolicyItemDaily != nil {
		in, out := &in.PolicyItemDaily, &out.PolicyItemDaily
		*out = make([]PolicyItemDailyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyItemHourly != nil {
		in, out := &in.PolicyItemHourly, &out.PolicyItemHourly
		*out = make([]PolicyItemHourlyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyItemMonthly != nil {
		in, out := &in.PolicyItemMonthly, &out.PolicyItemMonthly
		*out = make([]PolicyItemMonthlyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyItemWeekly != nil {
		in, out := &in.PolicyItemWeekly, &out.PolicyItemWeekly
		*out = make([]PolicyItemWeeklyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.RestoreWindowDays != nil {
		in, out := &in.RestoreWindowDays, &out.RestoreWindowDays
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompliancePolicyParameters.
func (in *CompliancePolicyParameters) DeepCopy() *CompliancePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(CompliancePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompliancePolicySpec) DeepCopyInto(out *CompliancePolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompliancePolicySpec.
func (in *CompliancePolicySpec) DeepCopy() *CompliancePolicySpec {
	if in == nil {
		return nil
	}
	out := new(CompliancePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompliancePolicyStatus) DeepCopyInto(out *CompliancePolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompliancePolicyStatus.
func (in *CompliancePolicyStatus) DeepCopy() *CompliancePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(CompliancePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnDemandPolicyItemObservation) DeepCopyInto(out *OnDemandPolicyItemObservation) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.FrequencyType != nil {
		in, out := &in.FrequencyType, &out.FrequencyType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnDemandPolicyItemObservation.
func (in *OnDemandPolicyItemObservation) DeepCopy() *OnDemandPolicyItemObservation {
	if in == nil {
		return nil
	}
	out := new(OnDemandPolicyItemObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnDemandPolicyItemParameters) DeepCopyInto(out *OnDemandPolicyItemParameters) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnDemandPolicyItemParameters.
func (in *OnDemandPolicyItemParameters) DeepCopy() *OnDemandPolicyItemParameters {
	if in == nil {
		return nil
	}
	out := new(OnDemandPolicyItemParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyItemDailyObservation) DeepCopyInto(out *PolicyItemDailyObservation) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.FrequencyType != nil {
		in, out := &in.FrequencyType, &out.FrequencyType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyItemDailyObservation.
func (in *PolicyItemDailyObservation) DeepCopy() *PolicyItemDailyObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyItemDailyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyItemDailyParameters) DeepCopyInto(out *PolicyItemDailyParameters) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyItemDailyParameters.
func (in *PolicyItemDailyParameters) DeepCopy() *PolicyItemDailyParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyItemDailyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyItemHourlyObservation) DeepCopyInto(out *PolicyItemHourlyObservation) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.FrequencyType != nil {
		in, out := &in.FrequencyType, &out.FrequencyType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyItemHourlyObservation.
func (in *PolicyItemHourlyObservation) DeepCopy() *PolicyItemHourlyObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyItemHourlyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyItemHourlyParameters) DeepCopyInto(out *PolicyItemHourlyParameters) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyItemHourlyParameters.
func (in *PolicyItemHourlyParameters) DeepCopy() *PolicyItemHourlyParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyItemHourlyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyItemMonthlyObservation) DeepCopyInto(out *PolicyItemMonthlyObservation) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.FrequencyType != nil {
		in, out := &in.FrequencyType, &out.FrequencyType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyItemMonthlyObservation.
func (in *PolicyItemMonthlyObservation) DeepCopy() *PolicyItemMonthlyObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyItemMonthlyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyItemMonthlyParameters) DeepCopyInto(out *PolicyItemMonthlyParameters) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyItemMonthlyParameters.
func (in *PolicyItemMonthlyParameters) DeepCopy() *PolicyItemMonthlyParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyItemMonthlyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyItemWeeklyObservation) DeepCopyInto(out *PolicyItemWeeklyObservation) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.FrequencyType != nil {
		in, out := &in.FrequencyType, &out.FrequencyType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyItemWeeklyObservation.
func (in *PolicyItemWeeklyObservation) DeepCopy() *PolicyItemWeeklyObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyItemWeeklyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyItemWeeklyParameters) DeepCopyInto(out *PolicyItemWeeklyParameters) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyItemWeeklyParameters.
func (in *PolicyItemWeeklyParameters) DeepCopy() *PolicyItemWeeklyParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyItemWeeklyParameters)
	in.DeepCopyInto(out)
	return out
}
