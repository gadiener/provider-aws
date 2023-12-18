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

type VPCPeeringConnectionAccepterAccepterInitParameters struct {

	// Indicates whether a local ClassicLink connection can communicate
	// with the peer VPC over the VPC Peering Connection.
	AllowClassicLinkToRemoteVPC *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// Indicates whether a local VPC can resolve public DNS hostnames to
	// private IP addresses when queried from instances in a peer VPC.
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// Indicates whether a local VPC can communicate with a ClassicLink
	// connection in the peer VPC over the VPC Peering Connection.
	AllowVPCToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VPCPeeringConnectionAccepterAccepterObservation struct {

	// Indicates whether a local ClassicLink connection can communicate
	// with the peer VPC over the VPC Peering Connection.
	AllowClassicLinkToRemoteVPC *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// Indicates whether a local VPC can resolve public DNS hostnames to
	// private IP addresses when queried from instances in a peer VPC.
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// Indicates whether a local VPC can communicate with a ClassicLink
	// connection in the peer VPC over the VPC Peering Connection.
	AllowVPCToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VPCPeeringConnectionAccepterAccepterParameters struct {

	// Indicates whether a local ClassicLink connection can communicate
	// with the peer VPC over the VPC Peering Connection.
	// +kubebuilder:validation:Optional
	AllowClassicLinkToRemoteVPC *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// Indicates whether a local VPC can resolve public DNS hostnames to
	// private IP addresses when queried from instances in a peer VPC.
	// +kubebuilder:validation:Optional
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// Indicates whether a local VPC can communicate with a ClassicLink
	// connection in the peer VPC over the VPC Peering Connection.
	// +kubebuilder:validation:Optional
	AllowVPCToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VPCPeeringConnectionAccepterInitParameters struct {

	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
	Accepter []VPCPeeringConnectionAccepterAccepterInitParameters `json:"accepter,omitempty" tf:"accepter,omitempty"`

	// Whether or not to accept the peering request. Defaults to false.
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
	Requester []VPCPeeringConnectionAccepterRequesterInitParameters `json:"requester,omitempty" tf:"requester,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VPCPeeringConnectionAccepterObservation struct {

	// The status of the VPC Peering Connection request.
	AcceptStatus *string `json:"acceptStatus,omitempty" tf:"accept_status,omitempty"`

	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
	Accepter []VPCPeeringConnectionAccepterAccepterObservation `json:"accepter,omitempty" tf:"accepter,omitempty"`

	// Whether or not to accept the peering request. Defaults to false.
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The ID of the VPC Peering Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The AWS account ID of the owner of the requester VPC.
	PeerOwnerID *string `json:"peerOwnerId,omitempty" tf:"peer_owner_id,omitempty"`

	// The region of the accepter VPC.
	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	// The ID of the requester VPC.
	PeerVPCID *string `json:"peerVpcId,omitempty" tf:"peer_vpc_id,omitempty"`

	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
	Requester []VPCPeeringConnectionAccepterRequesterObservation `json:"requester,omitempty" tf:"requester,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the accepter VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The VPC Peering Connection ID to manage.
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}

type VPCPeeringConnectionAccepterParameters struct {

	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
	// +kubebuilder:validation:Optional
	Accepter []VPCPeeringConnectionAccepterAccepterParameters `json:"accepter,omitempty" tf:"accepter,omitempty"`

	// Whether or not to accept the peering request. Defaults to false.
	// +kubebuilder:validation:Optional
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
	// +kubebuilder:validation:Optional
	Requester []VPCPeeringConnectionAccepterRequesterParameters `json:"requester,omitempty" tf:"requester,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC Peering Connection ID to manage.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCPeeringConnection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`

	// Reference to a VPCPeeringConnection in ec2 to populate vpcPeeringConnectionId.
	// +kubebuilder:validation:Optional
	VPCPeeringConnectionIDRef *v1.Reference `json:"vpcPeeringConnectionIdRef,omitempty" tf:"-"`

	// Selector for a VPCPeeringConnection in ec2 to populate vpcPeeringConnectionId.
	// +kubebuilder:validation:Optional
	VPCPeeringConnectionIDSelector *v1.Selector `json:"vpcPeeringConnectionIdSelector,omitempty" tf:"-"`
}

type VPCPeeringConnectionAccepterRequesterInitParameters struct {

	// Indicates whether a local ClassicLink connection can communicate
	// with the peer VPC over the VPC Peering Connection.
	AllowClassicLinkToRemoteVPC *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// Indicates whether a local VPC can resolve public DNS hostnames to
	// private IP addresses when queried from instances in a peer VPC.
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// Indicates whether a local VPC can communicate with a ClassicLink
	// connection in the peer VPC over the VPC Peering Connection.
	AllowVPCToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VPCPeeringConnectionAccepterRequesterObservation struct {

	// Indicates whether a local ClassicLink connection can communicate
	// with the peer VPC over the VPC Peering Connection.
	AllowClassicLinkToRemoteVPC *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// Indicates whether a local VPC can resolve public DNS hostnames to
	// private IP addresses when queried from instances in a peer VPC.
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// Indicates whether a local VPC can communicate with a ClassicLink
	// connection in the peer VPC over the VPC Peering Connection.
	AllowVPCToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VPCPeeringConnectionAccepterRequesterParameters struct {

	// Indicates whether a local ClassicLink connection can communicate
	// with the peer VPC over the VPC Peering Connection.
	// +kubebuilder:validation:Optional
	AllowClassicLinkToRemoteVPC *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// Indicates whether a local VPC can resolve public DNS hostnames to
	// private IP addresses when queried from instances in a peer VPC.
	// +kubebuilder:validation:Optional
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// Indicates whether a local VPC can communicate with a ClassicLink
	// connection in the peer VPC over the VPC Peering Connection.
	// +kubebuilder:validation:Optional
	AllowVPCToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

// VPCPeeringConnectionAccepterSpec defines the desired state of VPCPeeringConnectionAccepter
type VPCPeeringConnectionAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCPeeringConnectionAccepterParameters `json:"forProvider"`
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
	InitProvider VPCPeeringConnectionAccepterInitParameters `json:"initProvider,omitempty"`
}

// VPCPeeringConnectionAccepterStatus defines the observed state of VPCPeeringConnectionAccepter.
type VPCPeeringConnectionAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCPeeringConnectionAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCPeeringConnectionAccepter is the Schema for the VPCPeeringConnectionAccepters API. Manage the accepter's side of a VPC Peering Connection.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCPeeringConnectionAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCPeeringConnectionAccepterSpec   `json:"spec"`
	Status            VPCPeeringConnectionAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCPeeringConnectionAccepterList contains a list of VPCPeeringConnectionAccepters
type VPCPeeringConnectionAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCPeeringConnectionAccepter `json:"items"`
}

// Repository type metadata.
var (
	VPCPeeringConnectionAccepter_Kind             = "VPCPeeringConnectionAccepter"
	VPCPeeringConnectionAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCPeeringConnectionAccepter_Kind}.String()
	VPCPeeringConnectionAccepter_KindAPIVersion   = VPCPeeringConnectionAccepter_Kind + "." + CRDGroupVersion.String()
	VPCPeeringConnectionAccepter_GroupVersionKind = CRDGroupVersion.WithKind(VPCPeeringConnectionAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCPeeringConnectionAccepter{}, &VPCPeeringConnectionAccepterList{})
}
