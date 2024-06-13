package customer

import (
	"github.com/charliegreeny/checkout/pkg/model"
	"gorm.io/gorm"
)

type creator struct {
	db *gorm.DB
}

func NewCustomerCreator(db *gorm.DB) model.Creator[string, *Entity] {
	return &creator{db}
}

func (c creator) Create(id string) (*Entity, error) {
	e := &Entity{id}

	r := c.db.Create(e)
	if r.Error != nil {
		return nil, r.Error
	}
	return e, nil
}
