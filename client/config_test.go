package client

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func TestInitConfigNonNotExistError(t *testing.T) {
	tempDir := t.TempDir()
	subDir := filepath.Join(tempDir, "nonPerms")
	if err := os.Mkdir(subDir, 0o600); err != nil {
		t.Fatalf("Failed to create sub directory: %v", err)
	}
	cmd := &cobra.Command{}
	cmd.PersistentFlags().String(flags.FlagHome, "", "")
	if err := cmd.PersistentFlags().Set(flags.FlagHome, subDir); err != nil {
		t.Fatalf("Could not set home flag [%T] %v", err, err)
	}
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")
	cmd.PersistentFlags().StringP(cli.EncodingFlag, "e", "hex", "Binary encoding (hex|b64|btc)")
	cmd.PersistentFlags().StringP(cli.OutputFlag, "o", "text", "Output format (text|json)")

	if err := InitConfig(cmd); !os.IsPermission(err) {
		t.Fatalf("Failed to catch permissions error, got: [%T] %v", err, err)
	}
}
