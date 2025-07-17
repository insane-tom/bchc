package bchc

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"bchc/x/bchc/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "SayHello",
					Use:            "say-hello [name]",
					Short:          "Query say-hello",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},

				{
					RpcMethod: "ListPatient",
					Use:       "list-patient",
					Short:     "List all patient",
				},
				{
					RpcMethod:      "GetPatient",
					Use:            "get-patient [id]",
					Short:          "Gets a patient by id",
					Alias:          []string{"show-patient"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePatient",
					Use:            "create-patient [name] [hospital] [disease]",
					Short:          "Create patient",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "hospital"}, {ProtoField: "disease"}},
				},
				{
					RpcMethod:      "UpdatePatient",
					Use:            "update-patient [id] [name] [hospital] [disease]",
					Short:          "Update patient",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "hospital"}, {ProtoField: "disease"}},
				},
				{
					RpcMethod:      "DeletePatient",
					Use:            "delete-patient [id]",
					Short:          "Delete patient",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
