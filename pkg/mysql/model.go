package mysql

type Orders struct {
	OrderId       string `gorm:"primaryKey"`
	UserId        string
	Symbol        string
	PricePerStock float64
	Quantity      int32
	TotalPrice    float64
	OrderType     string
}
