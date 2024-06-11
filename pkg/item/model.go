package item


type Entity struct{
	SKU string `gorm:"primaryKey"`
	Price int 
	OffersId *string
}

func(Entity) TableName() string{
	return "items"
}

