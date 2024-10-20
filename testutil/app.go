package testutil

import (
	"encoding/hex"
	"encoding/json"
	"io"

	"cosmossdk.io/simapp/params"
	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/keys/eth/ethsecp256k1"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/evmos/evmos/v12/app"
	"github.com/evmos/evmos/v12/encoding"
	"github.com/evmos/evmos/v12/sdk/client/test"
	servercfg "github.com/evmos/evmos/v12/server/config"
	evmostypes "github.com/evmos/evmos/v12/types"
)

func NewTestApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	chainID string,
	options ...func(baseApp *baseapp.BaseApp),
) (*app.Evmos, params.EncodingConfig, error) {
	// create public key
	privVal := mock.NewPV()
	pubKey, _ := privVal.GetPubKey()

	// create validator set with single validator
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	// generate genesis account
	bz, _ := hex.DecodeString(test.TestPublicKey)
	faucetPubKey := &ethsecp256k1.PubKey{Key: bz}

	acc := authtypes.NewBaseAccount(faucetPubKey.Address().Bytes(), faucetPubKey, 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(test.TestTokenName, sdk.NewInt(100000000000000))),
	}

	encCfg := encoding.MakeConfig(app.ModuleBasics)
	options = append(options, baseapp.SetChainID(chainID))
	nApp := app.NewEvmos(
		logger,
		db,
		traceStore,
		loadLatest,
		app.DefaultNodeHome,
		0,
		encCfg,
		servercfg.NewDefaultAppConfig(evmostypes.AttoEvmos),
		simtestutil.EmptyAppOptions{},
		options...,
	)

	genesisState := app.NewDefaultGenesisState()
	genesisState, _ = simtestutil.GenesisStateWithValSet(nApp.AppCodec(), genesisState, valSet, []authtypes.GenesisAccount{acc}, balance)

	stateBytes, _ := json.MarshalIndent(genesisState, "", "  ")

	// Initialize the chain
	nApp.InitChain(
		abci.RequestInitChain{
			ChainId:       chainID,
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	nApp.Commit()

	return nApp, encCfg, nil
}
