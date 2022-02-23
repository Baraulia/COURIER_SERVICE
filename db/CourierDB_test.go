package db

import (
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestCourierPostgres_GetCouriersFromDB(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewCourierPostgres(db)

	tests := []struct {
		name    string
		mock    func()
		want    []SmallInfo
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id_courier", "courier_name", "phone_number", "photo", "surname"}).
					AddRow(1, "Tim", "1038812", "some photo", "").
					AddRow(2, "Kolya", "1022345", "this is photo", "").
					AddRow(3, "Vasya", "12312345", "my sexy photo", "").
					AddRow(14, "Michael", "293845697", "URL for photo", "Kors").
					AddRow(16, "Robert", "1039913", "some url for photo", "Kuchinskiy").
					AddRow(17, "Robert", "1039913", "some url for photo", "Kuchinskiy").
					AddRow(23, "Kseniya", "1039913", "some url for photo", "Lesnikova")

				mock.ExpectQuery(`SELECT id_courier,courier_name,phone_number,photo,surname FROM couriers WHERE (.+)`).
					WillReturnRows(rows)
				mock.ExpectCommit()
			},
			want: []SmallInfo{
				{1, "Tim", "1038812", "some photo", ""},
				{2, "Kolya", "1022345", "this is photo", ""},
				{3, "Vasya", "12312345", "my sexy photo", ""},
				{14, "Michael", "293845697", "URL for photo", "Kors"},
				{16, "Robert", "1039913", "some url for photo", "Kuchinskiy"},
				{17, "Robert", "1039913", "some url for photo", "Kuchinskiy"},
				{23, "Kseniya", "1039913", "some url for photo", "Lesnikova"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			get, err := r.GetCouriersFromDB()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, get)
			}
			assert.NoError(t, mock.ExpectationsWereMet())

		})
	}
}
