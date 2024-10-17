package client

import (
	"google.golang.org/grpc"

	"github.com/evmos/evmos/v12/sdk/keys"
)

// MechainClientOption configures how we set up the mechain client.
type MechainClientOption interface {
	Apply(*MechainClient)
}

// MechainClientOptionFunc defines an applied function for setting the mechain client.
type MechainClientOptionFunc func(*MechainClient)

// Apply set up the option field to the client instance.
func (f MechainClientOptionFunc) Apply(client *MechainClient) {
	f(client)
}

// WithKeyManager returns a MechainClientOption which configures a client key manager option.
func WithKeyManager(km keys.KeyManager) MechainClientOption {
	return MechainClientOptionFunc(func(client *MechainClient) {
		client.keyManager = km
	})
}

// WithGrpcConnectionAndDialOption returns a MechainClientOption which configures a grpc client connection with grpc dail options.
func WithGrpcConnectionAndDialOption(grpcAddr string, opts ...grpc.DialOption) MechainClientOption {
	return MechainClientOptionFunc(func(client *MechainClient) {
		client.grpcConn = grpcConn(grpcAddr, opts...)
	})
}

// WithWebSocketClient returns a MechainClientOption which specify that connection is a websocket connection
func WithWebSocketClient() MechainClientOption {
	return MechainClientOptionFunc(func(client *MechainClient) {
		client.useWebSocket = true
	})
}

// grpcConn is used to establish a connection with a given address and dial options.
func grpcConn(addr string, opts ...grpc.DialOption) *grpc.ClientConn {
	conn, err := grpc.Dial(
		addr,
		opts...,
	)
	if err != nil {
		panic(err)
	}
	return conn
}
