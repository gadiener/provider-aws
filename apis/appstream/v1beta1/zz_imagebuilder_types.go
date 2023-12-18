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

type AccessEndpointInitParameters struct {

	// Type of interface endpoint.
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// Identifier (ID) of the VPC in which the interface endpoint is used.
	VpceID *string `json:"vpceId,omitempty" tf:"vpce_id,omitempty"`
}

type AccessEndpointObservation struct {

	// Type of interface endpoint.
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// Identifier (ID) of the VPC in which the interface endpoint is used.
	VpceID *string `json:"vpceId,omitempty" tf:"vpce_id,omitempty"`
}

type AccessEndpointParameters struct {

	// Type of interface endpoint.
	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType" tf:"endpoint_type,omitempty"`

	// Identifier (ID) of the VPC in which the interface endpoint is used.
	// +kubebuilder:validation:Optional
	VpceID *string `json:"vpceId,omitempty" tf:"vpce_id,omitempty"`
}

type ImageBuilderDomainJoinInfoInitParameters struct {

	// Fully qualified name of the directory (for example, corp.example.com).
	DirectoryName *string `json:"directoryName,omitempty" tf:"directory_name,omitempty"`

	// Distinguished name of the organizational unit for computer accounts.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName,omitempty" tf:"organizational_unit_distinguished_name,omitempty"`
}

type ImageBuilderDomainJoinInfoObservation struct {

	// Fully qualified name of the directory (for example, corp.example.com).
	DirectoryName *string `json:"directoryName,omitempty" tf:"directory_name,omitempty"`

	// Distinguished name of the organizational unit for computer accounts.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName,omitempty" tf:"organizational_unit_distinguished_name,omitempty"`
}

type ImageBuilderDomainJoinInfoParameters struct {

	// Fully qualified name of the directory (for example, corp.example.com).
	// +kubebuilder:validation:Optional
	DirectoryName *string `json:"directoryName,omitempty" tf:"directory_name,omitempty"`

	// Distinguished name of the organizational unit for computer accounts.
	// +kubebuilder:validation:Optional
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName,omitempty" tf:"organizational_unit_distinguished_name,omitempty"`
}

type ImageBuilderInitParameters struct {

	// Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
	AccessEndpoint []AccessEndpointInitParameters `json:"accessEndpoint,omitempty" tf:"access_endpoint,omitempty"`

	// Version of the AppStream 2.0 agent to use for this image builder.
	AppstreamAgentVersion *string `json:"appstreamAgentVersion,omitempty" tf:"appstream_agent_version,omitempty"`

	// Description to display.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Human-readable friendly name for the AppStream image builder.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
	DomainJoinInfo []ImageBuilderDomainJoinInfoInitParameters `json:"domainJoinInfo,omitempty" tf:"domain_join_info,omitempty"`

	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess *bool `json:"enableDefaultInternetAccess,omitempty" tf:"enable_default_internet_access,omitempty"`

	// ARN of the public, private, or shared image to use.
	ImageArn *string `json:"imageArn,omitempty" tf:"image_arn,omitempty"`

	// Instance type to use when launching the image builder.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block for the VPC configuration for the image builder. See below.
	VPCConfig []ImageBuilderVPCConfigInitParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ImageBuilderObservation struct {

	// Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
	AccessEndpoint []AccessEndpointObservation `json:"accessEndpoint,omitempty" tf:"access_endpoint,omitempty"`

	// Version of the AppStream 2.0 agent to use for this image builder.
	AppstreamAgentVersion *string `json:"appstreamAgentVersion,omitempty" tf:"appstream_agent_version,omitempty"`

	// ARN of the appstream image builder.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Date and time, in UTC and extended RFC 3339 format, when the image builder was created.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// Description to display.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Human-readable friendly name for the AppStream image builder.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
	DomainJoinInfo []ImageBuilderDomainJoinInfoObservation `json:"domainJoinInfo,omitempty" tf:"domain_join_info,omitempty"`

	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess *bool `json:"enableDefaultInternetAccess,omitempty" tf:"enable_default_internet_access,omitempty"`

	// ARN of the IAM role to apply to the image builder.
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Name of the image builder.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the public, private, or shared image to use.
	ImageArn *string `json:"imageArn,omitempty" tf:"image_arn,omitempty"`

	// Name of the image used to create the image builder.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// Instance type to use when launching the image builder.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// State of the image builder. Can be: PENDING, UPDATING_AGENT, RUNNING, STOPPING, STOPPED, REBOOTING, SNAPSHOTTING, DELETING, FAILED, UPDATING, PENDING_QUALIFICATION
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Configuration block for the VPC configuration for the image builder. See below.
	VPCConfig []ImageBuilderVPCConfigObservation `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ImageBuilderParameters struct {

	// Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
	// +kubebuilder:validation:Optional
	AccessEndpoint []AccessEndpointParameters `json:"accessEndpoint,omitempty" tf:"access_endpoint,omitempty"`

	// Version of the AppStream 2.0 agent to use for this image builder.
	// +kubebuilder:validation:Optional
	AppstreamAgentVersion *string `json:"appstreamAgentVersion,omitempty" tf:"appstream_agent_version,omitempty"`

	// Description to display.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Human-readable friendly name for the AppStream image builder.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
	// +kubebuilder:validation:Optional
	DomainJoinInfo []ImageBuilderDomainJoinInfoParameters `json:"domainJoinInfo,omitempty" tf:"domain_join_info,omitempty"`

	// Enables or disables default internet access for the image builder.
	// +kubebuilder:validation:Optional
	EnableDefaultInternetAccess *bool `json:"enableDefaultInternetAccess,omitempty" tf:"enable_default_internet_access,omitempty"`

	// ARN of the IAM role to apply to the image builder.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnRef *v1.Reference `json:"iamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnSelector *v1.Selector `json:"iamRoleArnSelector,omitempty" tf:"-"`

	// ARN of the public, private, or shared image to use.
	// +kubebuilder:validation:Optional
	ImageArn *string `json:"imageArn,omitempty" tf:"image_arn,omitempty"`

	// Instance type to use when launching the image builder.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block for the VPC configuration for the image builder. See below.
	// +kubebuilder:validation:Optional
	VPCConfig []ImageBuilderVPCConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ImageBuilderVPCConfigInitParameters struct {

	// Identifiers of the security groups for the image builder or image builder.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`
}

type ImageBuilderVPCConfigObservation struct {

	// Identifiers of the security groups for the image builder or image builder.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Identifiers of the subnets to which a network interface is attached from the image builder instance or image builder instance.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type ImageBuilderVPCConfigParameters struct {

	// Identifiers of the security groups for the image builder or image builder.
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Identifiers of the subnets to which a network interface is attached from the image builder instance or image builder instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

// ImageBuilderSpec defines the desired state of ImageBuilder
type ImageBuilderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageBuilderParameters `json:"forProvider"`
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
	InitProvider ImageBuilderInitParameters `json:"initProvider,omitempty"`
}

// ImageBuilderStatus defines the observed state of ImageBuilder.
type ImageBuilderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageBuilderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ImageBuilder is the Schema for the ImageBuilders API. Provides an AppStream image builder
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ImageBuilder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceType) || (has(self.initProvider) && has(self.initProvider.instanceType))",message="spec.forProvider.instanceType is a required parameter"
	Spec   ImageBuilderSpec   `json:"spec"`
	Status ImageBuilderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageBuilderList contains a list of ImageBuilders
type ImageBuilderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageBuilder `json:"items"`
}

// Repository type metadata.
var (
	ImageBuilder_Kind             = "ImageBuilder"
	ImageBuilder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageBuilder_Kind}.String()
	ImageBuilder_KindAPIVersion   = ImageBuilder_Kind + "." + CRDGroupVersion.String()
	ImageBuilder_GroupVersionKind = CRDGroupVersion.WithKind(ImageBuilder_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageBuilder{}, &ImageBuilderList{})
}
