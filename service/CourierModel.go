package service

import (
	"context"
	"errors"
	"fmt"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPC"
	"github.com/Baraulia/COURIER_SERVICE/GRPC/grpcClient"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"log"
	"strings"
)

type CourierService struct {
	repo    dao.Repository
	grpcCli *grpcClient.GRPCClient
}

func NewProjectService(repo dao.Repository, grpcCli *grpcClient.GRPCClient) *CourierService {
	return &CourierService{
		repo:    repo,
		grpcCli: grpcCli,
	}
}

func (s *CourierService) GetCouriers() ([]dao.SmallInfo, error) {
	get, err := s.repo.GetCouriersFromDB()
	if get == nil {
		return []dao.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	if err != nil {
		return nil, fmt.Errorf("Error with database: %s", err)
	}
	return get, nil
}

func (s *CourierService) GetCourier(id int) (dao.SmallInfo, error) {
	get, err := s.repo.GetCourierFromDB(id)
	if (get == dao.SmallInfo{}) {
		return dao.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return dao.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	return get, nil
}

func (s *CourierService) SaveCourier(courier *dao.Courier) (*dao.Courier, error) {
	err := s.repo.SaveCourierInDB(courier)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return courier, nil
}

func (s *CourierService) UpdateCourier(id uint16) (uint16, error) {
	courierId, err := s.repo.CourierRep.UpdateCourierInDB(id)
	if err != nil {
		return 0, fmt.Errorf("Error with database: %s", err)
	}
	return courierId, nil
}

func (s *CourierService) ParseToken(token string) (*courierProto.UserRole, error) {
	return s.grpcCli.GetUserWithRights(context.Background(), &courierProto.AccessToken{AccessToken: token})
}

func (s *CourierService) CheckRoleRights(neededPerms []string, neededRole1 string, neededRole2 string, givenPerms string, givenRole string) error {
	if neededPerms != nil {
		ok := true
		for _, perm := range neededPerms {
			if !strings.Contains(givenPerms, perm) {
				ok = false
				return fmt.Errorf("not enough rights")
			} else {
				continue
			}
		}
		if ok == true {
			return nil
		}
	}
	if neededRole1 != givenRole || neededRole2 != givenRole {
		return fmt.Errorf("not enough rights")
	}
	return nil
}
