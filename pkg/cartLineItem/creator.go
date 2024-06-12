package cartLineItem

import (
	"github.com/charliegreeny/checkout/pkg/item"
	"github.com/charliegreeny/checkout/pkg/model"
	"github.com/charliegreeny/checkout/pkg/offer"
	"gorm.io/gorm"
)

type creator struct{
	db *gorm.DB
	itemGetter model.IDGetter[*item.Entity]
	offerGetter model.IDGetter[*offer.Entity]
}

func NewCreator(db *gorm.DB, ig model.IDGetter[*item.Entity], og model.IDGetter[*offer.Entity]) model.Creator[*Item, *Entity] {
	return &creator{db: db, itemGetter: ig, offerGetter: og}
}

func (c creator) Create(input *Item) (*Entity, error){
	item, err := c.itemGetter.GetById(input.ItemSKU)
	if err != nil {
		return nil, err
	}
	totalPrice := 0 
	if item.Offer != nil {
		a, err := offer.GetApplier(item.Offer)
		if err != nil {
			return nil, err
		}
		totalPrice = a.Apply(input.Quantity, item.Price)
	}
	// if totalPrice == 0 {
	// 	totalPrice = input.Quantity * item.Price
	// }
	return &Entity{
		ItemsSKU: item.SKU,
		CartID: input.CartId,
		Quantity: input.Quantity,
		TotalPrice: totalPrice,
	}, nil
}