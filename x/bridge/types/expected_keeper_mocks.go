// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bnb-chain/greenfield/x/bridge/types (interfaces: BankKeeper,StakingKeeper,CrossChainKeeper)

// Package types is a generated GoMock package.
package types

import (
	big "math/big"
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types.Context, addr types.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SendCoinsFromModuleToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToModule(ctx types.Context, senderModule, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToModule", ctx, senderModule, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToModule indicates an expected call of SendCoinsFromModuleToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToModule), ctx, senderModule, recipientModule, amt)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// BondDenom mocks base method.
func (m *MockStakingKeeper) BondDenom(ctx types.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BondDenom", ctx)
	ret0, _ := ret[0].(string)
	return ret0
}

// BondDenom indicates an expected call of BondDenom.
func (mr *MockStakingKeeperMockRecorder) BondDenom(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BondDenom", reflect.TypeOf((*MockStakingKeeper)(nil).BondDenom), ctx)
}

// MockCrossChainKeeper is a mock of CrossChainKeeper interface.
type MockCrossChainKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockCrossChainKeeperMockRecorder
}

// MockCrossChainKeeperMockRecorder is the mock recorder for MockCrossChainKeeper.
type MockCrossChainKeeperMockRecorder struct {
	mock *MockCrossChainKeeper
}

// NewMockCrossChainKeeper creates a new mock instance.
func NewMockCrossChainKeeper(ctrl *gomock.Controller) *MockCrossChainKeeper {
	mock := &MockCrossChainKeeper{ctrl: ctrl}
	mock.recorder = &MockCrossChainKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrossChainKeeper) EXPECT() *MockCrossChainKeeperMockRecorder {
	return m.recorder
}

// CreateRawIBCPackageWithFee mocks base method.
func (m *MockCrossChainKeeper) CreateRawIBCPackageWithFee(ctx types.Context, chainID types.ChainID, channelID types.ChannelID, packageType types.CrossChainPackageType, packageLoad []byte, relayerFee, ackRelayerFee *big.Int) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRawIBCPackageWithFee", ctx, chainID, channelID, packageType, packageLoad, relayerFee, ackRelayerFee)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRawIBCPackageWithFee indicates an expected call of CreateRawIBCPackageWithFee.
func (mr *MockCrossChainKeeperMockRecorder) CreateRawIBCPackageWithFee(ctx, chainID, channelID, packageType, packageLoad, relayerFee, ackRelayerFee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRawIBCPackageWithFee", reflect.TypeOf((*MockCrossChainKeeper)(nil).CreateRawIBCPackageWithFee), ctx, chainID, channelID, packageType, packageLoad, relayerFee, ackRelayerFee)
}

// GetDestBscChainID mocks base method.
func (m *MockCrossChainKeeper) GetDestBscChainID() types.ChainID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDestBscChainID")
	ret0, _ := ret[0].(types.ChainID)
	return ret0
}

// GetDestBscChainID indicates an expected call of GetDestBscChainID.
func (mr *MockCrossChainKeeperMockRecorder) GetDestBscChainID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDestBscChainID", reflect.TypeOf((*MockCrossChainKeeper)(nil).GetDestBscChainID))
}

// RegisterChannel mocks base method.
func (m *MockCrossChainKeeper) RegisterChannel(name string, id types.ChannelID, app types.CrossChainApplication) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterChannel", name, id, app)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterChannel indicates an expected call of RegisterChannel.
func (mr *MockCrossChainKeeperMockRecorder) RegisterChannel(name, id, app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterChannel", reflect.TypeOf((*MockCrossChainKeeper)(nil).RegisterChannel), name, id, app)
}