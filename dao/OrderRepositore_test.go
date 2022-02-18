package dao

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestRepository_AssigningOrderToCourier_InDB(t *testing.T) {
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
				rows := sqlmock.NewRows([]string{"courier_id", "id"}).
					AddRow(8, 8)

				mock.ExpectQuery("UPDATE delivery SET (.+)").
					WillReturnRows(rows)
				mock.ExpectCommit()
			},
			InputOrder: Order{
				0, 8, 8, "", "", "", "", "", false,
			},
			InputId:        8,
			InputCourierId: 8,
			expectedError:  nil,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.InputOrder)
			got := r.AssigningOrderToCourier_InDB(tt.InputOrder)
			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedError, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
