package app

import (
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

type AppConfig struct {
	serverconfig.Config
}

var CustomAppTemplate = serverconfig.DefaultConfigTemplate

func NewDefaultAppConfig() *AppConfig {
	srvCfg := serverconfig.DefaultConfig()
	// The SDK's default minimum gas price is set to "" (empty value) inside
	// app.toml. If left empty by validators, the node will halt on startup.
	// However, the chain developer can set a default app.toml value for their
	// validators here.
	//
	// In summary:
	// - if you leave srvCfg.MinGasPrices = "", all validators MUST tweak their
	//   own app.toml config,
	// - if you set srvCfg.MinGasPrices non-empty, validators CAN tweak their
	//   own app.toml to override, or use this default value.
	//
	// In simapp, we set the min gas prices to 0.
	srvCfg.MinGasPrices = "5000000000azkme" // 5gei

	return &AppConfig{
		Config: *srvCfg,
	}
}
