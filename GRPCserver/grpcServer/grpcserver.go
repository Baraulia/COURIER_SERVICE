package grpcServer

import (
	"context"
	"fmt"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPCserver"
	"github.com/Baraulia/COURIER_SERVICE/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

func (g *GRPCServer) SendOrderToCourierService(ctx context.Context, order *courierProto.OrderFields) (*courierProto.OrderStatusResponse, error) {
	res, err := g.service.OrderApp.GetOrderStatusByID(int(order.OrderId))
	if err != nil {
		logrus.Errorf("GetStatusById:%s", err)
		return nil, fmt.Errorf("GetStatusById:%w", err)
	}

	return res, nil
}
