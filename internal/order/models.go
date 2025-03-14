package order

type Orders struct{
	UserId string
	OrderId string
	Symbol string
	PricePerStock float64
	Quantity int32
	TotalPrice float64
	OrderType string
	OrderStatus string
}

type StockResponse struct {
	C float64 `json:"c"` // `c` is the current price
}