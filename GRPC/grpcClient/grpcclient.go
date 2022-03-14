package grpcClient

import (
	"context"
	"fmt"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type GRPCClient struct {
	cli courierProto.AuthClient
}

func NewGRPCClient(host string) *GRPCClient {
	Target := fmt.Sprintf("%s:8090", host)
	conn, err := grpc.Dial(Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("NewGRPCClient, Dial:%s", err)
	}
	cli := courierProto.NewAuthClient(conn)
	return &GRPCClient{cli: cli}
}

func (c *GRPCClient) GetUserWithRights(ctx context.Context, in *courierProto.AccessToken, opts ...grpc.CallOption) (*courierProto.UserRole, error) {
	return c.cli.GetUserWithRights(ctx, in)
}

func (c *GRPCClient) BindUserAndRole(ctx context.Context, in *courierProto.User, opts ...grpc.CallOption) (*courierProto.ResultBinding, error) {
	return nil, nil
}

func (c *GRPCClient) TokenGenerationByRefresh(ctx context.Context, in *courierProto.RefreshToken, opts ...grpc.CallOption) (*courierProto.GeneratedTokens, error) {
	return nil, nil
}

func (c *GRPCClient) TokenGenerationByUserId(ctx context.Context, in *courierProto.User, opts ...grpc.CallOption) (*courierProto.GeneratedTokens, error) {
	return nil, nil
}

func (c *GRPCClient) GetAllRoles(ctx context.Context, in *courierProto.Request, opts ...grpc.CallOption) (*courierProto.Roles, error) {
	return nil, nil
}