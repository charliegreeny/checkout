package offer

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/charliegreeny/checkout/pkg/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetApplier(t *testing.T) {
	validE := &Entity{
		ID:       "offer1",
		Quantity: Int(1),
		Price:    Int(1),
		Discount: nil,
	}
	invalidE := &Entity{
		ID:       "offer1",
		Quantity: nil,
		Price:    nil,
		Discount: nil,
	}
	tests := []struct {
		name    string
		args    *Entity
		want    Applier
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "valid offer to apply, returns ",
			args:    validE,
			want:    &multiApplier{offer: validE},
			wantErr: assert.NoError,
		},
		{
			name:    "invalid offer to apply, returns ",
			args:    invalidE,
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetApplier(tt.args)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getter_GetById(t *testing.T) {
	e := &Entity{
		ID:       "offer1",
		Quantity: Int(3),
		Price:    Int(3),
		Discount: nil,
	}
	tests := []struct {
		name     string
		id       string
		want     *Entity
		stubRows *sqlmock.Rows
		stubErr error
		wantErr  error
	}{
		{
			name:     "valid id, returns entity ",
			id:       "offer1",
			want:     e,
			stubRows: sqlmock.NewRows([]string{"id", "price", "quantity"}).AddRow("offer1", 3, 3),
			wantErr:  nil,
		},
		{
			name:     "record not found, ErrNotFound returned",
			id:       "offer1",
			want:     nil,
			stubRows: sqlmock.NewRows([]string{"id", "price", "quantity"}),
			stubErr: gorm.ErrRecordNotFound,
			wantErr:  model.ErrNotFound{Err:gorm.ErrRecordNotFound},
		},
		{
			name:     "Random db error, nil entity, random db error retur",
			id:       "offer1",
			want:     nil,
			stubRows: sqlmock.NewRows([]string{"id", "price", "quantity"}),
			stubErr:  gorm.ErrInvalidData,
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
			s := getter{
				db: gormDb,
			}
			if err != nil {
				t.Fatalf("an error '%s' with using sqlmock with gorm", err)
			}
			m := mock.ExpectQuery("SELECT (.+) FROM `offers` WHERE id =(.+)").
				WithArgs(tt.id, 1).WillReturnRows(tt.stubRows)

			if tt.stubErr != nil {
				m.WillReturnError(tt.stubErr)
			}

			got, err := s.GetById(tt.id)

			if err != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			}
			assert.Equalf(t, tt.want, got, "Entity returned not as expected")
		})
	}
}
