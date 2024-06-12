package cart

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/charliegreeny/checkout/pkg/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_service_GetById(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		want     *Entity
		stubRows *sqlmock.Rows
		wantErr  error
	}{
		{
			name:     "Valid id, entity returned, nil error",
			id:       "cart1",
			stubRows: sqlmock.NewRows([]string{"id", "total_price", "is_complete"}).AddRow("cart1", 100, true),
			want:     &Entity{ID: "cart1", TotalPrice: 100, IsComplete: true},
			wantErr:  nil,
		},
		{
			name:     "Invalid id, entity nil, ErrNotFound returned",
			id:       "cart1",
			stubRows: sqlmock.NewRows([]string{"id", "total_price", "is_complete"}),
			want:     nil,
			wantErr:  &model.ErrNotFound{},
		},
		{
			name:     "Random db error, nil entity, random db error return",
			id:       "cart1",
			stubRows: sqlmock.NewRows([]string{"id", "total_price", "is_complete"}),
			want:     nil,
			wantErr:  gorm.ErrInvalidData,
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
			s := service{
				db: gormDb,
			}
			if err != nil {
				t.Fatalf("an error '%s' with using sqlmock with gorm", err)
			}
			m := mock.ExpectQuery("SELECT (.+) FROM `carts` WHERE id =(.+)").
				WithArgs(tt.id, 1).WillReturnRows(tt.stubRows)

			if tt.wantErr != nil {
				m.WillReturnError(tt.wantErr)
			}

			got, err := s.GetById(tt.id)

			if err != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			}
			assert.Equalf(t, tt.want, got, "Entity returned not as expected")
		})
	}
}
