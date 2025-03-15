package mysql

type Orders struct {
	OrderId       string `gorm:"primaryKey"`
	UserId        string
	Symbol        string
	PricePerStock float64
	Quantity      int32
	TotalPrice    float64
	OrderType     string
	OrderStatus   string
}

type Holdings struct {
	UserId     string `gorm:"primaryKey"`
	Symbol     string `gorm:"primaryKey"`
	Quantity   int32
	TotalPrice float64
}
