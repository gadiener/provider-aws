// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Domain.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Domain) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionKey),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.EncryptionKeyRef,
			Selector:     mg.Spec.ForProvider.EncryptionKeySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionKey")
	}
	mg.Spec.ForProvider.EncryptionKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EncryptionKeyRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EncryptionKey),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.EncryptionKeyRef,
			Selector:     mg.Spec.InitProvider.EncryptionKeySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EncryptionKey")
	}
	mg.Spec.InitProvider.EncryptionKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EncryptionKeyRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DomainPermissionsPolicy.
func (mg *DomainPermissionsPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("codeartifact.aws.upbound.io", "v1beta1", "Domain", "DomainList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Domain),
			Extract:      resource.ExtractParamPath("domain", false),
			Reference:    mg.Spec.ForProvider.DomainRef,
			Selector:     mg.Spec.ForProvider.DomainSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Domain")
	}
	mg.Spec.ForProvider.Domain = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("codeartifact.aws.upbound.io", "v1beta1", "Domain", "DomainList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Domain),
			Extract:      resource.ExtractParamPath("domain", false),
			Reference:    mg.Spec.InitProvider.DomainRef,
			Selector:     mg.Spec.InitProvider.DomainSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Domain")
	}
	mg.Spec.InitProvider.Domain = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DomainRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Repository.
func (mg *Repository) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("codeartifact.aws.upbound.io", "v1beta1", "Domain", "DomainList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Domain),
			Extract:      resource.ExtractParamPath("domain", false),
			Reference:    mg.Spec.ForProvider.DomainRef,
			Selector:     mg.Spec.ForProvider.DomainSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Domain")
	}
	mg.Spec.ForProvider.Domain = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Upstream); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("codeartifact.aws.upbound.io", "v1beta1", "Repository", "RepositoryList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Upstream[i3].RepositoryName),
				Extract:      resource.ExtractParamPath("repository", false),
				Reference:    mg.Spec.ForProvider.Upstream[i3].RepositoryNameRef,
				Selector:     mg.Spec.ForProvider.Upstream[i3].RepositoryNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Upstream[i3].RepositoryName")
		}
		mg.Spec.ForProvider.Upstream[i3].RepositoryName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Upstream[i3].RepositoryNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("codeartifact.aws.upbound.io", "v1beta1", "Domain", "DomainList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Domain),
			Extract:      resource.ExtractParamPath("domain", false),
			Reference:    mg.Spec.InitProvider.DomainRef,
			Selector:     mg.Spec.InitProvider.DomainSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Domain")
	}
	mg.Spec.InitProvider.Domain = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DomainRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Upstream); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("codeartifact.aws.upbound.io", "v1beta1", "Repository", "RepositoryList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Upstream[i3].RepositoryName),
				Extract:      resource.ExtractParamPath("repository", false),
				Reference:    mg.Spec.InitProvider.Upstream[i3].RepositoryNameRef,
				Selector:     mg.Spec.InitProvider.Upstream[i3].RepositoryNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Upstream[i3].RepositoryName")
		}
		mg.Spec.InitProvider.Upstream[i3].RepositoryName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Upstream[i3].RepositoryNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this RepositoryPermissionsPolicy.
func (mg *RepositoryPermissionsPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("codeartifact.aws.upbound.io", "v1beta1", "Domain", "DomainList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Domain),
			Extract:      resource.ExtractParamPath("domain", false),
			Reference:    mg.Spec.ForProvider.DomainRef,
			Selector:     mg.Spec.ForProvider.DomainSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Domain")
	}
	mg.Spec.ForProvider.Domain = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("codeartifact.aws.upbound.io", "v1beta1", "Repository", "RepositoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
			Extract:      resource.ExtractParamPath("repository", false),
			Reference:    mg.Spec.ForProvider.RepositoryRef,
			Selector:     mg.Spec.ForProvider.RepositorySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("codeartifact.aws.upbound.io", "v1beta1", "Domain", "DomainList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Domain),
			Extract:      resource.ExtractParamPath("domain", false),
			Reference:    mg.Spec.InitProvider.DomainRef,
			Selector:     mg.Spec.InitProvider.DomainSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Domain")
	}
	mg.Spec.InitProvider.Domain = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DomainRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("codeartifact.aws.upbound.io", "v1beta1", "Repository", "RepositoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
			Extract:      resource.ExtractParamPath("repository", false),
			Reference:    mg.Spec.InitProvider.RepositoryRef,
			Selector:     mg.Spec.InitProvider.RepositorySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}
