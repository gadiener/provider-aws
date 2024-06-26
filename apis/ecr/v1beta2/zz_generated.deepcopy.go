//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationInitParameters) DeepCopyInto(out *DestinationInitParameters) {
	*out = *in
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationInitParameters.
func (in *DestinationInitParameters) DeepCopy() *DestinationInitParameters {
	if in == nil {
		return nil
	}
	out := new(DestinationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationObservation) DeepCopyInto(out *DestinationObservation) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationObservation.
func (in *DestinationObservation) DeepCopy() *DestinationObservation {
	if in == nil {
		return nil
	}
	out := new(DestinationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationParameters) DeepCopyInto(out *DestinationParameters) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationParameters.
func (in *DestinationParameters) DeepCopy() *DestinationParameters {
	if in == nil {
		return nil
	}
	out := new(DestinationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfigurationInitParameters) DeepCopyInto(out *EncryptionConfigurationInitParameters) {
	*out = *in
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.KMSKey != nil {
		in, out := &in.KMSKey, &out.KMSKey
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyRef != nil {
		in, out := &in.KMSKeyRef, &out.KMSKeyRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeySelector != nil {
		in, out := &in.KMSKeySelector, &out.KMSKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfigurationInitParameters.
func (in *EncryptionConfigurationInitParameters) DeepCopy() *EncryptionConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfigurationObservation) DeepCopyInto(out *EncryptionConfigurationObservation) {
	*out = *in
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.KMSKey != nil {
		in, out := &in.KMSKey, &out.KMSKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfigurationObservation.
func (in *EncryptionConfigurationObservation) DeepCopy() *EncryptionConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfigurationParameters) DeepCopyInto(out *EncryptionConfigurationParameters) {
	*out = *in
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.KMSKey != nil {
		in, out := &in.KMSKey, &out.KMSKey
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyRef != nil {
		in, out := &in.KMSKeyRef, &out.KMSKeyRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeySelector != nil {
		in, out := &in.KMSKeySelector, &out.KMSKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfigurationParameters.
func (in *EncryptionConfigurationParameters) DeepCopy() *EncryptionConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanningConfigurationInitParameters) DeepCopyInto(out *ImageScanningConfigurationInitParameters) {
	*out = *in
	if in.ScanOnPush != nil {
		in, out := &in.ScanOnPush, &out.ScanOnPush
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanningConfigurationInitParameters.
func (in *ImageScanningConfigurationInitParameters) DeepCopy() *ImageScanningConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ImageScanningConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanningConfigurationObservation) DeepCopyInto(out *ImageScanningConfigurationObservation) {
	*out = *in
	if in.ScanOnPush != nil {
		in, out := &in.ScanOnPush, &out.ScanOnPush
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanningConfigurationObservation.
func (in *ImageScanningConfigurationObservation) DeepCopy() *ImageScanningConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ImageScanningConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanningConfigurationParameters) DeepCopyInto(out *ImageScanningConfigurationParameters) {
	*out = *in
	if in.ScanOnPush != nil {
		in, out := &in.ScanOnPush, &out.ScanOnPush
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanningConfigurationParameters.
func (in *ImageScanningConfigurationParameters) DeepCopy() *ImageScanningConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ImageScanningConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfiguration) DeepCopyInto(out *ReplicationConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfiguration.
func (in *ReplicationConfiguration) DeepCopy() *ReplicationConfiguration {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigurationInitParameters) DeepCopyInto(out *ReplicationConfigurationInitParameters) {
	*out = *in
	if in.ReplicationConfiguration != nil {
		in, out := &in.ReplicationConfiguration, &out.ReplicationConfiguration
		*out = new(ReplicationConfigurationReplicationConfigurationInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigurationInitParameters.
func (in *ReplicationConfigurationInitParameters) DeepCopy() *ReplicationConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigurationList) DeepCopyInto(out *ReplicationConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicationConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigurationList.
func (in *ReplicationConfigurationList) DeepCopy() *ReplicationConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigurationObservation) DeepCopyInto(out *ReplicationConfigurationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
	if in.ReplicationConfiguration != nil {
		in, out := &in.ReplicationConfiguration, &out.ReplicationConfiguration
		*out = new(ReplicationConfigurationReplicationConfigurationObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigurationObservation.
func (in *ReplicationConfigurationObservation) DeepCopy() *ReplicationConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigurationParameters) DeepCopyInto(out *ReplicationConfigurationParameters) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ReplicationConfiguration != nil {
		in, out := &in.ReplicationConfiguration, &out.ReplicationConfiguration
		*out = new(ReplicationConfigurationReplicationConfigurationParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigurationParameters.
func (in *ReplicationConfigurationParameters) DeepCopy() *ReplicationConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigurationReplicationConfigurationInitParameters) DeepCopyInto(out *ReplicationConfigurationReplicationConfigurationInitParameters) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]RuleInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigurationReplicationConfigurationInitParameters.
func (in *ReplicationConfigurationReplicationConfigurationInitParameters) DeepCopy() *ReplicationConfigurationReplicationConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigurationReplicationConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigurationReplicationConfigurationObservation) DeepCopyInto(out *ReplicationConfigurationReplicationConfigurationObservation) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]RuleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigurationReplicationConfigurationObservation.
func (in *ReplicationConfigurationReplicationConfigurationObservation) DeepCopy() *ReplicationConfigurationReplicationConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigurationReplicationConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigurationReplicationConfigurationParameters) DeepCopyInto(out *ReplicationConfigurationReplicationConfigurationParameters) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]RuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigurationReplicationConfigurationParameters.
func (in *ReplicationConfigurationReplicationConfigurationParameters) DeepCopy() *ReplicationConfigurationReplicationConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigurationReplicationConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigurationSpec) DeepCopyInto(out *ReplicationConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigurationSpec.
func (in *ReplicationConfigurationSpec) DeepCopy() *ReplicationConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfigurationStatus) DeepCopyInto(out *ReplicationConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfigurationStatus.
func (in *ReplicationConfigurationStatus) DeepCopy() *ReplicationConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Repository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryFilterInitParameters) DeepCopyInto(out *RepositoryFilterInitParameters) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.FilterType != nil {
		in, out := &in.FilterType, &out.FilterType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryFilterInitParameters.
func (in *RepositoryFilterInitParameters) DeepCopy() *RepositoryFilterInitParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryFilterInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryFilterObservation) DeepCopyInto(out *RepositoryFilterObservation) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.FilterType != nil {
		in, out := &in.FilterType, &out.FilterType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryFilterObservation.
func (in *RepositoryFilterObservation) DeepCopy() *RepositoryFilterObservation {
	if in == nil {
		return nil
	}
	out := new(RepositoryFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryFilterParameters) DeepCopyInto(out *RepositoryFilterParameters) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.FilterType != nil {
		in, out := &in.FilterType, &out.FilterType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryFilterParameters.
func (in *RepositoryFilterParameters) DeepCopy() *RepositoryFilterParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryInitParameters) DeepCopyInto(out *RepositoryInitParameters) {
	*out = *in
	if in.EncryptionConfiguration != nil {
		in, out := &in.EncryptionConfiguration, &out.EncryptionConfiguration
		*out = make([]EncryptionConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.ImageScanningConfiguration != nil {
		in, out := &in.ImageScanningConfiguration, &out.ImageScanningConfiguration
		*out = new(ImageScanningConfigurationInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageTagMutability != nil {
		in, out := &in.ImageTagMutability, &out.ImageTagMutability
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryInitParameters.
func (in *RepositoryInitParameters) DeepCopy() *RepositoryInitParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryList) DeepCopyInto(out *RepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryList.
func (in *RepositoryList) DeepCopy() *RepositoryList {
	if in == nil {
		return nil
	}
	out := new(RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryObservation) DeepCopyInto(out *RepositoryObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.EncryptionConfiguration != nil {
		in, out := &in.EncryptionConfiguration, &out.EncryptionConfiguration
		*out = make([]EncryptionConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ImageScanningConfiguration != nil {
		in, out := &in.ImageScanningConfiguration, &out.ImageScanningConfiguration
		*out = new(ImageScanningConfigurationObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageTagMutability != nil {
		in, out := &in.ImageTagMutability, &out.ImageTagMutability
		*out = new(string)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
	if in.RepositoryURL != nil {
		in, out := &in.RepositoryURL, &out.RepositoryURL
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryObservation.
func (in *RepositoryObservation) DeepCopy() *RepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(RepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryParameters) DeepCopyInto(out *RepositoryParameters) {
	*out = *in
	if in.EncryptionConfiguration != nil {
		in, out := &in.EncryptionConfiguration, &out.EncryptionConfiguration
		*out = make([]EncryptionConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.ImageScanningConfiguration != nil {
		in, out := &in.ImageScanningConfiguration, &out.ImageScanningConfiguration
		*out = new(ImageScanningConfigurationParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageTagMutability != nil {
		in, out := &in.ImageTagMutability, &out.ImageTagMutability
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryParameters.
func (in *RepositoryParameters) DeepCopy() *RepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpec) DeepCopyInto(out *RepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpec.
func (in *RepositorySpec) DeepCopy() *RepositorySpec {
	if in == nil {
		return nil
	}
	out := new(RepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryStatus) DeepCopyInto(out *RepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryStatus.
func (in *RepositoryStatus) DeepCopy() *RepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleInitParameters) DeepCopyInto(out *RuleInitParameters) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = make([]DestinationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RepositoryFilter != nil {
		in, out := &in.RepositoryFilter, &out.RepositoryFilter
		*out = make([]RepositoryFilterInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleInitParameters.
func (in *RuleInitParameters) DeepCopy() *RuleInitParameters {
	if in == nil {
		return nil
	}
	out := new(RuleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleObservation) DeepCopyInto(out *RuleObservation) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = make([]DestinationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RepositoryFilter != nil {
		in, out := &in.RepositoryFilter, &out.RepositoryFilter
		*out = make([]RepositoryFilterObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleObservation.
func (in *RuleObservation) DeepCopy() *RuleObservation {
	if in == nil {
		return nil
	}
	out := new(RuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleParameters) DeepCopyInto(out *RuleParameters) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = make([]DestinationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RepositoryFilter != nil {
		in, out := &in.RepositoryFilter, &out.RepositoryFilter
		*out = make([]RepositoryFilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleParameters.
func (in *RuleParameters) DeepCopy() *RuleParameters {
	if in == nil {
		return nil
	}
	out := new(RuleParameters)
	in.DeepCopyInto(out)
	return out
}
