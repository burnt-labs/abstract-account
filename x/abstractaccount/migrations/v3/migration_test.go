package v3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	simapptesting "github.com/burnt-labs/abstract-account/simapp/testing"
	"github.com/burnt-labs/abstract-account/x/abstractaccount/types"
)

func TestMigrateStoreDisablesRegistrationUntilChainConfiguresBootstrap(t *testing.T) {
	app := simapptesting.MakeSimpleMockApp()
	ctx := app.NewContext(false)

	params, err := app.AbstractAccountKeeper.GetParams(ctx)
	require.NoError(t, err)
	params.BootstrapCodeID = 11
	require.NoError(t, app.AbstractAccountKeeper.SetParams(ctx, params))

	require.NoError(t, app.AbstractAccountKeeper.Migrator().Migrate2to3(ctx))

	migrated, err := app.AbstractAccountKeeper.GetParams(ctx)
	require.NoError(t, err)
	require.Zero(t, migrated.BootstrapCodeID)
	require.False(t, migrated.RegistrationEnabled)
	require.Equal(t, uint64(types.DefaultMaxGas), migrated.MaxGasBefore)
}
