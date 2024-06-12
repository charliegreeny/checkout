package cart

import (
	"time"

	"github.com/charliegreeny/checkout/pkg/cartLineItem"
)

type Entity struct {
	ID               string           `gorm:"column:id;primaryKey" json:"id"`
	CustomersID      string           `gorm:"column:customers_id;primaryKey" json:"customers_id"`
	TotalPrice       int              `gorm:"column:total_price"`
	UpdatedAt        time.Time        `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt        time.Time        `gorm:"column:created_at" json:"created_at"`
	IsComplete       bool             `gorm:"column:is_complete" json:"is_completed"`
	lineItemEntities []*cartLineItem.Entity `gorm:"-"`
}

func (Entity) TableName() string {
	return "carts"
}

func (e Entity) ToOutput() *output {
	var lineItems []*cartLineItem.Item
	for _, lineItemEntity := range e.lineItemEntities {
		lineItems = append(lineItems, lineItemEntity.ToOutput())
	}
	return &output{
		ID:          e.ID,
		CustomerId:  e.CustomersID,
		LineItems:   lineItems,
		UpdatedAt:   e.UpdatedAt,
		CreatedAt:   e.CreatedAt,
		IsCompleted: e.IsComplete,
		TotalPrice:  e.TotalPrice,
	}
}

type createInput struct {
	CustomerId string     `json:"customerId"`
	LineItems  []*cartLineItem.Item `json:"lineItems"`
}
type output struct {
	ID          string      `json:"id"`
	CustomerId  string      `json:"customerId"`
	LineItems   []*cartLineItem.Item `json:"lineItems"`
	TotalPrice  int         `json:"totalPrice"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	CreatedAt   time.Time   `json:"createdAt"`
	IsCompleted bool        `json:"isCompleted"`
}
