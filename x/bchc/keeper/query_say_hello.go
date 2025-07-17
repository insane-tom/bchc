package keeper

import (
    "context"
    "fmt"

    "bchc/x/bchc/types"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (q queryServer) SayHello(ctx context.Context, req *types.QuerySayHelloRequest) (*types.QuerySayHelloResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    // TODO: Process the query

    // Custom Response
    return &types.QuerySayHelloResponse{Name: fmt.Sprintf("Hello %s!", req.Name)}, nil
}