package cartLineItem

type Entity struct {
	CartID     string `gorm:"column:cart_id;primaryKey" json:"cart_id"`
	ItemsSKU   string `gorm:"column:items_SKU;primaryKey" json:"items_SKU"`
	TotalPrice int    `gorm:"column:total_price"`
	Quantity   int    `gorm:"column:quantity;primaryKey" json:"quantity"`
}

func (Entity) TableName() string {
	return "cart_line_items"
}
func (e Entity) ToOutput() *Item {
	return &Item{
		ItemSKU:    e.ItemsSKU,
		Quantity:   e.Quantity,
		TotalPrice: e.TotalPrice,
	}
}

type Item struct {
	ItemSKU    string `json:"itemSKU"`
	Quantity   int    `json:"quantity"`
	TotalPrice int    `json:"totalPrice"`
	CartId string `json:"-"`
}