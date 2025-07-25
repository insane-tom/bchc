package types_test

import (
	"testing"

	"bchc/x/bchc/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid genesis state",
			genState: &types.GenesisState{PatientList: []types.Patient{{Id: 0}, {Id: 1}}, PatientCount: 2}, valid: true,
		}, {
			desc: "duplicated patient",
			genState: &types.GenesisState{
				PatientList: []types.Patient{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		}, {
			desc: "invalid patient count",
			genState: &types.GenesisState{
				PatientList: []types.Patient{
					{
						Id: 1,
					},
				},
				PatientCount: 0,
			},
			valid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
