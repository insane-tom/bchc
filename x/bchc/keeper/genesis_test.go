package keeper_test

import (
	"testing"

	"bchc/x/bchc/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:       types.DefaultParams(),
		PatientList:  []types.Patient{{Id: 0}, {Id: 1}},
		PatientCount: 2,
	}
	f := initFixture(t)
	err := f.keeper.InitGenesis(f.ctx, genesisState)
	require.NoError(t, err)
	got, err := f.keeper.ExportGenesis(f.ctx)
	require.NoError(t, err)
	require.NotNil(t, got)

	require.EqualExportedValues(t, genesisState.Params, got.Params)
	require.EqualExportedValues(t, genesisState.PatientList, got.PatientList)
	require.Equal(t, genesisState.PatientCount, got.PatientCount)

}
