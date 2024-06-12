package item

import (
	"testing"

	"github.com/charliegreeny/checkout/pkg/model"
	"github.com/charliegreeny/checkout/pkg/offer"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// type mockOfferGetter struct{
// 	e *offer.Entity
// 	err error
// }

// func(m mockOfferGetter) GetById(id) (*offer.Entity, error) {
// 	return m.e, error
// }

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
