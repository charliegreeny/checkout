package customer

import (
	"github.com/charliegreeny/checkout/pkg/model"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func NewCustomerCreator(db *gorm.DB) model.Creator[string, *Entity] {
	return &service{db}
}

func (s service) Create(id string) (*Entity, error) {
	e := &Entity{id}

	r := s.db.Create(e)
	if r.Error != nil {
		return nil, r.Error
	}
	return e, nil
}
