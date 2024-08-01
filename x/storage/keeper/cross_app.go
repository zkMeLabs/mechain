package keeper

import (
	"github.com/evmos/evmos/v12/x/storage/types"
)

func RegisterCrossApps(keeper Keeper) {
	bucketApp := NewBucketApp(keeper)
	err := keeper.crossChainKeeper.RegisterChannel(types.BucketChannel, types.BucketChannelId, bucketApp)
	if err != nil {
		panic(err)
	}

	objectApp := NewObjectApp(keeper)
	err = keeper.crossChainKeeper.RegisterChannel(types.ObjectChannel, types.ObjectChannelId, objectApp)
	if err != nil {
		panic(err)
	}

	groupApp := NewGroupApp(keeper)
	err = keeper.crossChainKeeper.RegisterChannel(types.GroupChannel, types.GroupChannelId, groupApp)
	if err != nil {
		panic(err)
	}

	zkmesbtApp := NewZkmeSBTApp(keeper)
	err = keeper.crossChainKeeper.RegisterChannel(types.ZkmeSBTChannel, types.ZkmeSBTChannelId, zkmesbtApp)
	if err != nil {
		panic(err)
	}

	permissionApp := NewPermissionApp(keeper, keeper.permKeeper)
	err = keeper.crossChainKeeper.RegisterChannel(types.PermissionChannel, types.PermissionChannelId, permissionApp)
	if err != nil {
		panic(err)
	}
}
