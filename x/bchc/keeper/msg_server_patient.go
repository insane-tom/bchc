package keeper

import (
	"context"
	"errors"
	"fmt"

	"bchc/x/bchc/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePatient(ctx context.Context, msg *types.MsgCreatePatient) (*types.MsgCreatePatientResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	nextId, err := k.PatientSeq.Next(ctx)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get next id")
	}

	var patient = types.Patient{
		Id:       nextId,
		Creator:  msg.Creator,
		Name:     msg.Name,
		Hospital: msg.Hospital,
		Disease:  msg.Disease,
	}

	if err = k.Patient.Set(
		ctx,
		nextId,
		patient,
	); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set patient")
	}

	return &types.MsgCreatePatientResponse{
		Id: nextId,
	}, nil
}

func (k msgServer) UpdatePatient(ctx context.Context, msg *types.MsgUpdatePatient) (*types.MsgUpdatePatientResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	var patient = types.Patient{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Name:     msg.Name,
		Hospital: msg.Hospital,
		Disease:  msg.Disease,
	}

	// Checks that the element exists
	val, err := k.Patient.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get patient")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Patient.Set(ctx, msg.Id, patient); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update patient")
	}

	return &types.MsgUpdatePatientResponse{}, nil
}

func (k msgServer) DeletePatient(ctx context.Context, msg *types.MsgDeletePatient) (*types.MsgDeletePatientResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Checks that the element exists
	val, err := k.Patient.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get patient")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Patient.Remove(ctx, msg.Id); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to delete patient")
	}

	return &types.MsgDeletePatientResponse{}, nil
}
