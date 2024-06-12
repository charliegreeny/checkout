package cart

import (
	"errors"

	"github.com/charliegreeny/checkout/pkg/model"
	"gorm.io/gorm"
)

type service struct{
	db *gorm.DB
}

func NewService(db *gorm.DB) model.Service[*Entity]{
	return &service{db}
}

func(s service) GetById(id string) (*Entity, error){
	var e *Entity
  	result := s.db.First(&e, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		return nil, model.NotFoundErr{Err: result.Error}
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return e, nil
}