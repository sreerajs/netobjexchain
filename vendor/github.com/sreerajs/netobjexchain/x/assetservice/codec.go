package assetservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetAsset{}, "assetservice/SetAsset", nil)
	cdc.RegisterConcrete(MsgBuyAsset{}, "assetservice/BuyAsset", nil)
}
