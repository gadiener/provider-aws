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

type ReplicationTaskInitParameters struct {

	// Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see Determining a CDC native start point.
	CdcStartPosition *string `json:"cdcStartPosition,omitempty" tf:"cdc_start_position,omitempty"`

	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime *string `json:"cdcStartTime,omitempty" tf:"cdc_start_time,omitempty"`

	// The migration type. Can be one of full-load | cdc | full-load-and-cdc.
	MigrationType *string `json:"migrationType,omitempty" tf:"migration_type,omitempty"`

	// An escaped JSON string that contains the task settings. For a complete list of task settings, see Task Settings for AWS Database Migration Service Tasks.
	ReplicationTaskSettings *string `json:"replicationTaskSettings,omitempty" tf:"replication_task_settings,omitempty"`

	// Whether to run or stop the replication task.
	StartReplicationTask *bool `json:"startReplicationTask,omitempty" tf:"start_replication_task,omitempty"`

	// An escaped JSON string that contains the table mappings. For information on table mapping see Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data
	TableMappings *string `json:"tableMappings,omitempty" tf:"table_mappings,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ReplicationTaskObservation struct {

	// Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see Determining a CDC native start point.
	CdcStartPosition *string `json:"cdcStartPosition,omitempty" tf:"cdc_start_position,omitempty"`

	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime *string `json:"cdcStartTime,omitempty" tf:"cdc_start_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The migration type. Can be one of full-load | cdc | full-load-and-cdc.
	MigrationType *string `json:"migrationType,omitempty" tf:"migration_type,omitempty"`

	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn *string `json:"replicationInstanceArn,omitempty" tf:"replication_instance_arn,omitempty"`

	// The Amazon Resource Name (ARN) for the replication task.
	ReplicationTaskArn *string `json:"replicationTaskArn,omitempty" tf:"replication_task_arn,omitempty"`

	// An escaped JSON string that contains the task settings. For a complete list of task settings, see Task Settings for AWS Database Migration Service Tasks.
	ReplicationTaskSettings *string `json:"replicationTaskSettings,omitempty" tf:"replication_task_settings,omitempty"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn *string `json:"sourceEndpointArn,omitempty" tf:"source_endpoint_arn,omitempty"`

	// Whether to run or stop the replication task.
	StartReplicationTask *bool `json:"startReplicationTask,omitempty" tf:"start_replication_task,omitempty"`

	// Replication Task status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// An escaped JSON string that contains the table mappings. For information on table mapping see Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data
	TableMappings *string `json:"tableMappings,omitempty" tf:"table_mappings,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn *string `json:"targetEndpointArn,omitempty" tf:"target_endpoint_arn,omitempty"`
}

type ReplicationTaskParameters struct {

	// Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see Determining a CDC native start point.
	// +kubebuilder:validation:Optional
	CdcStartPosition *string `json:"cdcStartPosition,omitempty" tf:"cdc_start_position,omitempty"`

	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	// +kubebuilder:validation:Optional
	CdcStartTime *string `json:"cdcStartTime,omitempty" tf:"cdc_start_time,omitempty"`

	// The migration type. Can be one of full-load | cdc | full-load-and-cdc.
	// +kubebuilder:validation:Optional
	MigrationType *string `json:"migrationType,omitempty" tf:"migration_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of the replication instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/dms/v1beta1.ReplicationInstance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("replication_instance_arn",true)
	// +kubebuilder:validation:Optional
	ReplicationInstanceArn *string `json:"replicationInstanceArn,omitempty" tf:"replication_instance_arn,omitempty"`

	// Reference to a ReplicationInstance in dms to populate replicationInstanceArn.
	// +kubebuilder:validation:Optional
	ReplicationInstanceArnRef *v1.Reference `json:"replicationInstanceArnRef,omitempty" tf:"-"`

	// Selector for a ReplicationInstance in dms to populate replicationInstanceArn.
	// +kubebuilder:validation:Optional
	ReplicationInstanceArnSelector *v1.Selector `json:"replicationInstanceArnSelector,omitempty" tf:"-"`

	// An escaped JSON string that contains the task settings. For a complete list of task settings, see Task Settings for AWS Database Migration Service Tasks.
	// +kubebuilder:validation:Optional
	ReplicationTaskSettings *string `json:"replicationTaskSettings,omitempty" tf:"replication_task_settings,omitempty"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/dms/v1beta1.Endpoint
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("endpoint_arn",true)
	// +kubebuilder:validation:Optional
	SourceEndpointArn *string `json:"sourceEndpointArn,omitempty" tf:"source_endpoint_arn,omitempty"`

	// Reference to a Endpoint in dms to populate sourceEndpointArn.
	// +kubebuilder:validation:Optional
	SourceEndpointArnRef *v1.Reference `json:"sourceEndpointArnRef,omitempty" tf:"-"`

	// Selector for a Endpoint in dms to populate sourceEndpointArn.
	// +kubebuilder:validation:Optional
	SourceEndpointArnSelector *v1.Selector `json:"sourceEndpointArnSelector,omitempty" tf:"-"`

	// Whether to run or stop the replication task.
	// +kubebuilder:validation:Optional
	StartReplicationTask *bool `json:"startReplicationTask,omitempty" tf:"start_replication_task,omitempty"`

	// An escaped JSON string that contains the table mappings. For information on table mapping see Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data
	// +kubebuilder:validation:Optional
	TableMappings *string `json:"tableMappings,omitempty" tf:"table_mappings,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/dms/v1beta1.Endpoint
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("endpoint_arn",true)
	// +kubebuilder:validation:Optional
	TargetEndpointArn *string `json:"targetEndpointArn,omitempty" tf:"target_endpoint_arn,omitempty"`

	// Reference to a Endpoint in dms to populate targetEndpointArn.
	// +kubebuilder:validation:Optional
	TargetEndpointArnRef *v1.Reference `json:"targetEndpointArnRef,omitempty" tf:"-"`

	// Selector for a Endpoint in dms to populate targetEndpointArn.
	// +kubebuilder:validation:Optional
	TargetEndpointArnSelector *v1.Selector `json:"targetEndpointArnSelector,omitempty" tf:"-"`
}

// ReplicationTaskSpec defines the desired state of ReplicationTask
type ReplicationTaskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicationTaskParameters `json:"forProvider"`
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
	InitProvider ReplicationTaskInitParameters `json:"initProvider,omitempty"`
}

// ReplicationTaskStatus defines the observed state of ReplicationTask.
type ReplicationTaskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicationTaskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationTask is the Schema for the ReplicationTasks API. Provides a DMS (Data Migration Service) replication task resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReplicationTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.migrationType) || (has(self.initProvider) && has(self.initProvider.migrationType))",message="spec.forProvider.migrationType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tableMappings) || (has(self.initProvider) && has(self.initProvider.tableMappings))",message="spec.forProvider.tableMappings is a required parameter"
	Spec   ReplicationTaskSpec   `json:"spec"`
	Status ReplicationTaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationTaskList contains a list of ReplicationTasks
type ReplicationTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationTask `json:"items"`
}

// Repository type metadata.
var (
	ReplicationTask_Kind             = "ReplicationTask"
	ReplicationTask_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReplicationTask_Kind}.String()
	ReplicationTask_KindAPIVersion   = ReplicationTask_Kind + "." + CRDGroupVersion.String()
	ReplicationTask_GroupVersionKind = CRDGroupVersion.WithKind(ReplicationTask_Kind)
)

func init() {
	SchemeBuilder.Register(&ReplicationTask{}, &ReplicationTaskList{})
}
