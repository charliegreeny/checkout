package cart

import "time"

type Entity struct {
	ID              string    `gorm:"column:id;primaryKey" json:"id"`
	CustomersID     string    `gorm:"column:customers_id;primaryKey" json:"customers_id"`
	TotalPrice 		int  	  `gorm:"column:total_price"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	IsCompleted 	bool 	  `gorm:"column:is_completed" json:"is_completed"`
}

func (Entity) TableName() string {
	return "carts"
}

type LineItemEntity struct {
	CartID   string `gorm:"column:cart_id;primaryKey" json:"cart_id"`
	ItemsSKU string `gorm:"column:items_SKU;primaryKey" json:"items_SKU"`
	TotalPrice int  `gorm:"column:total_price"`
	Quantity int32  `gorm:"column:quantity;primaryKey" json:"quantity"`
}

func(LineItemEntity) TableName() string{
	return "cart_line_items"
}

