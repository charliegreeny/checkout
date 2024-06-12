package offer

type Entity struct {
	ID       string `gorm:"primaryKey"`
	Quantity *int
	Price    *int
	Discount *float32
}

func (Entity) TableName() string {
	return "offers"
}

type Applier interface {
	Apply(quantity, price int) int
}
