package types

func NewMsgCreatePatient(creator string, name string, hospital string, disease string) *MsgCreatePatient {
	return &MsgCreatePatient{
		Creator:  creator,
		Name:     name,
		Hospital: hospital,
		Disease:  disease,
	}
}

func NewMsgUpdatePatient(creator string, id uint64, name string, hospital string, disease string) *MsgUpdatePatient {
	return &MsgUpdatePatient{
		Id:       id,
		Creator:  creator,
		Name:     name,
		Hospital: hospital,
		Disease:  disease,
	}
}

func NewMsgDeletePatient(creator string, id uint64) *MsgDeletePatient {
	return &MsgDeletePatient{
		Id:      id,
		Creator: creator,
	}
}
