package service

import (
	"github.com/minio/minio-go"
	"log"
)

func InitClientDO() (*minio.Client, error) {

	ACCESS_KEY := "Z4AK5OV4JRTOPORJRW2V"                        //os.Getenv("ACCESS_KEY")
	SECRET_KEY := "uELoOpfK1rA/LGjDFPK6w0GZQ+fDumGtIMt16RK6Sfg" //os.Getenv("SECRET_KEY")
	endpoint := "fra1.digitaloceanspaces.com"
	ssl := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, ACCESS_KEY, SECRET_KEY, ssl)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return client, nil
}
