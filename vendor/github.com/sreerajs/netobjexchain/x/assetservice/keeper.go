package assetservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	assetsStoreKey  sdk.StoreKey // Unexposed key to access asset store from sdk.Context
	ownersStoreKey sdk.StoreKey // Unexposed key to access owners store from sdk.Context
	pricesStoreKey sdk.StoreKey // Unexposed key to access prices store from sdk.Context
	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the assetservice Keeper
func NewKeeper(coinKeeper bank.Keeper, assetsStoreKey sdk.StoreKey, ownersStoreKey sdk.StoreKey, priceStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper:     coinKeeper,
		assetsStoreKey:  assetsStoreKey,
		ownersStoreKey: ownersStoreKey,
		pricesStoreKey: priceStoreKey,
		cdc:            cdc,
	}
}

// ResolveAsset - returns the string that the asset resolves to
func (k Keeper) ResolveAsset(ctx sdk.Context, asset string) string {
	store := ctx.KVStore(k.assetsStoreKey)
	bz := store.Get([]byte(asset))
	return string(bz)
}

// SetAsset - sets the value string that a asset resolves to
func (k Keeper) SetAsset(ctx sdk.Context, asset string, value string) {
	store := ctx.KVStore(k.assetsStoreKey)
	store.Set([]byte(asset), []byte(value))
}

// HasOwner - returns whether or not the asset already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, asset string) bool {
	store := ctx.KVStore(k.ownersStoreKey)
	bz := store.Get([]byte(asset))
	return bz != nil
}

// GetOwner - get the current owner of a asset
func (k Keeper) GetOwner(ctx sdk.Context, asset string) sdk.AccAddress {
	store := ctx.KVStore(k.ownersStoreKey)
	bz := store.Get([]byte(asset))
	return bz
}

// SetOwner - sets the current owner of a asset
func (k Keeper) SetOwner(ctx sdk.Context, asset string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.ownersStoreKey)
	store.Set([]byte(asset), owner)
}

// GetPrice - gets the current price of a asset.  If price doesn't exist yet, set to 1mycoin.
func (k Keeper) GetPrice(ctx sdk.Context, asset string) sdk.Coins {
	if !k.HasOwner(ctx, asset) {
		return sdk.Coins{sdk.NewInt64Coin("mycoin", 1)}
	}
	store := ctx.KVStore(k.pricesStoreKey)
	bz := store.Get([]byte(asset))
	var price sdk.Coins
	// k.cdc.MustUnmarshalBinary(bz, &price)
	k.cdc.MustUnmarshalBinaryBare(bz, &price)
	return price
}

// SetPrice - sets the current price of a asset
func (k Keeper) SetPrice(ctx sdk.Context, asset string, price sdk.Coins) {
	store := ctx.KVStore(k.pricesStoreKey)
	store.Set([]byte(asset), k.cdc.MustMarshalBinaryBare(price))
}
