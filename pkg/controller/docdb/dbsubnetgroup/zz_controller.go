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

package dbsubnetgroup

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/docdb"
	svcsdk "github.com/aws/aws-sdk-go/service/docdb"
	svcsdkapi "github.com/aws/aws-sdk-go/service/docdb/docdbiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/docdb/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an DBSubnetGroup resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create DBSubnetGroup in AWS"
	errUpdate        = "cannot update DBSubnetGroup in AWS"
	errDescribe      = "failed to describe DBSubnetGroup"
	errDelete        = "failed to delete DBSubnetGroup"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.DBSubnetGroup)
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
	cr, ok := mg.(*svcapitypes.DBSubnetGroup)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeDBSubnetGroupsInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeDBSubnetGroupsWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.DBSubnetGroups) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateDBSubnetGroup(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

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
	cr, ok := mg.(*svcapitypes.DBSubnetGroup)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateDBSubnetGroupInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateDBSubnetGroupWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.DBSubnetGroup.DBSubnetGroupArn != nil {
		cr.Status.AtProvider.DBSubnetGroupARN = resp.DBSubnetGroup.DBSubnetGroupArn
	} else {
		cr.Status.AtProvider.DBSubnetGroupARN = nil
	}
	if resp.DBSubnetGroup.DBSubnetGroupDescription != nil {
		cr.Spec.ForProvider.DBSubnetGroupDescription = resp.DBSubnetGroup.DBSubnetGroupDescription
	} else {
		cr.Spec.ForProvider.DBSubnetGroupDescription = nil
	}
	if resp.DBSubnetGroup.DBSubnetGroupName != nil {
		cr.Status.AtProvider.DBSubnetGroupName = resp.DBSubnetGroup.DBSubnetGroupName
	} else {
		cr.Status.AtProvider.DBSubnetGroupName = nil
	}
	if resp.DBSubnetGroup.SubnetGroupStatus != nil {
		cr.Status.AtProvider.SubnetGroupStatus = resp.DBSubnetGroup.SubnetGroupStatus
	} else {
		cr.Status.AtProvider.SubnetGroupStatus = nil
	}
	if resp.DBSubnetGroup.Subnets != nil {
		f4 := []*svcapitypes.Subnet{}
		for _, f4iter := range resp.DBSubnetGroup.Subnets {
			f4elem := &svcapitypes.Subnet{}
			if f4iter.SubnetAvailabilityZone != nil {
				f4elemf0 := &svcapitypes.AvailabilityZone{}
				if f4iter.SubnetAvailabilityZone.Name != nil {
					f4elemf0.Name = f4iter.SubnetAvailabilityZone.Name
				}
				f4elem.SubnetAvailabilityZone = f4elemf0
			}
			if f4iter.SubnetIdentifier != nil {
				f4elem.SubnetIdentifier = f4iter.SubnetIdentifier
			}
			if f4iter.SubnetStatus != nil {
				f4elem.SubnetStatus = f4iter.SubnetStatus
			}
			f4 = append(f4, f4elem)
		}
		cr.Status.AtProvider.Subnets = f4
	} else {
		cr.Status.AtProvider.Subnets = nil
	}
	if resp.DBSubnetGroup.VpcId != nil {
		cr.Status.AtProvider.VPCID = resp.DBSubnetGroup.VpcId
	} else {
		cr.Status.AtProvider.VPCID = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.DBSubnetGroup)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateModifyDBSubnetGroupInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.ModifyDBSubnetGroupWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.DBSubnetGroup)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteDBSubnetGroupInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteDBSubnetGroupWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.DocDBAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		filterList:     nopFilterList,
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
	client         svcsdkapi.DocDBAPI
	preObserve     func(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.DescribeDBSubnetGroupsInput) error
	postObserve    func(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.DescribeDBSubnetGroupsOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.DBSubnetGroup, *svcsdk.DescribeDBSubnetGroupsOutput) *svcsdk.DescribeDBSubnetGroupsOutput
	lateInitialize func(*svcapitypes.DBSubnetGroupParameters, *svcsdk.DescribeDBSubnetGroupsOutput) error
	isUpToDate     func(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.DescribeDBSubnetGroupsOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.CreateDBSubnetGroupInput) error
	postCreate     func(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.CreateDBSubnetGroupOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.DeleteDBSubnetGroupInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.DeleteDBSubnetGroupOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.ModifyDBSubnetGroupInput) error
	postUpdate     func(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.ModifyDBSubnetGroupOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.DescribeDBSubnetGroupsInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.DBSubnetGroup, _ *svcsdk.DescribeDBSubnetGroupsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.DBSubnetGroup, list *svcsdk.DescribeDBSubnetGroupsOutput) *svcsdk.DescribeDBSubnetGroupsOutput {
	return list
}

func nopLateInitialize(*svcapitypes.DBSubnetGroupParameters, *svcsdk.DescribeDBSubnetGroupsOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.DescribeDBSubnetGroupsOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.CreateDBSubnetGroupInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.DBSubnetGroup, _ *svcsdk.CreateDBSubnetGroupOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.DeleteDBSubnetGroupInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.DBSubnetGroup, _ *svcsdk.DeleteDBSubnetGroupOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.DBSubnetGroup, *svcsdk.ModifyDBSubnetGroupInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.DBSubnetGroup, _ *svcsdk.ModifyDBSubnetGroupOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
