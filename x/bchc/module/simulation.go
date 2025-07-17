package bchc

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"bchc/testutil/sample"
	bchcsimulation "bchc/x/bchc/simulation"
	"bchc/x/bchc/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bchcGenesis := types.GenesisState{
		Params:      types.DefaultParams(),
		PatientList: []types.Patient{{Id: 0, Creator: sample.AccAddress()}, {Id: 1, Creator: sample.AccAddress()}}, PatientCount: 2,
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bchcGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreatePatient          = "op_weight_msg_bchc"
		defaultWeightMsgCreatePatient int = 100
	)

	var weightMsgCreatePatient int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePatient, &weightMsgCreatePatient, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePatient = defaultWeightMsgCreatePatient
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePatient,
		bchcsimulation.SimulateMsgCreatePatient(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdatePatient          = "op_weight_msg_bchc"
		defaultWeightMsgUpdatePatient int = 100
	)

	var weightMsgUpdatePatient int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePatient, &weightMsgUpdatePatient, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePatient = defaultWeightMsgUpdatePatient
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePatient,
		bchcsimulation.SimulateMsgUpdatePatient(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeletePatient          = "op_weight_msg_bchc"
		defaultWeightMsgDeletePatient int = 100
	)

	var weightMsgDeletePatient int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePatient, &weightMsgDeletePatient, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePatient = defaultWeightMsgDeletePatient
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePatient,
		bchcsimulation.SimulateMsgDeletePatient(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
