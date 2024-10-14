// Copyright 2022 Evmos Foundation
// This file is part of the Evmos Network packages.
//
// Evmos is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Evmos packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Evmos packages. If not, see https://github.com/evmos/evmos/blob/main/LICENSE

package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/ethereum/go-ethereum/core/vm"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"

	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cast"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/simapp"
	simappparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/store/iavl"
	"github.com/cosmos/cosmos-sdk/store/streaming"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/posthandler"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/crosschain"
	crosschainkeeper "github.com/cosmos/cosmos-sdk/x/crosschain/keeper"
	crosschaintypes "github.com/cosmos/cosmos-sdk/x/crosschain/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/gashub"
	gashubkeeper "github.com/cosmos/cosmos-sdk/x/gashub/keeper"
	gashubtypes "github.com/cosmos/cosmos-sdk/x/gashub/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/oracle"
	oraclekeeper "github.com/cosmos/cosmos-sdk/x/oracle/keeper"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctestingtypes "github.com/cosmos/ibc-go/v7/testing/types"
	bridgemodule "github.com/evmos/evmos/v12/x/bridge"

	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v7/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"

	ica "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"
	icahost "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/eth/eip712"
	ethante "github.com/evmos/evmos/v12/app/ante/evm"
	"github.com/evmos/evmos/v12/encoding"
	servercfg "github.com/evmos/evmos/v12/server/config"
	srvflags "github.com/evmos/evmos/v12/server/flags"
	evmostypes "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/evm"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"
	precompilesauthz "github.com/evmos/evmos/v12/x/evm/precompiles/authz"
	precompilesbank "github.com/evmos/evmos/v12/x/evm/precompiles/bank"
	precompilesgov "github.com/evmos/evmos/v12/x/evm/precompiles/gov"
	precompilesstorage "github.com/evmos/evmos/v12/x/evm/precompiles/storage"
	precompilesvirtualgroup "github.com/evmos/evmos/v12/x/evm/precompiles/virtualgroup"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
	"github.com/evmos/evmos/v12/x/feemarket"
	feemarketkeeper "github.com/evmos/evmos/v12/x/feemarket/keeper"
	feemarkettypes "github.com/evmos/evmos/v12/x/feemarket/types"

	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	// unnamed import of statik for swagger UI support
	_ "github.com/evmos/evmos/v12/client/docs/statik"

	"github.com/evmos/evmos/v12/app/ante"
	"github.com/evmos/evmos/v12/x/epochs"
	epochskeeper "github.com/evmos/evmos/v12/x/epochs/keeper"
	epochstypes "github.com/evmos/evmos/v12/x/epochs/types"
	"github.com/evmos/evmos/v12/x/erc20"
	erc20client "github.com/evmos/evmos/v12/x/erc20/client"
	erc20keeper "github.com/evmos/evmos/v12/x/erc20/keeper"
	erc20types "github.com/evmos/evmos/v12/x/erc20/types"

	// NOTE: override ICS20 keeper to support IBC transfers of ERC20 tokens
	"github.com/evmos/evmos/v12/x/ibc/transfer"
	transferkeeper "github.com/evmos/evmos/v12/x/ibc/transfer/keeper"

	// Force-load the tracer engines to trigger registration due to Go-Ethereum v1.10.15 changes
	_ "github.com/ethereum/go-ethereum/eth/tracers/js"
	_ "github.com/ethereum/go-ethereum/eth/tracers/native"

	bridgemodulekeeper "github.com/evmos/evmos/v12/x/bridge/keeper"
	bridgemoduletypes "github.com/evmos/evmos/v12/x/bridge/types"
	challengemodule "github.com/evmos/evmos/v12/x/challenge"
	challengemodulekeeper "github.com/evmos/evmos/v12/x/challenge/keeper"
	challengemoduletypes "github.com/evmos/evmos/v12/x/challenge/types"
	"github.com/evmos/evmos/v12/x/gensp"
	gensptypes "github.com/evmos/evmos/v12/x/gensp/types"
	paymentmodule "github.com/evmos/evmos/v12/x/payment"
	paymentmodulekeeper "github.com/evmos/evmos/v12/x/payment/keeper"
	paymentmoduletypes "github.com/evmos/evmos/v12/x/payment/types"
	permissionmodule "github.com/evmos/evmos/v12/x/permission"
	permissionmodulekeeper "github.com/evmos/evmos/v12/x/permission/keeper"
	permissionmoduletypes "github.com/evmos/evmos/v12/x/permission/types"
	spmodule "github.com/evmos/evmos/v12/x/sp"
	spmodulekeeper "github.com/evmos/evmos/v12/x/sp/keeper"
	spmoduletypes "github.com/evmos/evmos/v12/x/sp/types"
	storagemodule "github.com/evmos/evmos/v12/x/storage"
	storagemodulekeeper "github.com/evmos/evmos/v12/x/storage/keeper"
	storagemoduletypes "github.com/evmos/evmos/v12/x/storage/types"
	virtualgroupmodule "github.com/evmos/evmos/v12/x/virtualgroup"
	virtualgroupmodulekeeper "github.com/evmos/evmos/v12/x/virtualgroup/keeper"
	virtualgroupmoduletypes "github.com/evmos/evmos/v12/x/virtualgroup/types"
)

// Name defines the application binary name
const (
	Name      = "mechaind"
	ShortName = "mechaind"
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		gensp.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				paramsclient.ProposalHandler,
				ibcclientclient.UpdateClientProposalHandler, ibcclientclient.UpgradeProposalHandler,
				// Evmos proposal types
				erc20client.RegisterCoinProposalHandler, erc20client.RegisterERC20ProposalHandler, erc20client.ToggleTokenConversionProposalHandler,
			},
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		consensus.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ibctm.AppModuleBasic{},
		ica.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		crosschain.AppModuleBasic{},
		oracle.AppModuleBasic{},
		bridgemodule.AppModuleBasic{},

		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{AppModuleBasic: &ibctransfer.AppModuleBasic{}},
		evm.AppModuleBasic{},
		feemarket.AppModuleBasic{},
		erc20.AppModuleBasic{},
		epochs.AppModuleBasic{},
		gashub.AppModuleBasic{},
		spmodule.AppModuleBasic{},
		paymentmodule.AppModuleBasic{},
		permissionmodule.AppModuleBasic{},
		virtualgroupmodule.AppModuleBasic{},
		storagemodule.AppModuleBasic{},
		challengemodule.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:         nil,
		distrtypes.ModuleName:              nil,
		stakingtypes.BondedPoolName:        {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName:     {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:                {authtypes.Burner},
		ibctransfertypes.ModuleName:        {authtypes.Minter, authtypes.Burner},
		icatypes.ModuleName:                nil,
		evmtypes.ModuleName:                {authtypes.Minter, authtypes.Burner}, // used for secure addition and subtraction of balance using module account
		erc20types.ModuleName:              {authtypes.Minter, authtypes.Burner},
		paymentmoduletypes.ModuleName:      {authtypes.Burner, authtypes.Staking},
		crosschaintypes.ModuleName:         {authtypes.Minter},
		permissionmoduletypes.ModuleName:   nil,
		bridgemoduletypes.ModuleName:       nil,
		spmoduletypes.ModuleName:           {authtypes.Staking},
		virtualgroupmoduletypes.ModuleName: nil,
	}
)

var (
	_ servertypes.Application = (*Evmos)(nil)
	_ ibctesting.TestingApp   = (*Evmos)(nil)
	_ runtime.AppI            = (*Evmos)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+ShortName)

	// manually update the power reduction by replacing micro (u) -> atto (a) evmos
	sdk.DefaultPowerReduction = evmostypes.PowerReduction
	// modify fee market parameter defaults through global
	feemarkettypes.DefaultMinGasPrice = MainnetMinGasPrices
	feemarkettypes.DefaultMinGasMultiplier = MainnetMinGasMultiplier
	// modify default min commission to 5%
	stakingtypes.DefaultMinCommissionRate = sdk.NewDecWithPrec(5, 2)
}

// Evmos implements an extended ABCI application. It is an application
// that may process transactions through Ethereum's EVM running atop of
// Tendermint consensus.
type Evmos struct {
	*baseapp.BaseApp

	// encoding
	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	AuthzKeeper           authzkeeper.Keeper
	BankKeeper            bankkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	GovKeeper             govkeeper.Keeper
	CrisisKeeper          crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	CrossChainKeeper      crosschainkeeper.Keeper
	OracleKeeper          oraclekeeper.Keeper
	GashubKeeper          gashubkeeper.Keeper
	IBCKeeper             *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	ICAHostKeeper         icahostkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	TransferKeeper        transferkeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper

	BridgeKeeper           bridgemodulekeeper.Keeper
	SpKeeper               spmodulekeeper.Keeper
	PaymentKeeper          paymentmodulekeeper.Keeper
	ChallengeKeeper        challengemodulekeeper.Keeper
	PermissionmoduleKeeper permissionmodulekeeper.Keeper
	VirtualgroupKeeper     virtualgroupmodulekeeper.Keeper
	StorageKeeper          storagemodulekeeper.Keeper
	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper

	// Ethermint keepers
	EvmKeeper       *evmkeeper.Keeper
	FeeMarketKeeper feemarketkeeper.Keeper

	// Evmos keepers
	Erc20Keeper  erc20keeper.Keeper
	EpochsKeeper epochskeeper.Keeper

	// the module manager
	mm *module.Manager

	// the configurator
	configurator module.Configurator

	// simulation manager
	sm *module.SimulationManager

	tpsCounter *tpsCounter
	// app config
	appConfig *servercfg.AppConfig
}

// SimulationManager implements runtime.AppI
func (app *Evmos) SimulationManager() *module.SimulationManager {
	return app.sm
}

// NewEvmos returns a reference to a new initialized Ethermint application.
func NewEvmos(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig simappparams.EncodingConfig,
	customAppConfig *servercfg.AppConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *Evmos {
	appCodec := encodingConfig.Codec
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	eip712.AminoCodec = encodingConfig.Amino
	eip712.ProtoCodec = codec.NewProtoCodec(encodingConfig.InterfaceRegistry)

	// Setup Mempool and Proposal Handlers
	baseAppOptions = append(baseAppOptions, func(app *baseapp.BaseApp) {
		mempool := mempool.NoOpMempool{}
		app.SetMempool(mempool)
		handler := baseapp.NewDefaultProposalHandler(mempool, app)
		app.SetPrepareProposal(handler.PrepareProposalHandler())
		app.SetProcessProposal(handler.ProcessProposalHandler())
	})

	// NOTE we use custom transaction decoder that supports the sdk.Tx interface instead of sdk.StdTx
	bApp := baseapp.NewBaseApp(
		Name,
		logger,
		db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...,
	)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		// SDK keys
		authtypes.StoreKey, authzkeeper.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, upgradetypes.StoreKey,
		evidencetypes.StoreKey, capabilitytypes.StoreKey, consensusparamtypes.StoreKey,
		feegrant.StoreKey, crisistypes.StoreKey,
		group.StoreKey,
		crosschaintypes.StoreKey,
		oracletypes.StoreKey,
		bridgemoduletypes.StoreKey,
		gashubtypes.StoreKey,
		spmoduletypes.StoreKey,
		virtualgroupmoduletypes.StoreKey,
		paymentmoduletypes.StoreKey,
		permissionmoduletypes.StoreKey,
		storagemoduletypes.StoreKey,
		challengemoduletypes.StoreKey,
		reconStoreKey,
		// ibc keys
		ibcexported.StoreKey, ibctransfertypes.StoreKey,
		// ica keys
		icahosttypes.StoreKey,
		// ethermint keys
		evmtypes.StoreKey, feemarkettypes.StoreKey,
		// evmos keys
		erc20types.StoreKey,
		epochstypes.StoreKey,
	)

	// Add the EVM transient store key
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey, evmtypes.TransientKey, feemarkettypes.TransientKey, challengemoduletypes.TStoreKey, storagemoduletypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey, challengemoduletypes.MemStoreKey)

	// load state streaming if enabled
	if _, _, err := streaming.LoadStreamingServices(bApp, appOpts, appCodec, logger, keys); err != nil {
		fmt.Printf("failed to load state streaming: %s", err)
		os.Exit(1)
	}

	app := &Evmos{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		appConfig:         customAppConfig,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	// init params keeper and subspaces
	app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// get authority address
	authAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	// set the BaseApp's parameter store
	app.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(appCodec, keys[consensusparamtypes.StoreKey], authAddr)
	bApp.SetParamStore(&app.ConsensusParamsKeeper)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)

	// Applications that wish to enforce statically created ScopedKeepers should call `Seal` after creating
	// their scoped modules in `NewApp` with `ScopeToModule`
	app.CapabilityKeeper.Seal()

	// use custom Ethermint account for contracts
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey],
		evmostypes.ProtoAccount, maccPerms,
		authAddr,
	)
	app.AuthzKeeper = authzkeeper.NewKeeper(keys[authzkeeper.StoreKey], appCodec, app.MsgServiceRouter(), app.AccountKeeper)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.BlockedAddrs(), authAddr,
	)
	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec, keys[stakingtypes.StoreKey], app.AccountKeeper, app.AuthzKeeper, app.BankKeeper, authAddr,
	)
	app.CrossChainKeeper = crosschainkeeper.NewKeeper(
		appCodec,
		keys[crosschaintypes.StoreKey],
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		app.StakingKeeper,
		app.BankKeeper,
	)
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], app.AccountKeeper, app.BankKeeper,
		app.StakingKeeper, authtypes.FeeCollectorName, authAddr,
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, app.LegacyAmino(), keys[slashingtypes.StoreKey], app.StakingKeeper, authAddr,
	)
	app.OracleKeeper = oraclekeeper.NewKeeper(
		appCodec,
		keys[crosschaintypes.StoreKey],
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		app.CrossChainKeeper,
		app.BankKeeper,
		app.StakingKeeper,
	)
	app.CrisisKeeper = *crisiskeeper.NewKeeper(
		appCodec, keys[crisistypes.StoreKey], invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName, authAddr,
	)
	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.AccountKeeper)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp)

	tracer := cast.ToString(appOpts.Get(srvflags.EVMTracer))

	// Create Ethermint keepers
	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		appCodec, authtypes.NewModuleAddress(govtypes.ModuleName),
		keys[feemarkettypes.StoreKey],
		tkeys[feemarkettypes.TransientKey],
		app.GetSubspace(feemarkettypes.ModuleName),
	)

	app.EvmKeeper = evmkeeper.NewKeeper(
		appCodec, keys[evmtypes.StoreKey], tkeys[evmtypes.TransientKey], authtypes.NewModuleAddress(govtypes.ModuleName),
		app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.FeeMarketKeeper,
		tracer, app.GetSubspace(evmtypes.ModuleName),
	)

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibcexported.StoreKey], app.GetSubspace(ibcexported.ModuleName), app.StakingKeeper, app.UpgradeKeeper, scopedIBCKeeper,
	)

	// register the proposal types
	govRouter := govv1beta1.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper)).
		AddRoute(erc20types.RouterKey, erc20.NewErc20ProposalHandler(&app.Erc20Keeper))

	govConfig := govtypes.DefaultConfig()
	/*
		Example of setting gov params:
		govConfig.MaxMetadataLen = 10000
	*/
	govKeeper := govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], app.AccountKeeper, app.BankKeeper,
		app.StakingKeeper, app.CrossChainKeeper, app.MsgServiceRouter(), govConfig, authAddr,
	)

	// Set legacy router for backwards compatibility with gov v1beta1
	govKeeper.SetLegacyRouter(govRouter)

	// Evmos Keeper

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	// NOTE: Distr, Slashing and Claim must be created before calling the Hooks method to avoid returning a Keeper without its table generated
	app.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			app.DistrKeeper.Hooks(),
			app.SlashingKeeper.Hooks(),
		),
	)

	app.Erc20Keeper = erc20keeper.NewKeeper(
		keys[erc20types.StoreKey], appCodec, authtypes.NewModuleAddress(govtypes.ModuleName),
		app.AccountKeeper, app.BankKeeper, app.EvmKeeper, app.StakingKeeper,
	)

	epochsKeeper := epochskeeper.NewKeeper(appCodec, keys[epochstypes.StoreKey])
	app.EpochsKeeper = *epochsKeeper.SetHooks(
		epochskeeper.NewMultiEpochHooks(
		// insert epoch hooks receivers here
		),
	)

	app.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(),
	)

	app.EvmKeeper = app.EvmKeeper.SetHooks(
		evmkeeper.NewMultiEvmHooks(
			app.Erc20Keeper.Hooks(),
		),
	)

	app.TransferKeeper = transferkeeper.NewKeeper(
		appCodec, keys[ibctransfertypes.StoreKey], app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper, // ICS4 Wrapper: claims IBC middleware
		app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		app.AccountKeeper, app.BankKeeper, scopedTransferKeeper,
		app.Erc20Keeper, // Add ERC20 Keeper for ERC20 transfers
	)

	// Override the ICS20 app module
	transferModule := transfer.NewAppModule(app.TransferKeeper)

	// Create the app.ICAHostKeeper
	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec, app.keys[icahosttypes.StoreKey],
		app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		scopedICAHostKeeper,
		bApp.MsgServiceRouter(),
	)

	// create host IBC module
	icaHostIBCModule := icahost.NewIBCModule(app.ICAHostKeeper)

	// create IBC module from top to bottom of stack
	var transferStack porttypes.IBCModule

	transferStack = transfer.NewIBCModule(app.TransferKeeper)
	transferStack = erc20.NewIBCMiddleware(app.Erc20Keeper, transferStack)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.
		AddRoute(icahosttypes.SubModuleName, icaHostIBCModule).
		AddRoute(ibctransfertypes.ModuleName, transferStack)

	app.IBCKeeper.SetRouter(ibcRouter)

	// create evidence keeper with router
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], app.StakingKeeper, app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	// greenfield keeper
	app.BridgeKeeper = *bridgemodulekeeper.NewKeeper(
		appCodec,
		keys[bridgemoduletypes.StoreKey],
		app.BankKeeper,
		app.StakingKeeper,
		app.CrossChainKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	bridgeModule := bridgemodule.NewAppModule(appCodec, app.BridgeKeeper, app.AccountKeeper, app.BankKeeper)

	app.GashubKeeper = gashubkeeper.NewKeeper(
		appCodec,
		keys[gashubtypes.StoreKey],
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	gashubModule := gashub.NewAppModule(app.GashubKeeper)

	app.SpKeeper = *spmodulekeeper.NewKeeper(
		appCodec,
		keys[spmoduletypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.AuthzKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	spModule := spmodule.NewAppModule(appCodec, app.SpKeeper, app.AccountKeeper, app.BankKeeper)

	app.PaymentKeeper = *paymentmodulekeeper.NewKeeper(
		appCodec,
		keys[paymentmoduletypes.StoreKey],
		app.BankKeeper,
		app.AccountKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	paymentModule := paymentmodule.NewAppModule(appCodec, app.PaymentKeeper, app.AccountKeeper, app.BankKeeper)

	app.VirtualgroupKeeper = *virtualgroupmodulekeeper.NewKeeper(
		appCodec,
		keys[virtualgroupmoduletypes.StoreKey],
		tkeys[virtualgroupmoduletypes.TStoreKey],
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		app.SpKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.PaymentKeeper,
	)

	app.PermissionmoduleKeeper = *permissionmodulekeeper.NewKeeper(
		appCodec,
		keys[permissionmoduletypes.StoreKey],
		app.AccountKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	permissionModule := permissionmodule.NewAppModule(appCodec, app.PermissionmoduleKeeper, app.AccountKeeper, app.BankKeeper)

	app.StorageKeeper = *storagemodulekeeper.NewKeeper(
		appCodec,
		keys[storagemoduletypes.StoreKey],
		tkeys[storagemoduletypes.TStoreKey],
		app.AccountKeeper,
		app.SpKeeper,
		app.PaymentKeeper,
		app.PermissionmoduleKeeper,
		app.CrossChainKeeper,
		app.VirtualgroupKeeper,
		app.EvmKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	storageModule := storagemodule.NewAppModule(appCodec, app.StorageKeeper, app.AccountKeeper, app.BankKeeper, app.SpKeeper)

	app.VirtualgroupKeeper.SetStorageKeeper(&app.StorageKeeper)
	virtualgroupModule := virtualgroupmodule.NewAppModule(appCodec, app.VirtualgroupKeeper, app.SpKeeper)

	app.ChallengeKeeper = *challengemodulekeeper.NewKeeper(
		appCodec,
		keys[challengemoduletypes.StoreKey],
		tkeys[challengemoduletypes.TStoreKey],
		app.BankKeeper,
		app.StorageKeeper,
		app.SpKeeper,
		app.StakingKeeper,
		app.PaymentKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	challengeModule := challengemodule.NewAppModule(appCodec, app.ChallengeKeeper, app.AccountKeeper, app.BankKeeper)
	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		// SDK app modules
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		gensp.NewAppModule(app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.PaymentKeeper, app.GetSubspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants, app.GetSubspace(crisistypes.ModuleName)),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, &app.GovKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(slashingtypes.ModuleName)),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
		crosschain.NewAppModule(app.CrossChainKeeper, app.BankKeeper, app.StakingKeeper),
		oracle.NewAppModule(app.OracleKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		params.NewAppModule(app.ParamsKeeper),
		consensus.NewAppModule(appCodec, app.ConsensusParamsKeeper),
		bridgeModule,
		gashubModule,
		spModule,
		virtualgroupModule,
		paymentModule,
		permissionModule,
		storageModule,
		challengeModule,

		// ibc modules
		ibc.NewAppModule(app.IBCKeeper),
		ica.NewAppModule(nil, &app.ICAHostKeeper),
		transferModule,
		// Ethermint app modules
		evm.NewAppModule(app.EvmKeeper, app.AccountKeeper, app.GetSubspace(evmtypes.ModuleName)),
		feemarket.NewAppModule(app.FeeMarketKeeper, app.GetSubspace(feemarkettypes.ModuleName)),
		// Evmos app modules
		erc20.NewAppModule(app.Erc20Keeper, app.AccountKeeper,
			app.GetSubspace(erc20types.ModuleName)),
		epochs.NewAppModule(appCodec, app.EpochsKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: upgrade module must go first to handle software upgrades.
	// NOTE: staking module is required if HistoricalEntries param > 0.
	// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		// Note: epochs' begin should be "real" start of epochs, we keep epochs beginblock at the beginning
		epochstypes.ModuleName,
		feemarkettypes.ModuleName,
		evmtypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		ibcexported.ModuleName,
		// no-op modules
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		consensusparamtypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		crosschaintypes.ModuleName,
		oracletypes.ModuleName,
		paramstypes.ModuleName,
		erc20types.ModuleName,
		bridgemoduletypes.ModuleName,
		gashubtypes.ModuleName,
		spmoduletypes.ModuleName,
		virtualgroupmoduletypes.ModuleName,
		paymentmoduletypes.ModuleName,
		permissionmoduletypes.ModuleName,
		storagemoduletypes.ModuleName,
		gensptypes.ModuleName,
		challengemoduletypes.ModuleName,
	)

	// NOTE: fee market module must go last in order to retrieve the block gas used.
	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		evmtypes.ModuleName,
		feemarkettypes.ModuleName,
		// Note: epochs' endblock should be "real" end of epochs, we keep epochs endblock at the end
		epochstypes.ModuleName,
		// no-op modules
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		consensusparamtypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		crosschaintypes.ModuleName,
		oracletypes.ModuleName,
		upgradetypes.ModuleName,
		// Evmos modules
		erc20types.ModuleName,
		bridgemoduletypes.ModuleName,
		gashubtypes.ModuleName,
		spmoduletypes.ModuleName,
		virtualgroupmoduletypes.ModuleName,
		paymentmoduletypes.ModuleName,
		permissionmoduletypes.ModuleName,
		storagemoduletypes.ModuleName,
		gensptypes.ModuleName,
		challengemoduletypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		// SDK modules
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		// NOTE: staking requires the claiming hook
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		gashubtypes.ModuleName,
		ibcexported.ModuleName,
		// Ethermint modules
		// evm module denomination is used by the revenue module, in AnteHandle
		evmtypes.ModuleName,
		// NOTE: feemarket module needs to be initialized before genutil module:
		// gentx transactions use MinGasPriceDecorator.AnteHandle
		feemarkettypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		consensusparamtypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		crosschaintypes.ModuleName,
		oracletypes.ModuleName,
		// Evmos modules
		erc20types.ModuleName,
		epochstypes.ModuleName,
		// NOTE: crisis module must go at the end to check for invariants on each module
		crisistypes.ModuleName,
		bridgemoduletypes.ModuleName,
		spmoduletypes.ModuleName,
		virtualgroupmoduletypes.ModuleName,
		paymentmoduletypes.ModuleName,
		permissionmoduletypes.ModuleName,
		storagemoduletypes.ModuleName,
		gensptypes.ModuleName,
		challengemoduletypes.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.mm.Modules))

	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// add test gRPC service for testing gRPC queries in isolation
	// testdata.RegisterTestServiceServer(app.GRPCQueryRouter(), testdata.TestServiceImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.mm.Modules, overrideModules)
	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	maxGasWanted := cast.ToUint64(appOpts.Get(srvflags.EVMMaxTxGasWanted))

	app.setAnteHandler(encodingConfig.TxConfig, maxGasWanted)
	app.setPostHandler()
	app.SetEndBlocker(app.EndBlocker)
	app.setupUpgradeHandlers()
	app.SetUpgradeChecker(app.UpgradeKeeper.IsUpgraded)
	app.EvmPrecompiled()

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	err = app.RegisterUpgradeHandlers(app.ChainID(), &app.appConfig.Config)
	if err != nil {
		panic(err)
	}
	ms := app.CommitMultiStore()
	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: app.ChainID(), Height: app.LastBlockHeight()}, true, app.UpgradeKeeper.IsUpgraded, app.Logger())
	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			logger.Error("error on loading last version", "err", err)
			os.Exit(1)
		}
		// Execute the upgraded register, such as the newly added Msg type
		// ex.
		// app.GovKeeper.Router().RegisterService(...)
		err = app.UpgradeKeeper.InitUpgraded(ctx)
		if err != nil {
			panic(err)
		}
	}
	if app.IsIavlStore() {
		// enable diff for reconciliation
		bankIavl, ok := ms.GetCommitStore(keys[banktypes.StoreKey]).(*iavl.Store)
		if !ok {
			os.Exit(1)
		}
		bankIavl.EnableDiff()
		paymentIavl, ok := ms.GetCommitStore(keys[paymentmoduletypes.StoreKey]).(*iavl.Store)
		if !ok {
			os.Exit(1)
		}
		paymentIavl.EnableDiff()
	}
	app.initModules(ctx)
	// add eth query router
	ethRouter := app.BaseApp.EthQueryRouter()
	ethRouter.RegisterConstHandler()
	ethRouter.RegisterEthQueryBalanceHandler(app.BankKeeper, bankkeeper.EthQueryBalanceHandlerGen)

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper

	// Finally start the tpsCounter.
	app.tpsCounter = newTPSCounter(logger)
	go func() {
		// Unfortunately golangci-lint is so pedantic
		// so we have to ignore this error explicitly.
		_ = app.tpsCounter.start(context.Background())
	}()

	return app
}

func (app *Evmos) initModules(_ sdk.Context) {
	app.initCrossChain()
	app.initBridge()
	app.initStorage()
	app.initGov()
}

func (app *Evmos) initCrossChain() {
	app.CrossChainKeeper.SetSrcChainID(sdk.ChainID(app.appConfig.CrossChain.SrcChainId))
	app.CrossChainKeeper.SetDestBscChainID(sdk.ChainID(app.appConfig.CrossChain.DestBscChainId))
	app.CrossChainKeeper.SetDestPolygonChainID(sdk.ChainID(app.appConfig.CrossChain.DestPolygonChainId))
	app.CrossChainKeeper.SetDestScrollChainID(sdk.ChainID(app.appConfig.CrossChain.DestScrollChainId))
	app.CrossChainKeeper.SetDestLineaChainID(sdk.ChainID(app.appConfig.CrossChain.DestLineaChainId))
	app.CrossChainKeeper.SetDestMantleChainID(sdk.ChainID(app.appConfig.CrossChain.DestMantleChainId))
	app.CrossChainKeeper.SetDestArbitrumChainID(sdk.ChainID(app.appConfig.CrossChain.DestArbitrumChainId))
	app.CrossChainKeeper.SetDestOptimismChainID(sdk.ChainID(app.appConfig.CrossChain.DestOptimismChainId))
}

func (app *Evmos) initBridge() {
	bridgemodulekeeper.RegisterCrossApps(app.BridgeKeeper)
}

func (app *Evmos) initStorage() {
	storagemodulekeeper.RegisterCrossApps(app.StorageKeeper)
	storagemodulekeeper.InitPaymentCheck(app.StorageKeeper, app.appConfig.PaymentCheck.Enabled,
		app.appConfig.PaymentCheck.Interval)
}

func (app *Evmos) initGov() {
	err := app.GovKeeper.RegisterCrossChainSyncParamsApp()
	if err != nil {
		panic(err)
	}
}

// Name returns the name of the App
func (app *Evmos) Name() string { return app.BaseApp.Name() }

func (app *Evmos) setAnteHandler(txConfig client.TxConfig, maxGasWanted uint64) {
	options := ante.HandlerOptions{
		Cdc:                    app.appCodec,
		AccountKeeper:          app.AccountKeeper,
		BankKeeper:             app.BankKeeper,
		ExtensionOptionChecker: evmostypes.HasDynamicFeeExtensionOption,
		EvmKeeper:              app.EvmKeeper,
		FeegrantKeeper:         app.FeeGrantKeeper,
		GashubKeeper:           app.GashubKeeper,
		DistributionKeeper:     app.DistrKeeper,
		IBCKeeper:              app.IBCKeeper,
		FeeMarketKeeper:        app.FeeMarketKeeper,
		SignModeHandler:        txConfig.SignModeHandler(),
		SigGasConsumer:         ante.SigVerificationGasConsumer,
		MaxTxGasWanted:         maxGasWanted,
		TxFeeChecker:           ethante.NewDynamicFeeChecker(app.EvmKeeper),
	}

	if err := options.Validate(); err != nil {
		panic(err)
	}

	app.SetAnteHandler(ante.NewAnteHandler(options))
}

func (app *Evmos) setPostHandler() {
	postHandler, err := posthandler.NewPostHandler(
		posthandler.HandlerOptions{},
	)
	if err != nil {
		panic(err)
	}

	app.SetPostHandler(postHandler)
}

// BeginBlocker runs the Tendermint ABCI BeginBlock logic. It executes state changes at the beginning
// of the new block for every registered module. If there is a registered fork at the current height,
// BeginBlocker will schedule the upgrade plan and perform the state migration (if any).
func (app *Evmos) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	// Perform any scheduled forks before executing the modules logic
	app.ScheduleForkUpgrade(ctx)
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker updates every end block
func (app *Evmos) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	resp := app.mm.EndBlock(ctx, req)
	if app.IsIavlStore() {
		bankIavl, _ := app.CommitMultiStore().GetCommitStore(app.GetKey(banktypes.StoreKey)).(*iavl.Store)
		paymentIavl, _ := app.CommitMultiStore().GetCommitStore(app.GetKey(paymentmoduletypes.StoreKey)).(*iavl.Store)

		reconCtx, _ := ctx.CacheContext()
		reconCtx = reconCtx.WithGasMeter(sdk.NewInfiniteGasMeter())
		app.reconcile(reconCtx, bankIavl, paymentIavl)
	}
	return resp
}

// The DeliverTx method is intentionally decomposed to calculate the transactions per second.
func (app *Evmos) DeliverTx(req abci.RequestDeliverTx) (res abci.ResponseDeliverTx) {
	defer func() {
		// TODO: Record the count along with the code and or reason so as to display
		// in the transactions per second live dashboards.
		if res.IsErr() {
			app.tpsCounter.incrementFailure()
		} else {
			app.tpsCounter.incrementSuccess()
		}
	}()
	return app.BaseApp.DeliverTx(req)
}

// InitChainer updates at chain initialization
func (app *Evmos) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState simapp.GenesisState
	if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())

	// init cross chain channel permissions
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestBscChainId), bridgemoduletypes.TransferOutChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestBscChainId), bridgemoduletypes.TransferInChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestBscChainId), bridgemoduletypes.SyncParamsChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestBscChainId), storagemoduletypes.BucketChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestBscChainId), storagemoduletypes.ObjectChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestBscChainId), storagemoduletypes.GroupChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestBscChainId), storagemoduletypes.ZkmeSBTChannelId, sdk.ChannelAllow)

	// support polygon
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestPolygonChainId), bridgemoduletypes.TransferOutChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestPolygonChainId), bridgemoduletypes.TransferInChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestPolygonChainId), bridgemoduletypes.SyncParamsChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestPolygonChainId), storagemoduletypes.BucketChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestPolygonChainId), storagemoduletypes.ObjectChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestPolygonChainId), storagemoduletypes.GroupChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestPolygonChainId), storagemoduletypes.ZkmeSBTChannelId, sdk.ChannelAllow)

	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestScrollChainId), bridgemoduletypes.TransferOutChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestScrollChainId), bridgemoduletypes.TransferInChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestScrollChainId), bridgemoduletypes.SyncParamsChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestScrollChainId), storagemoduletypes.BucketChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestScrollChainId), storagemoduletypes.ObjectChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestScrollChainId), storagemoduletypes.GroupChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestScrollChainId), storagemoduletypes.ZkmeSBTChannelId, sdk.ChannelAllow)

	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestLineaChainId), bridgemoduletypes.TransferOutChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestLineaChainId), bridgemoduletypes.TransferInChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestLineaChainId), bridgemoduletypes.SyncParamsChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestLineaChainId), storagemoduletypes.BucketChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestLineaChainId), storagemoduletypes.ObjectChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestLineaChainId), storagemoduletypes.GroupChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestLineaChainId), storagemoduletypes.ZkmeSBTChannelId, sdk.ChannelAllow)

	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestMantleChainId), bridgemoduletypes.TransferOutChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestMantleChainId), bridgemoduletypes.TransferInChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestMantleChainId), bridgemoduletypes.SyncParamsChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestMantleChainId), storagemoduletypes.BucketChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestMantleChainId), storagemoduletypes.ObjectChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestMantleChainId), storagemoduletypes.GroupChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestMantleChainId), storagemoduletypes.ZkmeSBTChannelId, sdk.ChannelAllow)

	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestArbitrumChainId), bridgemoduletypes.TransferOutChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestArbitrumChainId), bridgemoduletypes.TransferInChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestArbitrumChainId), bridgemoduletypes.SyncParamsChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestArbitrumChainId), storagemoduletypes.BucketChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestArbitrumChainId), storagemoduletypes.ObjectChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestArbitrumChainId), storagemoduletypes.GroupChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestArbitrumChainId), storagemoduletypes.ZkmeSBTChannelId, sdk.ChannelAllow)

	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestOptimismChainId), bridgemoduletypes.TransferOutChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestOptimismChainId), bridgemoduletypes.TransferInChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestOptimismChainId), bridgemoduletypes.SyncParamsChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestOptimismChainId), storagemoduletypes.BucketChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestOptimismChainId), storagemoduletypes.ObjectChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestOptimismChainId), storagemoduletypes.GroupChannelID, sdk.ChannelAllow)
	app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestOptimismChainId), storagemoduletypes.ZkmeSBTChannelId, sdk.ChannelAllow)

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads state at a particular height
func (app *Evmos) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *Evmos) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)

	accs := make([]string, 0, len(maccPerms))
	for k := range maccPerms {
		accs = append(accs, k)
	}
	sort.Strings(accs)

	for _, acc := range accs {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlockedAddrs returns all the app's module account addresses that are not
// allowed to receive external tokens.
func (app *Evmos) BlockedAddrs() map[string]bool {
	blockedAddrs := make(map[string]bool)

	accs := make([]string, 0, len(maccPerms))
	for k := range maccPerms {
		accs = append(accs, k)
	}
	sort.Strings(accs)

	for _, acc := range accs {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return blockedAddrs
}

// LegacyAmino returns Evmos's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *Evmos) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns Evmos's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *Evmos) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Evmos's InterfaceRegistry
func (app *Evmos) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *Evmos) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *Evmos) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *Evmos) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *Evmos) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *Evmos) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx

	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register node gRPC service for grpc-gateway.
	node.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

func (app *Evmos) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *Evmos) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

// RegisterNodeService registers the node gRPC service on the provided
// application gRPC query router.
func (app *Evmos) RegisterNodeService(clientCtx client.Context) {
	node.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}

// IBC Go TestingApp functions

// GetBaseApp implements the TestingApp interface.
func (app *Evmos) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

// GetStakingKeeper implements the TestingApp interface.
func (app *Evmos) GetStakingKeeper() ibctestingtypes.StakingKeeper {
	return app.StakingKeeper
}

// GetStakingKeeperSDK implements the TestingApp interface.
func (app *Evmos) GetStakingKeeperSDK() stakingkeeper.Keeper {
	return *app.StakingKeeper
}

// GetIBCKeeper implements the TestingApp interface.
func (app *Evmos) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *Evmos) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetTxConfig implements the TestingApp interface.
func (app *Evmos) GetTxConfig() client.TxConfig {
	cfg := encoding.MakeConfig(ModuleBasics)
	return cfg.TxConfig
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(_ client.Context, rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}

	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(
	appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey,
) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	// SDK subspaces
	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable()) //nolint: staticcheck
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibcexported.ModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	// ethermint subspaces
	paramsKeeper.Subspace(evmtypes.ModuleName).WithKeyTable(evmtypes.ParamKeyTable()) //nolint: staticcheck
	paramsKeeper.Subspace(feemarkettypes.ModuleName).WithKeyTable(feemarkettypes.ParamKeyTable())
	// evmos subspaces
	paramsKeeper.Subspace(erc20types.ModuleName)
	return paramsKeeper
}

// EvmPrecompiled  set evm precompiled contracts
func (app *Evmos) EvmPrecompiled() {
	precompiled := evmkeeper.BerlinPrecompiled()

	// bank precompile
	precompiled[precompilesbank.GetAddress()] = func(ctx sdk.Context) vm.PrecompiledContract {
		return precompilesbank.NewPrecompiledContract(ctx, app.BankKeeper)
	}

	// authz precompile
	precompiled[precompilesauthz.GetAddress()] = func(ctx sdk.Context) vm.PrecompiledContract {
		return precompilesauthz.NewPrecompiledContract(ctx, app.AuthzKeeper)
	}

	// gov precompile
	precompiled[precompilesgov.GetAddress()] = func(ctx sdk.Context) vm.PrecompiledContract {
		return precompilesgov.NewPrecompiledContract(ctx, app.GovKeeper, app.AccountKeeper)
	}

	// storage precompile
	precompiled[precompilesstorage.GetAddress()] = func(ctx sdk.Context) vm.PrecompiledContract {
		return precompilesstorage.NewPrecompiledContract(ctx, app.StorageKeeper)
	}

	// virtualgroup precompile
	precompiled[precompilesvirtualgroup.GetAddress()] = func(ctx sdk.Context) vm.PrecompiledContract {
		return precompilesvirtualgroup.NewPrecompiledContract(ctx, app.VirtualgroupKeeper)
	}

	// set precompiled contracts
	app.EvmKeeper.WithPrecompiled(precompiled)
}

func (app *Evmos) setupUpgradeHandlers() {
	// When a planned update height is reached, the old binary will panic
	// writing on disk the height and name of the update that triggered it
	// This will read that value, and execute the preparations for the upgrade.
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Errorf("failed to read upgrade info from disk: %w", err))
	}

	var storeUpgrades *storetypes.StoreUpgrades

	if storeUpgrades != nil {
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, storeUpgrades))
	}
}

func MakeEncodingConfig() simappparams.EncodingConfig {
	return encoding.MakeConfig(ModuleBasics)
}
