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

package resolverendpoint

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/route53resolver"
	svcsdk "github.com/aws/aws-sdk-go/service/route53resolver"
	svcsdkapi "github.com/aws/aws-sdk-go/service/route53resolver/route53resolveriface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/route53resolver/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an ResolverEndpoint resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create ResolverEndpoint in AWS"
	errUpdate        = "cannot update ResolverEndpoint in AWS"
	errDescribe      = "failed to describe ResolverEndpoint"
	errDelete        = "failed to delete ResolverEndpoint"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.ResolverEndpoint)
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
	cr, ok := mg.(*svcapitypes.ResolverEndpoint)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetResolverEndpointInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetResolverEndpointWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateResolverEndpoint(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

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
	cr, ok := mg.(*svcapitypes.ResolverEndpoint)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateResolverEndpointInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateResolverEndpointWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.ResolverEndpoint.Arn != nil {
		cr.Status.AtProvider.ARN = resp.ResolverEndpoint.Arn
	} else {
		cr.Status.AtProvider.ARN = nil
	}
	if resp.ResolverEndpoint.CreationTime != nil {
		cr.Status.AtProvider.CreationTime = resp.ResolverEndpoint.CreationTime
	} else {
		cr.Status.AtProvider.CreationTime = nil
	}
	if resp.ResolverEndpoint.CreatorRequestId != nil {
		cr.Status.AtProvider.CreatorRequestID = resp.ResolverEndpoint.CreatorRequestId
	} else {
		cr.Status.AtProvider.CreatorRequestID = nil
	}
	if resp.ResolverEndpoint.Direction != nil {
		cr.Spec.ForProvider.Direction = resp.ResolverEndpoint.Direction
	} else {
		cr.Spec.ForProvider.Direction = nil
	}
	if resp.ResolverEndpoint.HostVPCId != nil {
		cr.Status.AtProvider.HostVPCID = resp.ResolverEndpoint.HostVPCId
	} else {
		cr.Status.AtProvider.HostVPCID = nil
	}
	if resp.ResolverEndpoint.Id != nil {
		cr.Status.AtProvider.ID = resp.ResolverEndpoint.Id
	} else {
		cr.Status.AtProvider.ID = nil
	}
	if resp.ResolverEndpoint.IpAddressCount != nil {
		cr.Status.AtProvider.IPAddressCount = resp.ResolverEndpoint.IpAddressCount
	} else {
		cr.Status.AtProvider.IPAddressCount = nil
	}
	if resp.ResolverEndpoint.ModificationTime != nil {
		cr.Status.AtProvider.ModificationTime = resp.ResolverEndpoint.ModificationTime
	} else {
		cr.Status.AtProvider.ModificationTime = nil
	}
	if resp.ResolverEndpoint.Name != nil {
		cr.Spec.ForProvider.Name = resp.ResolverEndpoint.Name
	} else {
		cr.Spec.ForProvider.Name = nil
	}
	if resp.ResolverEndpoint.OutpostArn != nil {
		cr.Spec.ForProvider.OutpostARN = resp.ResolverEndpoint.OutpostArn
	} else {
		cr.Spec.ForProvider.OutpostARN = nil
	}
	if resp.ResolverEndpoint.PreferredInstanceType != nil {
		cr.Spec.ForProvider.PreferredInstanceType = resp.ResolverEndpoint.PreferredInstanceType
	} else {
		cr.Spec.ForProvider.PreferredInstanceType = nil
	}
	if resp.ResolverEndpoint.ResolverEndpointType != nil {
		cr.Spec.ForProvider.ResolverEndpointType = resp.ResolverEndpoint.ResolverEndpointType
	} else {
		cr.Spec.ForProvider.ResolverEndpointType = nil
	}
	if resp.ResolverEndpoint.SecurityGroupIds != nil {
		f12 := []*string{}
		for _, f12iter := range resp.ResolverEndpoint.SecurityGroupIds {
			var f12elem string
			f12elem = *f12iter
			f12 = append(f12, &f12elem)
		}
		cr.Status.AtProvider.SecurityGroupIDs = f12
	} else {
		cr.Status.AtProvider.SecurityGroupIDs = nil
	}
	if resp.ResolverEndpoint.Status != nil {
		cr.Status.AtProvider.Status = resp.ResolverEndpoint.Status
	} else {
		cr.Status.AtProvider.Status = nil
	}
	if resp.ResolverEndpoint.StatusMessage != nil {
		cr.Status.AtProvider.StatusMessage = resp.ResolverEndpoint.StatusMessage
	} else {
		cr.Status.AtProvider.StatusMessage = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.ResolverEndpoint)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateResolverEndpointInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateResolverEndpointWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.ResolverEndpoint)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteResolverEndpointInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteResolverEndpointWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.Route53ResolverAPI, opts []option) *external {
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
	client         svcsdkapi.Route53ResolverAPI
	preObserve     func(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.GetResolverEndpointInput) error
	postObserve    func(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.GetResolverEndpointOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.ResolverEndpointParameters, *svcsdk.GetResolverEndpointOutput) error
	isUpToDate     func(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.GetResolverEndpointOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.CreateResolverEndpointInput) error
	postCreate     func(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.CreateResolverEndpointOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.DeleteResolverEndpointInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.DeleteResolverEndpointOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.UpdateResolverEndpointInput) error
	postUpdate     func(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.UpdateResolverEndpointOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.GetResolverEndpointInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.ResolverEndpoint, _ *svcsdk.GetResolverEndpointOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.ResolverEndpointParameters, *svcsdk.GetResolverEndpointOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.GetResolverEndpointOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.CreateResolverEndpointInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.ResolverEndpoint, _ *svcsdk.CreateResolverEndpointOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.DeleteResolverEndpointInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.ResolverEndpoint, _ *svcsdk.DeleteResolverEndpointOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.ResolverEndpoint, *svcsdk.UpdateResolverEndpointInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.ResolverEndpoint, _ *svcsdk.UpdateResolverEndpointOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
