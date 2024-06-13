package cartLineItem

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/charliegreeny/checkout/pkg/item"
	"github.com/charliegreeny/checkout/pkg/model"
	"github.com/charliegreeny/checkout/pkg/offer"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func s(s string) *string {
	return &s
}

func i(i int) *int {
	return &i
}

type mockItemGetter struct {
	i   *item.Entity
	err error
}

func (m mockItemGetter) GetById(_ string) (*item.Entity, error) {
	return m.i, m.err
}

func Test_creator_Create(t *testing.T) {
	tests := []struct {
		name           string
		arg            *Item
		want           *Entity
		noInsert       bool
		mockItemGetter mockItemGetter
		stubErr        error
		wantErr        assert.ErrorAssertionFunc
	}{
		{
			name:           "creates carteLineItem from item w/ 1 quantity and no offer, returns entity and nil error",
			arg:            &Item{
				ItemSKU:    "A",
				Quantity:   1,
				TotalPrice: 0,
				CartId:     "cartId1",
			},
			want:           &Entity{
				CartID:     "cartId1",
				ItemsSKU:   "A",
				TotalPrice: 10,
				Quantity:   1,
			},
			mockItemGetter: mockItemGetter{
				i:   &item.Entity{
					SKU:      "A",
					Price:    10,
					OffersId: s("offerId"),
					Offer:    nil,
				},
				err: nil,
			},
			stubErr:        nil,
			wantErr: 		assert.NoError,
		},
		{
			name:           "creates carteLineItem from item w/ 2 quantity and no offer, returns entity and nil error",
			arg:            &Item{
				ItemSKU:    "A",
				Quantity:   2,
				TotalPrice: 0,
				CartId:     "cartId1",
			},
			want:           &Entity{
				CartID:     "cartId1",
				ItemsSKU:   "A",
				TotalPrice: 20,
				Quantity:   2,
			},
			mockItemGetter: mockItemGetter{
				i:   &item.Entity{
					SKU:      "A",
					Price:    10,
					OffersId: nil,
					Offer:    nil,
				},
				err: nil,
			},
			stubErr:        nil,
			wantErr: 		assert.NoError,
		},
		{
			name:           "creates carteLineItem from item w/ 2 quantity and offer, returns entity with offer applied and nil error",
			arg:            &Item{
				ItemSKU:    "A",
				Quantity:   2,
				TotalPrice: 0,
				CartId:     "cartId1",
			},
			want:           &Entity{
				CartID:     "cartId1",
				ItemsSKU:   "A",
				TotalPrice: 15,
				Quantity:   2,
			},
			mockItemGetter: mockItemGetter{
				i:   &item.Entity{
					SKU:      "A",
					Price:    10,
					OffersId: s("offerId"),
					Offer:    &offer.Entity{
						ID:       "offerId",
						Quantity: i(2),
						Price:    i(15),
						Discount: nil,
					},
				},
				err: nil,
			},
			stubErr:        nil,
			wantErr: 		assert.NoError,
		},
		{
			name:           "creates carteLineItem from item w/ 2 quantity and offer, itemGetter returns ErrNotFound, nil entity ErrNotFound return",
			arg:            &Item{
				ItemSKU:    "A",
				Quantity:   2,
				TotalPrice: 0,
				CartId:     "cartId1",
			},
			want:         nil,
			mockItemGetter: mockItemGetter{
				i:   nil,
				err: model.ErrNotFound{Err: errors.New("not found")},
			},
			noInsert: true,
			stubErr:        nil,
			wantErr: 		func(t assert.TestingT , err error, msgAndArgs ...interface{}) bool {
				return errors.Is(err, model.ErrNotFound{Err: errors.New("not found")})},
		},
		{
			name: "error when inserting cart_line_item, nil entity error return",
			arg: &Item{
				ItemSKU:    "A",
				Quantity:   2,
				TotalPrice: 0,
				CartId:     "cartId1",
			},
			want: nil,
			mockItemGetter: mockItemGetter{
				i: &item.Entity{
					SKU:      "A",
					Price:    10,
					OffersId: s("offerId"),
					Offer: &offer.Entity{
						ID:       "offerId",
						Quantity: i(2),
						Price:    i(15),
						Discount: nil,
					},
				},
				err: nil,
			},
			stubErr: errors.New("DB error"),
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
				db:         gormDb,
				itemGetter: tt.mockItemGetter,
			}
			if !tt.noInsert {
				mock.ExpectBegin()
				m := mock.ExpectExec("INSERT INTO `cart_line_items` (.+) VALUES (.+)").WillReturnResult(sqlmock.NewResult(1, 1))

				if tt.stubErr != nil {
					m.WillReturnError(tt.stubErr)
				}
				mock.ExpectCommit()
				mock.ExpectationsWereMet()
			}
			got, err := c.Create(tt.arg)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
