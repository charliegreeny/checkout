package cart

import "gorm.io/gorm"

type service struct{
	db *gorm.DB
}

func NewService(db *gorm.DB) *service{
	return &service{db}
}

func(s service) GetById() *Entity{
  	return &Entity{}
}