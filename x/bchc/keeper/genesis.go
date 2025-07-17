package keeper

import (
	"context"

	"bchc/x/bchc/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.PatientList {
		if err := k.Patient.Set(ctx, elem.Id, elem); err != nil {
			return err
		}
	}

	if err := k.PatientSeq.Set(ctx, genState.PatientCount); err != nil {
		return err
	}
	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	err = k.Patient.Walk(ctx, nil, func(key uint64, elem types.Patient) (bool, error) {
		genesis.PatientList = append(genesis.PatientList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.PatientCount, err = k.PatientSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}

	return genesis, nil
}
