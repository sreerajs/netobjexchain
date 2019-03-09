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
func NewKeeper(coinKeeper bank.Keeper, assetStoreKey sdk.StoreKey, ownersStoreKey sdk.StoreKey, priceStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper:     coinKeeper,
		assetsStoreKey:  assetsStoreKey,
		ownersStoreKey: ownersStoreKey,
		pricesStoreKey: priceStoreKey,
		cdc:            cdc,
	}
}

// ResolveAsset - returns the string that the asset resolves to
func (k Keeper) ResolveAsset(ctx sdk.Context, name string) string {
	store := ctx.KVStore(k.assetsStoreKey)
	bz := store.Get([]byte(name))
	return string(bz)
}

// SetAsset - sets the value string that a asset resolves to
func (k Keeper) SetAsset(ctx sdk.Context, name string, value string) {
	store := ctx.KVStore(k.assetsStoreKey)
	store.Set([]byte(name), []byte(value))
}

// HasOwner - returns whether or not the asset already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.ownersStoreKey)
	bz := store.Get([]byte(name))
	return bz != nil
}

// GetOwner - get the current owner of a asset
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	store := ctx.KVStore(k.ownersStoreKey)
	bz := store.Get([]byte(name))
	return bz
}

// SetOwner - sets the current owner of a asset
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.ownersStoreKey)
	store.Set([]byte(name), owner)
}

// GetPrice - gets the current price of a asset.  If price doesn't exist yet, set to 1mycoin.
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	if !k.HasOwner(ctx, name) {
		return sdk.Coins{sdk.NewInt64Coin("mycoin", 1)}
	}
	store := ctx.KVStore(k.pricesStoreKey)
	bz := store.Get([]byte(name))
	var price sdk.Coins
	// k.cdc.MustUnmarshalBinary(bz, &price)
	k.cdc.MustUnmarshalBinaryBare(bz, &price)
	return price
}

// SetPrice - sets the current price of a asset
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	store := ctx.KVStore(k.pricesStoreKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(price))
}
