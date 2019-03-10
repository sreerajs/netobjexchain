package assetservice

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the assetservice Querier
const (
	QueryResolve = "resolve"
	QueryWhois   = "whois"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, keeper)
		case QueryWhois:
			return queryWhois(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown assetservice query endpoint")
		}
	}
}

// nolint: unparam
func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	asset := path[0]

	value := keeper.ResolveAsset(ctx, asset)

	if value == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve asset")
	}

	return []byte(value), nil
}

// Whois represents a asset -> value lookup
type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

// nolint: unparam
func queryWhois(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	asset := path[0]

	whois := Whois{}

	whois.Value = keeper.ResolveAsset(ctx, asset)
	whois.Owner = keeper.GetOwner(ctx, asset)
	whois.Price = keeper.GetPrice(ctx, asset)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, whois)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
