package customer

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_creator_Create(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		want    *Entity
		stubErr error
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "adds customer to DB, returns entity and nil error ",
			id:      "customer1",
			want:    &Entity{"customer1"},
			stubErr: nil,
			wantErr: assert.NoError,
		},
		{
			name:    "db returns error, returns nil entity and error ",
			id:      "customer1",
			want:   nil,
			stubErr: gorm.ErrInvalidData,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			d := mysql.New(mysql.Config{
				DSN:                       "sqlmock_db_0",
				DriverName:                "mysql",
				Conn:                      db,
				SkipInitializeWithVersion: true,
			})
			gormDb, err := gorm.Open(d, &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' with using sqlmock with gorm", err)
			}
			c := creator{
				db: gormDb,
			}

			mock.ExpectBegin()
			m := mock.ExpectExec("INSERT INTO `customers` (.+) VALUES (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
			
			if tt.stubErr != nil {
				m.WillReturnError(tt.stubErr)
			}
			mock.ExpectCommit()
			mock.ExpectationsWereMet()

			got, err := c.Create(tt.id)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
