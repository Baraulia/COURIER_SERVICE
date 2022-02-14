package dao

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"log"
	"testing"
)
func TestRepository_GetCourierCompletedOrdersWithPage_fromDB(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db)

	testTable := []struct {
		name          string
		mock          func(courier_id,limit,page int)
		courier_id    int
		limit         int
		page          int
		expectedOrder  []Order
	}{
		{
			name: "OK",
			mock: func(courier_id,limit,page int) {
			mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"courier_id", "id", "delivery_service_id" ,"delivery_time", "status", "customer_address"}).
					AddRow(1, 1, 1,"12:00:00","completed","address")

				mock.ExpectQuery(`SELECT courier_id,id,delivery_service_id,delivery_time,status,customer_address FROM delivery WHERE status='completed' (.+)`).
					WillReturnRows(rows)

				rows2 := sqlmock.NewRows([]string{"courier_id"}).
					AddRow(1)
				mock.ExpectQuery(`SELECT courier_id FROM delivery WHERE status='completed' (.+)`).
					WillReturnRows(rows2)

				mock.ExpectCommit()
			},
			courier_id: 1,
			limit: 1,
			page: 1,
			expectedOrder:  []Order{
				{
					IdDeliveryService: 1,
					IdOrder:           1,
					IdCourier:         1,
					DeliveryTime:      "12:00:00",
					CustomerAddress:   "address",
					Status:            "completed",
				},
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.courier_id,tt.limit,tt.page)
			got, _ := r.GetCourierCompletedOrdersWithPage_fromDB(tt.courier_id,tt.limit,tt.page)

			assert.Equal(t, tt.expectedOrder, got)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}


func TestRepository_GetAllOrdersOfCourierServiceWithPage_fromDB(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db)

	testTable := []struct {
		name          string
		mock          func(delivery_service_id,limit,page int)
		delivery_service_id    int
		limit         int
		page          int
		expectedOrder  []Order
	}{
		{
			name: "OK",
			mock: func(delivery_service_id,limit,page int) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"courier_id", "id","delivery_time", "status", "customer_address"}).
					AddRow(1, 1,"12:00:00","completed","address")

				mock.ExpectQuery(`SELECT courier_id,id,delivery_time,status,customer_address FROM delivery WHERE (.+)`).
					WillReturnRows(rows)

				rows2 := sqlmock.NewRows([]string{"courier_id"}).
					AddRow(1)
				mock.ExpectQuery(`SELECT courier_id FROM delivery WHERE (.+)`).
					WillReturnRows(rows2)

				mock.ExpectCommit()
			},
			delivery_service_id: 1,
			limit: 1,
			page: 1,
			expectedOrder:  []Order{
				{
					IdOrder:           1,
					IdCourier:         1,
					DeliveryTime:      "12:00:00",
					CustomerAddress:   "address",
					Status:            "completed",
				},
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.delivery_service_id,tt.limit,tt.page)
			got, _ := r.GetAllOrdersOfCourierServiceWithPage_fromDB(tt.delivery_service_id,tt.limit,tt.page)

			assert.Equal(t, tt.expectedOrder, got)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}


func TestRepository_GetCourierCompletedOrdersByMouthWithPage_fromDB(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db)

	testTable := []struct {
		name          string
		mock          func(courier_id,limit,page int)
		courier_id    int
		limit         int
		page          int
		month 		  int
		expectedOrder  []Order
	}{
		{
			name: "OK",
			mock: func(courier_id,limit,page int) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"courier_id", "id", "delivery_service_id" ,"delivery_time", "order_date","status", "customer_address"}).
					AddRow(1, 1, 1,"12:00:00", "2022-02-02","completed","address")

				mock.ExpectQuery(`SELECT courier_id ,id ,delivery_service_id ,delivery_time ,order_date ,status ,customer_address FROM delivery where (.+)`).
					WillReturnRows(rows)

				rows2 := sqlmock.NewRows([]string{"courier_id"}).
					AddRow(1)
				mock.ExpectQuery(`SELECT courier_id FROM delivery WHERE (.+)`).
					WillReturnRows(rows2)

				mock.ExpectCommit()
			},
			courier_id: 1,
			limit: 1,
			page: 1,
			month: 1,
			expectedOrder:  []Order{
				{
					IdDeliveryService: 1,
					IdOrder:           1,
					IdCourier:         1,
					DeliveryTime:      "12:00:00",
					OrderDate:         "2022-02-02",
					CustomerAddress:   "address",
					Status:            "completed",
				},
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.courier_id,tt.limit,tt.page)
			got, _ := r.GetCourierCompletedOrdersByMouthWithPage_fromDB(tt.courier_id,tt.limit,tt.page,tt.month)

			assert.Equal(t, tt.expectedOrder, got)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}


func TestRepository_AssigningOrderToCourier_InDB(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db)

	testTable := []struct {
		name           string
		mock           func(order Order)
		InputOrder     Order
		InputId        int
		InputCourierId int
		expectedError  error
	}{
		{
			name: "OK",
			mock: func(order Order) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"courier_id","id"}).
					AddRow(8,8)

				mock.ExpectQuery("UPDATE delivery SET (.+)").
					WillReturnRows(rows)
				mock.ExpectCommit()
			},
			InputOrder: Order{
				0,8,8,"","","","",
			},
			InputId: 8,
			InputCourierId: 8,
			expectedError: nil,
			},

		}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.InputOrder)
			got := r.AssigningOrderToCourier_InDB(tt.InputOrder)
			if tt.expectedError!=nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedError, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
	}
