// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RegexPatternSetInitParameters struct {

	// A friendly description of the regular expression pattern set.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A friendly name of the regular expression pattern set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as B[a@]dB[o0]t. See Regular Expression below for details.
	RegularExpression []RegularExpressionInitParameters `json:"regularExpression,omitempty" tf:"regular_expression,omitempty"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are CLOUDFRONT or REGIONAL. To work with CloudFront, you must also specify the region us-east-1 (N. Virginia) on the AWS provider.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RegexPatternSetObservation struct {

	// The Amazon Resource Name (ARN) that identifies the cluster.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A friendly description of the regular expression pattern set.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique identifier for the set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LockToken *string `json:"lockToken,omitempty" tf:"lock_token,omitempty"`

	// A friendly name of the regular expression pattern set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as B[a@]dB[o0]t. See Regular Expression below for details.
	RegularExpression []RegularExpressionObservation `json:"regularExpression,omitempty" tf:"regular_expression,omitempty"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are CLOUDFRONT or REGIONAL. To work with CloudFront, you must also specify the region us-east-1 (N. Virginia) on the AWS provider.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RegexPatternSetParameters struct {

	// A friendly description of the regular expression pattern set.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A friendly name of the regular expression pattern set.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as B[a@]dB[o0]t. See Regular Expression below for details.
	// +kubebuilder:validation:Optional
	RegularExpression []RegularExpressionParameters `json:"regularExpression,omitempty" tf:"regular_expression,omitempty"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are CLOUDFRONT or REGIONAL. To work with CloudFront, you must also specify the region us-east-1 (N. Virginia) on the AWS provider.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RegularExpressionInitParameters struct {

	// The string representing the regular expression, see the AWS WAF documentation for more information.
	RegexString *string `json:"regexString,omitempty" tf:"regex_string,omitempty"`
}

type RegularExpressionObservation struct {

	// The string representing the regular expression, see the AWS WAF documentation for more information.
	RegexString *string `json:"regexString,omitempty" tf:"regex_string,omitempty"`
}

type RegularExpressionParameters struct {

	// The string representing the regular expression, see the AWS WAF documentation for more information.
	// +kubebuilder:validation:Optional
	RegexString *string `json:"regexString" tf:"regex_string,omitempty"`
}

// RegexPatternSetSpec defines the desired state of RegexPatternSet
type RegexPatternSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegexPatternSetParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RegexPatternSetInitParameters `json:"initProvider,omitempty"`
}

// RegexPatternSetStatus defines the observed state of RegexPatternSet.
type RegexPatternSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegexPatternSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegexPatternSet is the Schema for the RegexPatternSets API. Provides an AWS WAFv2 Regex Pattern Set resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RegexPatternSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scope) || (has(self.initProvider) && has(self.initProvider.scope))",message="spec.forProvider.scope is a required parameter"
	Spec   RegexPatternSetSpec   `json:"spec"`
	Status RegexPatternSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegexPatternSetList contains a list of RegexPatternSets
type RegexPatternSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegexPatternSet `json:"items"`
}

// Repository type metadata.
var (
	RegexPatternSet_Kind             = "RegexPatternSet"
	RegexPatternSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegexPatternSet_Kind}.String()
	RegexPatternSet_KindAPIVersion   = RegexPatternSet_Kind + "." + CRDGroupVersion.String()
	RegexPatternSet_GroupVersionKind = CRDGroupVersion.WithKind(RegexPatternSet_Kind)
)

func init() {
	SchemeBuilder.Register(&RegexPatternSet{}, &RegexPatternSetList{})
}
