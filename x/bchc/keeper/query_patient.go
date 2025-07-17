package keeper

import (
	"context"
	"errors"

	"bchc/x/bchc/types"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListPatient(ctx context.Context, req *types.QueryAllPatientRequest) (*types.QueryAllPatientResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	patients, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Patient,
		req.Pagination,
		func(_ uint64, value types.Patient) (types.Patient, error) {
			return value, nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPatientResponse{Patient: patients, Pagination: pageRes}, nil
}

func (q queryServer) GetPatient(ctx context.Context, req *types.QueryGetPatientRequest) (*types.QueryGetPatientResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	patient, err := q.k.Patient.Get(ctx, req.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetPatientResponse{Patient: patient}, nil
}
