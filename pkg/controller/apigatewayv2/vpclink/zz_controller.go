/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package vpclink

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/apigatewayv2/v1beta1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an VPCLink resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create VPCLink in AWS"
	errUpdate        = "cannot update VPCLink in AWS"
	errDescribe      = "failed to describe VPCLink"
	errDelete        = "failed to delete VPCLink"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.VPCLink)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := connectaws.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.VPCLink)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetVpcLinkInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetVpcLinkWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateVPCLink(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, diff, err := e.isUpToDate(ctx, cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		Diff:                    diff,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.VPCLink)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateVpcLinkInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateVpcLinkWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.CreatedDate != nil {
		cr.Status.AtProvider.CreatedDate = &metav1.Time{*resp.CreatedDate}
	} else {
		cr.Status.AtProvider.CreatedDate = nil
	}
	if resp.Name != nil {
		cr.Spec.ForProvider.Name = resp.Name
	} else {
		cr.Spec.ForProvider.Name = nil
	}
	if resp.SecurityGroupIds != nil {
		f2 := []*string{}
		for _, f2iter := range resp.SecurityGroupIds {
			var f2elem string
			f2elem = *f2iter
			f2 = append(f2, &f2elem)
		}
		cr.Status.AtProvider.SecurityGroupIDs = f2
	} else {
		cr.Status.AtProvider.SecurityGroupIDs = nil
	}
	if resp.SubnetIds != nil {
		f3 := []*string{}
		for _, f3iter := range resp.SubnetIds {
			var f3elem string
			f3elem = *f3iter
			f3 = append(f3, &f3elem)
		}
		cr.Status.AtProvider.SubnetIDs = f3
	} else {
		cr.Status.AtProvider.SubnetIDs = nil
	}
	if resp.Tags != nil {
		f4 := map[string]*string{}
		for f4key, f4valiter := range resp.Tags {
			var f4val string
			f4val = *f4valiter
			f4[f4key] = &f4val
		}
		cr.Spec.ForProvider.Tags = f4
	} else {
		cr.Spec.ForProvider.Tags = nil
	}
	if resp.VpcLinkId != nil {
		cr.Status.AtProvider.VPCLinkID = resp.VpcLinkId
	} else {
		cr.Status.AtProvider.VPCLinkID = nil
	}
	if resp.VpcLinkStatus != nil {
		cr.Status.AtProvider.VPCLinkStatus = resp.VpcLinkStatus
	} else {
		cr.Status.AtProvider.VPCLinkStatus = nil
	}
	if resp.VpcLinkStatusMessage != nil {
		cr.Status.AtProvider.VPCLinkStatusMessage = resp.VpcLinkStatusMessage
	} else {
		cr.Status.AtProvider.VPCLinkStatusMessage = nil
	}
	if resp.VpcLinkVersion != nil {
		cr.Status.AtProvider.VPCLinkVersion = resp.VpcLinkVersion
	} else {
		cr.Status.AtProvider.VPCLinkVersion = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.VPCLink)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateVpcLinkInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateVpcLinkWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.VPCLink)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteVpcLinkInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteVpcLinkWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ApiGatewayV2API, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.ApiGatewayV2API
	preObserve     func(context.Context, *svcapitypes.VPCLink, *svcsdk.GetVpcLinkInput) error
	postObserve    func(context.Context, *svcapitypes.VPCLink, *svcsdk.GetVpcLinkOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.VPCLinkParameters, *svcsdk.GetVpcLinkOutput) error
	isUpToDate     func(context.Context, *svcapitypes.VPCLink, *svcsdk.GetVpcLinkOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.VPCLink, *svcsdk.CreateVpcLinkInput) error
	postCreate     func(context.Context, *svcapitypes.VPCLink, *svcsdk.CreateVpcLinkOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.VPCLink, *svcsdk.DeleteVpcLinkInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.VPCLink, *svcsdk.DeleteVpcLinkOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.VPCLink, *svcsdk.UpdateVpcLinkInput) error
	postUpdate     func(context.Context, *svcapitypes.VPCLink, *svcsdk.UpdateVpcLinkOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.VPCLink, *svcsdk.GetVpcLinkInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.VPCLink, _ *svcsdk.GetVpcLinkOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.VPCLinkParameters, *svcsdk.GetVpcLinkOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.VPCLink, *svcsdk.GetVpcLinkOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.VPCLink, *svcsdk.CreateVpcLinkInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.VPCLink, _ *svcsdk.CreateVpcLinkOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.VPCLink, *svcsdk.DeleteVpcLinkInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.VPCLink, _ *svcsdk.DeleteVpcLinkOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.VPCLink, *svcsdk.UpdateVpcLinkInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.VPCLink, _ *svcsdk.UpdateVpcLinkOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
