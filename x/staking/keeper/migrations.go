package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v043 "github.com/cosmos/cosmos-sdk/x/staking/migrations/v043"
	v045 "github.com/cosmos/cosmos-sdk/x/staking/migrations/v045"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v043.MigrateStore(ctx, m.keeper.storeKey)
}

// Migrate2to3 migrates x/staking state from consensus version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v045.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
