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

package accelerator

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/globalaccelerator"
	svcsdk "github.com/aws/aws-sdk-go/service/globalaccelerator"
	svcsdkapi "github.com/aws/aws-sdk-go/service/globalaccelerator/globalacceleratoriface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/globalaccelerator/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an Accelerator resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Accelerator in AWS"
	errUpdate        = "cannot update Accelerator in AWS"
	errDescribe      = "failed to describe Accelerator"
	errDelete        = "failed to delete Accelerator"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Accelerator)
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
	cr, ok := mg.(*svcapitypes.Accelerator)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeAcceleratorInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeAcceleratorWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateAccelerator(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

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
	cr, ok := mg.(*svcapitypes.Accelerator)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateAcceleratorInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateAcceleratorWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.Accelerator.AcceleratorArn != nil {
		cr.Status.AtProvider.AcceleratorARN = resp.Accelerator.AcceleratorArn
	} else {
		cr.Status.AtProvider.AcceleratorARN = nil
	}
	if resp.Accelerator.CreatedTime != nil {
		cr.Status.AtProvider.CreatedTime = &metav1.Time{*resp.Accelerator.CreatedTime}
	} else {
		cr.Status.AtProvider.CreatedTime = nil
	}
	if resp.Accelerator.DnsName != nil {
		cr.Status.AtProvider.DNSName = resp.Accelerator.DnsName
	} else {
		cr.Status.AtProvider.DNSName = nil
	}
	if resp.Accelerator.DualStackDnsName != nil {
		cr.Status.AtProvider.DualStackDNSName = resp.Accelerator.DualStackDnsName
	} else {
		cr.Status.AtProvider.DualStackDNSName = nil
	}
	if resp.Accelerator.Enabled != nil {
		cr.Spec.ForProvider.Enabled = resp.Accelerator.Enabled
	} else {
		cr.Spec.ForProvider.Enabled = nil
	}
	if resp.Accelerator.Events != nil {
		f5 := []*svcapitypes.AcceleratorEvent{}
		for _, f5iter := range resp.Accelerator.Events {
			f5elem := &svcapitypes.AcceleratorEvent{}
			if f5iter.Message != nil {
				f5elem.Message = f5iter.Message
			}
			if f5iter.Timestamp != nil {
				f5elem.Timestamp = &metav1.Time{*f5iter.Timestamp}
			}
			f5 = append(f5, f5elem)
		}
		cr.Status.AtProvider.Events = f5
	} else {
		cr.Status.AtProvider.Events = nil
	}
	if resp.Accelerator.IpAddressType != nil {
		cr.Spec.ForProvider.IPAddressType = resp.Accelerator.IpAddressType
	} else {
		cr.Spec.ForProvider.IPAddressType = nil
	}
	if resp.Accelerator.IpSets != nil {
		f7 := []*svcapitypes.IPSet{}
		for _, f7iter := range resp.Accelerator.IpSets {
			f7elem := &svcapitypes.IPSet{}
			if f7iter.IpAddressFamily != nil {
				f7elem.IPAddressFamily = f7iter.IpAddressFamily
			}
			if f7iter.IpAddresses != nil {
				f7elemf1 := []*string{}
				for _, f7elemf1iter := range f7iter.IpAddresses {
					var f7elemf1elem string
					f7elemf1elem = *f7elemf1iter
					f7elemf1 = append(f7elemf1, &f7elemf1elem)
				}
				f7elem.IPAddresses = f7elemf1
			}
			if f7iter.IpFamily != nil {
				f7elem.IPFamily = f7iter.IpFamily
			}
			f7 = append(f7, f7elem)
		}
		cr.Status.AtProvider.IPSets = f7
	} else {
		cr.Status.AtProvider.IPSets = nil
	}
	if resp.Accelerator.LastModifiedTime != nil {
		cr.Status.AtProvider.LastModifiedTime = &metav1.Time{*resp.Accelerator.LastModifiedTime}
	} else {
		cr.Status.AtProvider.LastModifiedTime = nil
	}
	if resp.Accelerator.Name != nil {
		cr.Spec.ForProvider.Name = resp.Accelerator.Name
	} else {
		cr.Spec.ForProvider.Name = nil
	}
	if resp.Accelerator.Status != nil {
		cr.Status.AtProvider.Status = resp.Accelerator.Status
	} else {
		cr.Status.AtProvider.Status = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.Accelerator)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateAcceleratorInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateAcceleratorWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.Accelerator)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteAcceleratorInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteAcceleratorWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.GlobalAcceleratorAPI, opts []option) *external {
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
	client         svcsdkapi.GlobalAcceleratorAPI
	preObserve     func(context.Context, *svcapitypes.Accelerator, *svcsdk.DescribeAcceleratorInput) error
	postObserve    func(context.Context, *svcapitypes.Accelerator, *svcsdk.DescribeAcceleratorOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.AcceleratorParameters, *svcsdk.DescribeAcceleratorOutput) error
	isUpToDate     func(context.Context, *svcapitypes.Accelerator, *svcsdk.DescribeAcceleratorOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.Accelerator, *svcsdk.CreateAcceleratorInput) error
	postCreate     func(context.Context, *svcapitypes.Accelerator, *svcsdk.CreateAcceleratorOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Accelerator, *svcsdk.DeleteAcceleratorInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.Accelerator, *svcsdk.DeleteAcceleratorOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.Accelerator, *svcsdk.UpdateAcceleratorInput) error
	postUpdate     func(context.Context, *svcapitypes.Accelerator, *svcsdk.UpdateAcceleratorOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Accelerator, *svcsdk.DescribeAcceleratorInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.Accelerator, _ *svcsdk.DescribeAcceleratorOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.AcceleratorParameters, *svcsdk.DescribeAcceleratorOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.Accelerator, *svcsdk.DescribeAcceleratorOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.Accelerator, *svcsdk.CreateAcceleratorInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.Accelerator, _ *svcsdk.CreateAcceleratorOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.Accelerator, *svcsdk.DeleteAcceleratorInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.Accelerator, _ *svcsdk.DeleteAcceleratorOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.Accelerator, *svcsdk.UpdateAcceleratorInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.Accelerator, _ *svcsdk.UpdateAcceleratorOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
