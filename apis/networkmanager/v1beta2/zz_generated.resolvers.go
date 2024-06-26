// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ConnectAttachment) ResolveReferences( // ResolveReferences of this ConnectAttachment.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta1", "CoreNetwork", "CoreNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CoreNetworkID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CoreNetworkIDRef,
			Selector:     mg.Spec.ForProvider.CoreNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CoreNetworkID")
	}
	mg.Spec.ForProvider.CoreNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CoreNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta2", "VPCAttachment", "VPCAttachmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EdgeLocation),
			Extract:      resource.ExtractParamPath("edge_location", true),
			Reference:    mg.Spec.ForProvider.EdgeLocationRef,
			Selector:     mg.Spec.ForProvider.EdgeLocationSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EdgeLocation")
	}
	mg.Spec.ForProvider.EdgeLocation = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EdgeLocationRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta2", "VPCAttachment", "VPCAttachmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransportAttachmentID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.TransportAttachmentIDRef,
			Selector:     mg.Spec.ForProvider.TransportAttachmentIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransportAttachmentID")
	}
	mg.Spec.ForProvider.TransportAttachmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransportAttachmentIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta1", "CoreNetwork", "CoreNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CoreNetworkID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.CoreNetworkIDRef,
			Selector:     mg.Spec.InitProvider.CoreNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CoreNetworkID")
	}
	mg.Spec.InitProvider.CoreNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CoreNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta2", "VPCAttachment", "VPCAttachmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EdgeLocation),
			Extract:      resource.ExtractParamPath("edge_location", true),
			Reference:    mg.Spec.InitProvider.EdgeLocationRef,
			Selector:     mg.Spec.InitProvider.EdgeLocationSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EdgeLocation")
	}
	mg.Spec.InitProvider.EdgeLocation = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EdgeLocationRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta2", "VPCAttachment", "VPCAttachmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TransportAttachmentID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.TransportAttachmentIDRef,
			Selector:     mg.Spec.InitProvider.TransportAttachmentIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TransportAttachmentID")
	}
	mg.Spec.InitProvider.TransportAttachmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TransportAttachmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Device.
func (mg *Device) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta1", "GlobalNetwork", "GlobalNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
			Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta2", "Site", "SiteList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SiteIDRef,
			Selector:     mg.Spec.ForProvider.SiteIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SiteID")
	}
	mg.Spec.ForProvider.SiteID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SiteIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta1", "GlobalNetwork", "GlobalNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.GlobalNetworkID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.GlobalNetworkIDRef,
			Selector:     mg.Spec.InitProvider.GlobalNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GlobalNetworkID")
	}
	mg.Spec.InitProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GlobalNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta2", "Site", "SiteList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SiteID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SiteIDRef,
			Selector:     mg.Spec.InitProvider.SiteIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SiteID")
	}
	mg.Spec.InitProvider.SiteID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SiteIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Link.
func (mg *Link) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta1", "GlobalNetwork", "GlobalNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
			Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta2", "Site", "SiteList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.SiteIDRef,
			Selector:     mg.Spec.ForProvider.SiteIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SiteID")
	}
	mg.Spec.ForProvider.SiteID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SiteIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta1", "GlobalNetwork", "GlobalNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.GlobalNetworkID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.GlobalNetworkIDRef,
			Selector:     mg.Spec.InitProvider.GlobalNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GlobalNetworkID")
	}
	mg.Spec.InitProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GlobalNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta2", "Site", "SiteList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SiteID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.SiteIDRef,
			Selector:     mg.Spec.InitProvider.SiteIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SiteID")
	}
	mg.Spec.InitProvider.SiteID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SiteIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Site.
func (mg *Site) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta1", "GlobalNetwork", "GlobalNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
			Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta1", "GlobalNetwork", "GlobalNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.GlobalNetworkID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.GlobalNetworkIDRef,
			Selector:     mg.Spec.InitProvider.GlobalNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GlobalNetworkID")
	}
	mg.Spec.InitProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GlobalNetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VPCAttachment.
func (mg *VPCAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta1", "CoreNetwork", "CoreNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CoreNetworkID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CoreNetworkIDRef,
			Selector:     mg.Spec.ForProvider.CoreNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CoreNetworkID")
	}
	mg.Spec.ForProvider.CoreNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CoreNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetArns),
			Extract:       common.ARNExtractor(),
			References:    mg.Spec.ForProvider.SubnetArnsRefs,
			Selector:      mg.Spec.ForProvider.SubnetArnsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetArns")
	}
	mg.Spec.ForProvider.SubnetArns = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetArnsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.VPCArnRef,
			Selector:     mg.Spec.ForProvider.VPCArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCArn")
	}
	mg.Spec.ForProvider.VPCArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("networkmanager.aws.upbound.io", "v1beta1", "CoreNetwork", "CoreNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CoreNetworkID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.CoreNetworkIDRef,
			Selector:     mg.Spec.InitProvider.CoreNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CoreNetworkID")
	}
	mg.Spec.InitProvider.CoreNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CoreNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SubnetArns),
			Extract:       common.ARNExtractor(),
			References:    mg.Spec.InitProvider.SubnetArnsRefs,
			Selector:      mg.Spec.InitProvider.SubnetArnsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetArns")
	}
	mg.Spec.InitProvider.SubnetArns = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetArnsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.VPCArnRef,
			Selector:     mg.Spec.InitProvider.VPCArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCArn")
	}
	mg.Spec.InitProvider.VPCArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCArnRef = rsp.ResolvedReference

	return nil
}
