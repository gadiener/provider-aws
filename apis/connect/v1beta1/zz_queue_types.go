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

type OutboundCallerConfigInitParameters struct {

	// Specifies the caller ID name.
	OutboundCallerIDName *string `json:"outboundCallerIdName,omitempty" tf:"outbound_caller_id_name,omitempty"`

	// Specifies the caller ID number.
	OutboundCallerIDNumberID *string `json:"outboundCallerIdNumberId,omitempty" tf:"outbound_caller_id_number_id,omitempty"`

	// Specifies outbound whisper flow to be used during an outbound call.
	OutboundFlowID *string `json:"outboundFlowId,omitempty" tf:"outbound_flow_id,omitempty"`
}

type OutboundCallerConfigObservation struct {

	// Specifies the caller ID name.
	OutboundCallerIDName *string `json:"outboundCallerIdName,omitempty" tf:"outbound_caller_id_name,omitempty"`

	// Specifies the caller ID number.
	OutboundCallerIDNumberID *string `json:"outboundCallerIdNumberId,omitempty" tf:"outbound_caller_id_number_id,omitempty"`

	// Specifies outbound whisper flow to be used during an outbound call.
	OutboundFlowID *string `json:"outboundFlowId,omitempty" tf:"outbound_flow_id,omitempty"`
}

type OutboundCallerConfigParameters struct {

	// Specifies the caller ID name.
	// +kubebuilder:validation:Optional
	OutboundCallerIDName *string `json:"outboundCallerIdName,omitempty" tf:"outbound_caller_id_name,omitempty"`

	// Specifies the caller ID number.
	// +kubebuilder:validation:Optional
	OutboundCallerIDNumberID *string `json:"outboundCallerIdNumberId,omitempty" tf:"outbound_caller_id_number_id,omitempty"`

	// Specifies outbound whisper flow to be used during an outbound call.
	// +kubebuilder:validation:Optional
	OutboundFlowID *string `json:"outboundFlowId,omitempty" tf:"outbound_flow_id,omitempty"`
}

type QueueInitParameters struct {

	// Specifies the description of the Queue.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the maximum number of contacts that can be in the queue before it is considered full. Minimum value of 0.
	MaxContacts *float64 `json:"maxContacts,omitempty" tf:"max_contacts,omitempty"`

	// Specifies the name of the Queue.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that defines the outbound caller ID name, number, and outbound whisper flow. The Outbound Caller Config block is documented below.
	OutboundCallerConfig []OutboundCallerConfigInitParameters `json:"outboundCallerConfig,omitempty" tf:"outbound_caller_config,omitempty"`

	// Specifies a list of quick connects ids that determine the quick connects available to agents who are working the queue.
	// +listType=set
	QuickConnectIds []*string `json:"quickConnectIds,omitempty" tf:"quick_connect_ids,omitempty"`

	// Specifies the description of the Queue. Valid values are ENABLED, DISABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type QueueObservation struct {

	// The Amazon Resource Name (ARN) of the Queue.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies the description of the Queue.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the identifier of the Hours of Operation.
	HoursOfOperationID *string `json:"hoursOfOperationId,omitempty" tf:"hours_of_operation_id,omitempty"`

	// The identifier of the hosting Amazon Connect Instance and identifier of the Queue separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the maximum number of contacts that can be in the queue before it is considered full. Minimum value of 0.
	MaxContacts *float64 `json:"maxContacts,omitempty" tf:"max_contacts,omitempty"`

	// Specifies the name of the Queue.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that defines the outbound caller ID name, number, and outbound whisper flow. The Outbound Caller Config block is documented below.
	OutboundCallerConfig []OutboundCallerConfigObservation `json:"outboundCallerConfig,omitempty" tf:"outbound_caller_config,omitempty"`

	// The identifier for the Queue.
	QueueID *string `json:"queueId,omitempty" tf:"queue_id,omitempty"`

	// Specifies a list of quick connects ids that determine the quick connects available to agents who are working the queue.
	// +listType=set
	QuickConnectIds []*string `json:"quickConnectIds,omitempty" tf:"quick_connect_ids,omitempty"`

	// +listType=set
	QuickConnectIdsAssociated []*string `json:"quickConnectIdsAssociated,omitempty" tf:"quick_connect_ids_associated,omitempty"`

	// Specifies the description of the Queue. Valid values are ENABLED, DISABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type QueueParameters struct {

	// Specifies the description of the Queue.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the identifier of the Hours of Operation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.HoursOfOperation
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("hours_of_operation_id",true)
	// +kubebuilder:validation:Optional
	HoursOfOperationID *string `json:"hoursOfOperationId,omitempty" tf:"hours_of_operation_id,omitempty"`

	// Reference to a HoursOfOperation in connect to populate hoursOfOperationId.
	// +kubebuilder:validation:Optional
	HoursOfOperationIDRef *v1.Reference `json:"hoursOfOperationIdRef,omitempty" tf:"-"`

	// Selector for a HoursOfOperation in connect to populate hoursOfOperationId.
	// +kubebuilder:validation:Optional
	HoursOfOperationIDSelector *v1.Selector `json:"hoursOfOperationIdSelector,omitempty" tf:"-"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the maximum number of contacts that can be in the queue before it is considered full. Minimum value of 0.
	// +kubebuilder:validation:Optional
	MaxContacts *float64 `json:"maxContacts,omitempty" tf:"max_contacts,omitempty"`

	// Specifies the name of the Queue.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that defines the outbound caller ID name, number, and outbound whisper flow. The Outbound Caller Config block is documented below.
	// +kubebuilder:validation:Optional
	OutboundCallerConfig []OutboundCallerConfigParameters `json:"outboundCallerConfig,omitempty" tf:"outbound_caller_config,omitempty"`

	// Specifies a list of quick connects ids that determine the quick connects available to agents who are working the queue.
	// +kubebuilder:validation:Optional
	// +listType=set
	QuickConnectIds []*string `json:"quickConnectIds,omitempty" tf:"quick_connect_ids,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the description of the Queue. Valid values are ENABLED, DISABLED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// QueueSpec defines the desired state of Queue
type QueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueParameters `json:"forProvider"`
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
	InitProvider QueueInitParameters `json:"initProvider,omitempty"`
}

// QueueStatus defines the observed state of Queue.
type QueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Queue is the Schema for the Queues API. Provides details about a specific Amazon Connect Queue
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Queue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   QueueSpec   `json:"spec"`
	Status QueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueList contains a list of Queues
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Queue `json:"items"`
}

// Repository type metadata.
var (
	Queue_Kind             = "Queue"
	Queue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Queue_Kind}.String()
	Queue_KindAPIVersion   = Queue_Kind + "." + CRDGroupVersion.String()
	Queue_GroupVersionKind = CRDGroupVersion.WithKind(Queue_Kind)
)

func init() {
	SchemeBuilder.Register(&Queue{}, &QueueList{})
}
