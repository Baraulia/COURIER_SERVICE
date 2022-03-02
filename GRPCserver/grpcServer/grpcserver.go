package grpcServer

import (
	"context"
	"fmt"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPCserver"
	"github.com/Baraulia/COURIER_SERVICE/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
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
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logrus.Fatalf("NewGRPCServer, Listen:%s", err)
	}
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		logrus.Fatalf("NewGRPCServer, Serve:%s", err)
	}

}

func (g *GRPCServer) CreateOrder(ctx context.Context, order *courierProto.OrderCourierServer) (*emptypb.Empty, error) {
	return nil, g.service.OrderApp.CreateOrder(order)
}

func (g *GRPCServer) GetDeliveryService(ctx context.Context, in *emptypb.Empty) (*courierProto.ServiceResponse, error) {
	res, err := g.service.OrderApp.GetServices()
	if err != nil {
		logrus.Errorf("GetService:%s", err)
		return nil, fmt.Errorf("GetService:%w", err)
	}
	return res, nil
}
