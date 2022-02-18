package dao

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestRepository_SaveDeliveryService_InDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db)

	testTable := []struct {
		name            string
		mock            func(service *DeliveryService)
		InputService    *DeliveryService
		expectedService *DeliveryService
		expectedError   error
	}{
		{
			name: "OK",
			mock: func(service *DeliveryService) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1)

				mock.ExpectQuery("INSERT INTO delivery_service").
					WillReturnRows(rows)

				mock.ExpectCommit()
			},
			InputService: &DeliveryService{
				Name:        "del",
				Email:       "del@",
				WorkingNow:  true,
				Description: "info",
				Deleted:     false,
			},
			expectedService: &DeliveryService{
				Id: 1,
			},
			expectedError: nil,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.InputService)
			got, err := r.SaveDeliveryService_InDB(tt.InputService)
			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedService, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
