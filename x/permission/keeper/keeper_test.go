package keeper_test

import (
	"math/rand"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/evmos/evmos/v12/testutil/sample"
	"github.com/evmos/evmos/v12/x/permission/types"
)

func (s *TestSuite) TestPruneAccountPolicies() {
	now := s.ctx.BlockTime()
	oneDayAfter := now.AddDate(0, 0, 1)

	resourceIDs := []math.Uint{math.NewUint(rand.Uint64()), math.NewUint(rand.Uint64()), math.NewUint(rand.Uint64())} //nolint: gosec
	policyIDs := make([]math.Uint, 3)

	// policy without expiry
	policy := types.Policy{
		Principal: &types.Principal{
			Type:  types.PRINCIPAL_TYPE_GNFD_ACCOUNT,
			Value: sample.RandAccAddressHex(),
		},
		ResourceType:   1,
		ResourceId:     resourceIDs[0],
		Statements:     nil,
		ExpirationTime: nil,
	}
	policyID, err := s.permissionKeeper.PutPolicy(s.ctx, &policy)
	s.NoError(err)
	policyIDs[0] = policyID

	policy.ResourceId = resourceIDs[2]
	policyID, err = s.permissionKeeper.PutPolicy(s.ctx, &policy)
	s.NoError(err)
	policyIDs[2] = policyID

	// policy with expiry
	policy.ResourceId = resourceIDs[1]
	policy.ExpirationTime = &oneDayAfter
	policyID, err = s.permissionKeeper.PutPolicy(s.ctx, &policy)
	s.NoError(err)
	policyIDs[1] = policyID

	testCases := []struct {
		name       string
		ctx        sdk.Context
		resourceID math.Uint
		policyID   math.Uint
		found      bool
		preRun     func()
		postRun    func()
	}{
		{
			name:       "no expiry and no prune",
			ctx:        s.ctx.WithBlockTime(oneDayAfter),
			resourceID: resourceIDs[0],
			policyID:   policyIDs[0],
			found:      true,
		},
		{
			name:       "expiry and no prune",
			ctx:        s.ctx.WithBlockTime(oneDayAfter),
			resourceID: resourceIDs[1],
			policyID:   policyIDs[1],
			found:      true,
		},
		{
			name:       "expiry and prune",
			ctx:        s.ctx.WithBlockTime(oneDayAfter.Add(time.Second)),
			resourceID: resourceIDs[1],
			policyID:   policyIDs[1],
		},
		{
			name:       "update from no expiry to expiry and prune",
			ctx:        s.ctx.WithBlockTime(oneDayAfter.Add(time.Second)),
			resourceID: resourceIDs[0],
			policyID:   policyIDs[0],
			preRun: func() {
				oldPolicy, found := s.permissionKeeper.GetPolicyByID(s.ctx, policyIDs[0])
				s.True(found)
				oldPolicy.ExpirationTime = &oneDayAfter
				newID, err := s.permissionKeeper.PutPolicy(s.ctx, oldPolicy)
				s.NoError(err)
				s.Equal(policyIDs[0], newID)
			},
		},
		{
			name:       "update from expiry to no expiry and no prune",
			ctx:        s.ctx.WithBlockTime(oneDayAfter.Add(time.Second)),
			resourceID: resourceIDs[2],
			policyID:   policyIDs[2],
			found:      true,
			preRun: func() {
				oldPolicy, found := s.permissionKeeper.GetPolicyByID(s.ctx, policyIDs[2])
				s.True(found)
				oldPolicy.ExpirationTime = nil
				newID, err := s.permissionKeeper.PutPolicy(s.ctx, oldPolicy)
				s.NoError(err)
				s.Equal(policyIDs[2], newID)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			if tc.preRun != nil {
				tc.preRun()
			}
			_, found := s.permissionKeeper.GetPolicyByID(tc.ctx, tc.policyID)
			s.True(found)
			s.permissionKeeper.RemoveExpiredPolicies(tc.ctx)
			_, found = s.permissionKeeper.GetPolicyByID(tc.ctx, tc.policyID)
			s.Equal(tc.found, found)
			if tc.postRun != nil {
				tc.postRun()
			}
		})
	}
}

func (s *TestSuite) TestPruneGroupPolicies() {
	now := s.ctx.BlockTime()
	oneDayAfter := now.AddDate(0, 0, 1)

	resourceIDs := []math.Uint{math.NewUint(rand.Uint64()), math.NewUint(rand.Uint64()), math.NewUint(rand.Uint64())} //nolint: gosec
	policyIDs := make([]math.Uint, 3)

	// member without expiry
	policy := types.Policy{
		Principal: &types.Principal{
			Type:  types.PRINCIPAL_TYPE_GNFD_GROUP,
			Value: sample.RandAccAddressHex(),
		},
		ResourceType:   1,
		ResourceId:     resourceIDs[0],
		Statements:     nil,
		ExpirationTime: nil,
	}
	policyID, err := s.permissionKeeper.PutPolicy(s.ctx, &policy)
	s.NoError(err)
	policyIDs[0] = policyID

	policy.ResourceId = resourceIDs[2]
	policyID, err = s.permissionKeeper.PutPolicy(s.ctx, &policy)
	s.NoError(err)
	policyIDs[2] = policyID

	// member with expiry
	policy.ResourceId = resourceIDs[1]
	policy.ExpirationTime = &oneDayAfter
	policyID, err = s.permissionKeeper.PutPolicy(s.ctx, &policy)
	s.NoError(err)
	policyIDs[1] = policyID

	testCases := []struct {
		name       string
		ctx        sdk.Context
		resourceID math.Uint
		policyID   math.Uint
		found      bool
		preRun     func()
		postRun    func()
	}{
		{
			name:       "no expiry and no prune",
			ctx:        s.ctx.WithBlockTime(oneDayAfter),
			resourceID: resourceIDs[0],
			policyID:   policyIDs[0],
			found:      true,
		},
		{
			name:       "expiry and no prune",
			ctx:        s.ctx.WithBlockTime(oneDayAfter),
			resourceID: resourceIDs[1],
			policyID:   policyIDs[1],
			found:      true,
		},
		{
			name:       "expiry and prune",
			ctx:        s.ctx.WithBlockTime(oneDayAfter.Add(time.Second)),
			resourceID: resourceIDs[1],
			policyID:   policyIDs[1],
		},
		{
			name:       "update from no expiry to expiry and prune",
			ctx:        s.ctx.WithBlockTime(oneDayAfter.Add(time.Second)),
			resourceID: resourceIDs[0],
			policyID:   policyIDs[0],
			preRun: func() {
				oldPolicy, found := s.permissionKeeper.GetPolicyByID(s.ctx, policyIDs[0])
				s.True(found)
				oldPolicy.ExpirationTime = &oneDayAfter
				newID, err := s.permissionKeeper.PutPolicy(s.ctx, oldPolicy)
				s.NoError(err)
				s.Equal(policyIDs[0], newID)
			},
		},
		{
			name:       "update from expiry to no expiry and no prune",
			ctx:        s.ctx.WithBlockTime(oneDayAfter.Add(time.Second)),
			resourceID: resourceIDs[2],
			policyID:   policyIDs[2],
			found:      true,
			preRun: func() {
				oldPolicy, found := s.permissionKeeper.GetPolicyByID(s.ctx, policyIDs[2])
				s.True(found)
				oldPolicy.ExpirationTime = nil
				newID, err := s.permissionKeeper.PutPolicy(s.ctx, oldPolicy)
				s.NoError(err)
				s.Equal(policyIDs[2], newID)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			if tc.preRun != nil {
				tc.preRun()
			}
			_, found := s.permissionKeeper.GetPolicyByID(tc.ctx, tc.policyID)
			s.True(found)
			s.permissionKeeper.RemoveExpiredPolicies(tc.ctx)
			_, found = s.permissionKeeper.GetPolicyByID(tc.ctx, tc.policyID)
			s.Equal(tc.found, found)
			if tc.postRun != nil {
				tc.postRun()
			}
		})
	}
}
