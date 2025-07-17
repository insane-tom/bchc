package keeper_test

import (
	"context"
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bchc/x/bchc/keeper"
	"bchc/x/bchc/types"
)

func createNPatient(keeper keeper.Keeper, ctx context.Context, n int) []types.Patient {
	items := make([]types.Patient, n)
	for i := range items {
		iu := uint64(i)
		items[i].Id = iu
		items[i].Name = strconv.Itoa(i)
		items[i].Hospital = strconv.Itoa(i)
		items[i].Disease = strconv.Itoa(i)
		_ = keeper.Patient.Set(ctx, iu, items[i])
		_ = keeper.PatientSeq.Set(ctx, iu)
	}
	return items
}

func TestPatientQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNPatient(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetPatientRequest
		response *types.QueryGetPatientResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetPatientRequest{Id: msgs[0].Id},
			response: &types.QueryGetPatientResponse{Patient: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetPatientRequest{Id: msgs[1].Id},
			response: &types.QueryGetPatientResponse{Patient: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetPatientRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetPatient(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestPatientQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNPatient(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPatientRequest {
		return &types.QueryAllPatientRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListPatient(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Patient), step)
			require.Subset(t, msgs, resp.Patient)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListPatient(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Patient), step)
			require.Subset(t, msgs, resp.Patient)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListPatient(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Patient)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListPatient(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
