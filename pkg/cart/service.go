package cart

import (
	"errors"
	"fmt"
	"time"

	"github.com/charliegreeny/checkout/pkg/cartLineItem"
	"github.com/charliegreeny/checkout/pkg/customer"
	"github.com/charliegreeny/checkout/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type service struct {
	db              *gorm.DB
	lineItemCreator model.Creator[*cartLineItem.Item, *cartLineItem.Entity]
	customerCreator model.Creator[string, *customer.Entity]
}

func NewService(db *gorm.DB,
	cc model.Creator[string, *customer.Entity],
	lic model.Creator[*cartLineItem.Item, *cartLineItem.Entity]) model.IDGetterCreator[*createInput, *Entity] {
	return &service{db: db, customerCreator: cc, lineItemCreator: lic}
}

func (s service) GetById(id string) (*Entity, error) {
	var e *Entity
	result := s.db.First(&e, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, model.ErrNotFound{Err: result.Error}
		}
		return nil, result.Error
	}
	return e, nil
}

func (s service) Create(input *createInput) (*Entity, error) {
	if input.CustomerId == "" {
		c, err := s.customerCreator.Create(uuid.NewString())
		if err != nil {
			return nil, fmt.Errorf("could not create customer: %w", err)
		}
		input.CustomerId = c.ID
	}
	now := time.Now()
	e := &Entity{
		ID:          uuid.NewString(),
		CustomersID: input.CustomerId,
		CreatedAt:   now,
		UpdatedAt:   now,
		IsComplete:  false,
	}
	if r := s.db.Create(e); r.Error != nil {
		return nil, fmt.Errorf("could not create cart: %w", r.Error)
	}
	cartPrice := 0
	for _, lineItem := range input.LineItems {
		lineItem.CartId = e.ID
		lineItemEntity, err := s.lineItemCreator.Create(lineItem)
		if err != nil {
			return nil, err
		}
		cartPrice += lineItemEntity.TotalPrice
		e.lineItemEntities = append(e.lineItemEntities, lineItemEntity)
	}
	e.TotalPrice = cartPrice
	if r := s.db.Save(e); r.Error != nil {
		return nil, fmt.Errorf("could not update cart: %w", r.Error)
	}
	return e, nil
}
