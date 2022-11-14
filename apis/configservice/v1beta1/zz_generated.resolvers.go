/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/lambda/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ConfigRule.
func (mg *ConfigRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Source); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Source[i3].SourceIdentifier),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Source[i3].SourceIdentifierRef,
			Selector:     mg.Spec.ForProvider.Source[i3].SourceIdentifierSelector,
			To: reference.To{
				List:    &v1beta1.FunctionList{},
				Managed: &v1beta1.Function{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Source[i3].SourceIdentifier")
		}
		mg.Spec.ForProvider.Source[i3].SourceIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Source[i3].SourceIdentifierRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ConfigurationAggregator.
func (mg *ConfigurationAggregator) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.OrganizationAggregationSource); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrganizationAggregationSource[i3].RoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.OrganizationAggregationSource[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.OrganizationAggregationSource[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta11.RoleList{},
				Managed: &v1beta11.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.OrganizationAggregationSource[i3].RoleArn")
		}
		mg.Spec.ForProvider.OrganizationAggregationSource[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.OrganizationAggregationSource[i3].RoleArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ConfigurationRecorder.
func (mg *ConfigurationRecorder) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DeliveryChannel.
func (mg *DeliveryChannel) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3BucketName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.S3BucketNameRef,
		Selector:     mg.Spec.ForProvider.S3BucketNameSelector,
		To: reference.To{
			List:    &v1beta12.BucketList{},
			Managed: &v1beta12.Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.S3BucketName")
	}
	mg.Spec.ForProvider.S3BucketName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.S3BucketNameRef = rsp.ResolvedReference

	return nil
}