package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	assetservicecmd "github.com/sreerajs/netobjexchain/x/assetservice/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group assetservice queries under a subcommand
	assetsvcQueryCmd := &cobra.Command{
		Use:   "assetservice",
		Short: "Querying commands for the assetservice module",
	}

	assetsvcQueryCmd.AddCommand(client.GetCommands(
		assetservicecmd.GetCmdResolveAsset(mc.storeKey, mc.cdc),
		assetservicecmd.GetCmdWhois(mc.storeKey, mc.cdc),
	)...)

	return assetsvcQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	assetsvcTxCmd := &cobra.Command{
		Use:   "assetservice",
		Short: "Assetservice transactions subcommands",
	}

	assetsvcTxCmd.AddCommand(client.PostCommands(
		assetservicecmd.GetCmdBuyAsset(mc.cdc),
		assetservicecmd.GetCmdSetAsset(mc.cdc),
	)...)

	return assetsvcTxCmd
}
