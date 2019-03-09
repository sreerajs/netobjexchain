package assetservice

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgSetAsset defines a SetAsset message
type MsgSetAsset struct {
	Name  string
	Value string
	Owner sdk.AccAddress
}

// NewMsgSetAsset is a constructor function for MsgSetAsset
func NewMsgSetAsset(name string, value string, owner sdk.AccAddress) MsgSetAsset {
	return MsgSetAsset{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

// Route should return the asset of the module
func (msg MsgSetAsset) Route() string { return "assetservice" }

// Type should return the action
func (msg MsgSetAsset) Type() string { return "set_asset" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetAsset) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Asset) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Asset and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetAsset) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgSetAsset) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgBuyAsset defines the BuyAsset message
type MsgBuyAsset struct {
	Name  string
	Bid   sdk.Coins
	Buyer sdk.AccAddress
}

// NewMsgBuyAsset is the constructor function for MsgBuyAsset
func NewMsgBuyAsset(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyAsset {
	return MsgBuyAsset{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}

// Route should return the asset of the module
func (msg MsgBuyAsset) Route() string { return "assetservice" }

// Type should return the action
func (msg MsgBuyAsset) Type() string { return "buy_asset" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyAsset) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.Asset) == 0 {
		return sdk.ErrUnknownRequest("Asset cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBuyAsset) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgBuyAsset) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}
