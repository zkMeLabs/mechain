package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/evmos/evmos/v12/x/virtualgroup/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdSettle())

	return cmd
}

func CmdSettle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settle [gvg family id] [gvg ids]",
		Short: "Broadcast message settle",
		Long: `Settle will do the settlement of a GVG family or several GVGs (by specifying comma separated ids). 
If zero is provided for GVG family, then the provided GVGs will be settled.
If none zero is provided for GVG family, then the provided GVG family will be settled and the provided GVGs will be ignored.`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			gvgFamilyID, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil || gvgFamilyID < 0 {
				return fmt.Errorf("invalid GVG family id %s", args[0])
			}
			gvgIDs := make([]uint32, 0)
			splits := strings.Split(args[1], ",")
			for _, split := range splits {
				gvgID, err := strconv.ParseInt(split, 10, 32)
				if err != nil || gvgID < 0 {
					return fmt.Errorf("invalid GVG id %s", args[1])
				}
				gvgIDs = append(gvgIDs, uint32(gvgID))
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSettle(
				clientCtx.GetFromAddress(),
				uint32(gvgFamilyID),
				gvgIDs,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
