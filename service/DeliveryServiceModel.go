package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/minio/minio-go"
	"log"
	"os"
	"strconv"
)

type DeliveryService struct {
	repo dao.Repository
}

func NewDeliveryService(repo dao.Repository) *DeliveryService {
	return &DeliveryService{
		repo: repo,
	}
}
func (s *DeliveryService) CreateDeliveryService(DeliveryService dao.DeliveryService) (int, error) {
	id, err := s.repo.SaveDeliveryServiceInDB(&DeliveryService)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("Error in DeliveryServiceService: %s", err)
	}
	return id, nil
}

func (s *DeliveryService) GetDeliveryServiceById(Id int) (*dao.DeliveryService, error) {
	var service *dao.DeliveryService
	service, err := s.repo.GetDeliveryServiceByIdFromDB(Id)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Error in DeliveryService: %s", err)
	}
	if service.Id == 0 {
		err = errors.New("not found")
		log.Println(err)
		return nil, fmt.Errorf("Error in DeliveryService: %s", err)
	}
	return service, nil
}

func (s *DeliveryService) GetAllDeliveryServices() ([]dao.DeliveryService, error) {
	var services = []dao.DeliveryService{}
	services, err := s.repo.GetAllDeliveryServicesFromDB()
	if err != nil {
		log.Println(err)
		return []dao.DeliveryService{}, fmt.Errorf("Error in DeliveryService: %s", err)
	}
	return services, nil
}

func (s *DeliveryService) UpdateDeliveryService(service dao.DeliveryService) error {
	if err := s.repo.UpdateDeliveryServiceInDB(service); err != nil {
		log.Println(err)
		return fmt.Errorf("Error in DeliveryService: %s", err)
	}
	return nil
}

func (s *DeliveryService) SaveLogoFile(cover []byte, id int) error {
	ACCESS_KEY := os.Getenv("ACCESS_KEY")
	SECRET_KEY := os.Getenv("SECRET_KEY")
	endpoint := "fra1.digitaloceanspaces.com"
	ssl := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, ACCESS_KEY, SECRET_KEY, ssl)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// List all Spaces.
	spaces, err := client.ListBuckets()
	if err != nil {
		log.Println(err)
		return err
	}
	for _, space := range spaces {
		log.Println(space.Name)
	}

	_, err1 := client.PutObject("storage-like-s3", fmt.Sprintf("%s", strconv.Itoa(id)),
		bytes.NewReader(cover), int64(len(cover)), minio.PutObjectOptions{ContentType: "image/jpeg", UserMetadata: map[string]string{"x-amz-acl": "public-read"}})
	if err1 != nil {
		log.Println(err1)
		return err1
	}

	var service dao.DeliveryService
	service.Id = id
	service.Photo = "https://storage-like-s3.fra1.digitaloceanspaces.com/" + strconv.Itoa(id)

	log.Println("https://storage-like-s3.fra1.digitaloceanspaces.com/" + strconv.Itoa(id))

	if err := s.repo.UpdateDeliveryServiceInDB(service); err != nil {
		log.Println(err)
		return fmt.Errorf("Error in DeliveryService: %s", err)
	}

	log.Println("Downloaded cover")
	return nil
}
