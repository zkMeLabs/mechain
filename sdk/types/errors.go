package types

import (
	"errors"
)

var (
	ErrKeyManagerNotInit     = errors.New("key manager is not initialized yet ")
	ErrChainIDNotSet         = errors.New("chainID is not set yet ")
	ErrSimulatedGasPrice     = errors.New("simulated gas price is 0 ")
	ErrFeeAmountNotValid     = errors.New("fee Amount coin should only be azkme")
	ErrGasInfoNotProvided    = errors.New("gas limit and(or) Fee Amount missing in txOpt")
	ErrRpcAddressNotProvided = errors.New("rpc address is not provided")
)
