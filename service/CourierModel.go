package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/minio/minio-go"
	"log"
	"strconv"
)

type CourierService struct {
	repo dao.Repository
}

func NewCourierService(repo dao.Repository) *CourierService {
	return &CourierService{repo: repo}
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

func (s *CourierService) GetCourier(id int) (dao.Courier, error) {
	get, err := s.repo.GetCourierFromDB(id)
	if (get == dao.Courier{}) {
		return dao.Courier{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return dao.Courier{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	if get.Deleted == true {
		err := errors.New("account deleted")
		log.Println("account deleted")
		return dao.Courier{}, err
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

func (s *CourierService) UpdateCourier(id uint16, status bool) (uint16, error) {
	courierId, err := s.repo.CourierRep.UpdateCourierInDB(id, status)
	if err != nil {
		return 0, fmt.Errorf("Error with database: %s", err)
	}
	return courierId, nil
}

func (s *CourierService) SaveCourierPhoto(cover []byte, id int) error {
	client, err := InitClientDO()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err1 := client.PutObject("storage-like-s3", fmt.Sprintf("courier_photo/%s", strconv.Itoa(id)),
		bytes.NewReader(cover), int64(len(cover)), minio.PutObjectOptions{ContentType: "image/jpeg", UserMetadata: map[string]string{"x-amz-acl": "public-read"}})
	if err1 != nil {
		log.Println(err1)
		return err1
	}
	var courier dao.Courier
	courier.Id = uint16(id)
	courier.Photo = "https://storage-like-s3.fra1.digitaloceanspaces.com/courier_photo/" + strconv.Itoa(id)

	if err := s.repo.UpdateCourierDB(courier); err != nil {
		log.Println(err)
		return fmt.Errorf("Error in DeliveryService: %s", err)
	}

	log.Println("Uploaded logo with link https://storage-like-s3.fra1.digitaloceanspaces.com/courier_photo/" + strconv.Itoa(id))
	return nil
}

func (s *CourierService) GetCouriersOfCourierService(limit, page, idService int) ([]dao.Courier, error) {
	var Couriers = []dao.Courier{}
	Couriers, totalCount := s.repo.GetCouriersOfCourierServiceFromDB(limit, page, idService)
	LimitOfPages := (totalCount / limit) + 1
	if LimitOfPages < page {
		err := errors.New("no page")
		log.Println("no more pages")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	return Couriers, nil
}
