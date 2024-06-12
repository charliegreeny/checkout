package item

import "github.com/charliegreeny/checkout/pkg/offer"


type Entity struct {
	SKU      string `gorm:"primaryKey"`
	Price    int
	OffersId *string
	Offer    *offer.Entity `gorm:"-"`
}

func (Entity) TableName() string {
	return "items"
}
