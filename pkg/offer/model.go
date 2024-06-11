package offer

type Entity struct {
	ID string `gorm:"primaryKey"`
	Quantity int 
	Price int 
	Discount *float32
}

func(Entity) TableName() string{
	return "offers"
}