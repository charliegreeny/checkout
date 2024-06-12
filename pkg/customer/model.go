package customer

type Entity struct{
	ID string `gorm:"column:id;primaryKey"`
}

func (Entity) TableName()string{
	return "customers"
}