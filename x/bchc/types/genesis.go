package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:      DefaultParams(),
		PatientList: []Patient{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	patientIdMap := make(map[uint64]bool)
	patientCount := gs.GetPatientCount()
	for _, elem := range gs.PatientList {
		if _, ok := patientIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for patient")
		}
		if elem.Id >= patientCount {
			return fmt.Errorf("patient id should be lower or equal than the last id")
		}
		patientIdMap[elem.Id] = true
	}

	return gs.Params.Validate()
}
