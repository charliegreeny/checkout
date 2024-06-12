package item

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/charliegreeny/checkout/pkg/model"
	"github.com/charliegreeny/checkout/pkg/offer"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mockOfferGetter struct {
	e   *offer.Entity
	err error
}

func (m mockOfferGetter) GetById(_ string) (*offer.Entity, error) {
	return m.e, m.err
}

func Test_getter_GetById(t *testing.T) {
	e1 := &Entity{
		SKU:   "item1",
		Price: 1,
		Offer: &offer.Entity{},
	}
	e2 := &Entity{
		SKU:   "item2",
		Price: 7,
		Offer: &offer.Entity{},
	}
	cache := map[string]*Entity{
		"item1": e1,
		"item2": e2,
	}

	tests := []struct {
		name    string
		id      string
		want    *Entity
		wantErr error
	}{
		{
			name:    "id is item1, returns related item",
			id:      "item1",
			want:    e1,
			wantErr: nil,
		},
		{
			name:    "id is item2, returns related item",
			id:      "item2",
			want:    e2,
			wantErr: nil,
		},
		{
			name:    "id is item3, returns ErrNotFound",
			id:      "item3",
			want:    nil,
			wantErr: &model.ErrNotFound{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := getter{
				cache: cache,
				log:   zap.NewNop(),
			}
			got, err := g.GetById(tt.id)
			if err != nil {
				assert.ErrorAs(t, err, tt.wantErr)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getter_AddToCache(t *testing.T) {
	type fields struct {
		cache       map[string]*Entity
		offerGetter model.IDGetter[*offer.Entity]
		db          *gorm.DB
		log         *zap.Logger
	}
	tests := []struct {
		name         string
		offerGetter  mockOfferGetter
		stubRows     *sqlmock.Rows
		stubErr      error
		wantErr      assert.ErrorAssertionFunc
		wantLenCache int
	}{
		{
			name: "2 valid items added to cache, nil error",
			offerGetter: mockOfferGetter{
				e: &offer.Entity{
					ID:       "offer1",
					Quantity: new(int),
					Price:    new(int),
					Discount: new(float32),
				},
				err: nil,
			},
			stubRows:     sqlmock.NewRows([]string{"SKU", "price", "offers_id"}).AddRow("Sku1", 1, nil).AddRow("sku1", 1, "offerId"),
			stubErr:      nil,
			wantErr:      assert.NoError,
			wantLenCache: 2,
		},
		{
			name: "1 valid items added to cache, one with offerId returns ErrNotFound, nil error",
			offerGetter: mockOfferGetter{
				e: &offer.Entity{
					ID:       "offer1",
					Quantity: new(int),
					Price:    new(int),
					Discount: new(float32),
				},
				err: model.ErrNotFound{},
			},
			stubRows:     sqlmock.NewRows([]string{"SKU", "price", "offers_id"}).AddRow("Sku1", 1, nil).AddRow("sku1", 1, "offerId"),
			stubErr:      nil,
			wantErr:      assert.NoError,
			wantLenCache: 2,
		},
		{
			name: "generic error from offersGetter, error returned",
			offerGetter: mockOfferGetter{
				e: &offer.Entity{
					ID:       "offer1",
					Quantity: new(int),
					Price:    new(int),
					Discount: new(float32),
				},
				err: gorm.ErrInvalidData,
			},
			stubRows:     sqlmock.NewRows([]string{"SKU", "price", "offers_id"}).AddRow("Sku1", 1, nil).AddRow("sku1", 1, "offerId"),
			stubErr:     nil ,
			wantErr:      assert.Error,
			wantLenCache: 1,
		},
		{
			name: "generic error from itemGetter, error returned",
			offerGetter: mockOfferGetter{
				e: &offer.Entity{
					ID:       "offer1",
					Quantity: new(int),
					Price:    new(int),
					Discount: new(float32),
				},
				err: nil,
			},
			stubRows:     sqlmock.NewRows([]string{"SKU", "price", "offers_id"}).AddRow("Sku1", 1, nil).AddRow("sku1", 1, "offerId"),
			stubErr:   	  gorm.ErrInvalidData,
			wantErr:      assert.Error,
			wantLenCache: 0,
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
			g := getter{
				cache:       map[string]*Entity{},
				offerGetter: tt.offerGetter,
				db:          gormDb,
				log:         zap.NewNop(),
			}
			if err != nil {
				t.Fatalf("an error '%s' with using sqlmock with gorm", err)
			}

			m := mock.ExpectQuery("SELECT (.+) FROM `items`").WillReturnRows(tt.stubRows)

			if tt.stubErr != nil {
				m.WillReturnError(tt.stubErr)
			}
			err = g.AddToCache()
			if tt.wantErr != nil {
				tt.wantErr(t, err)
			}
			assert.Equal(t, tt.wantLenCache, len(g.cache))
		})
	}
}
