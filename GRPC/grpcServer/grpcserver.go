package grpcServer

import (
	"context"
	"fmt"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPC"
	"github.com/Baraulia/COURIER_SERVICE/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type GRPCServer struct {
	service *service.Service
	courierProto.UnimplementedCourierServerServer
}

func NewGRPCServer(service *service.Service) {
	s := grpc.NewServer()
	str := &GRPCServer{service: service}
	courierProto.RegisterCourierServerServer(s, str)
	lis, err := net.Listen("tcp", ":8091")
	if err != nil {
		log.Fatalf("NewGRPCServer, Listen:%s", err)
	}
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("NewGRPCServer, Serve:%s", err)
	}

}

func (g *GRPCServer) CreateOrder(ctx context.Context, order *courierProto.OrderCourierServer) (*emptypb.Empty, error) {
	return nil, g.service.OrderApp.CreateOrder(order)
}

func (g *GRPCServer) GetDeliveryService(ctx context.Context, in *emptypb.Empty) (*courierProto.ServiceResponse, error) {
	res, err := g.service.OrderApp.GetServices()
	if err != nil {
		log.Fatalf("GetService:%s", err)
		return nil, fmt.Errorf("GetService:%w", err)
	}
	return res, nil
}