package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

// GetCmdResolveAsset queries information about a asset
func GetCmdResolveAsset(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "resolve [asset]",
		Short: "resolve asset",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", queryRoute, asset), nil)
			if err != nil {
				fmt.Printf("could not resolve asset - %s \n", string(asset))
				return nil
			}

			return cliCtx.PrintOutput(resolveRes{string(res)})
		},
	}
}

type resolveRes struct {
	Value string `json:"value"`
}

func (r resolveRes) String() string {
	return r.Value
}

// GetCmdWhois queries information about a domain
func GetCmdWhois(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "whois [asset]",
		Short: "Query whois info of asset",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", queryRoute, asset), nil)
			if err != nil {
				fmt.Printf("could not resolve whois - %s \n", string(asset))
				return nil
			}

			var out whoIsRes
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

type whoIsRes struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

func (w whoIsRes) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}
