package order

import (
	"errors"
	"fmt"

	"github.com/tanmaygupta069/order-service-go/pkg/mysql"
	Redis "github.com/tanmaygupta069/order-service-go/pkg/redis"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type OrderRepository interface {
	PlaceOrder(order *Orders) (*Orders, error)
	CacheStockPrice(symbol, price string, exp int) error
	GetCachedStockPrice(symbol string) (string, error)
	DeleteOrder(orderId string) error
	GetOrder(orderId string) (*mysql.Orders, error)
	GetOrders(userId string) ([]*mysql.Orders, error)
	UpdateOrderStatus(orderId string,status string) (*mysql.Orders,error)
}

type OrderRepositoryImp struct {
	mysql       mysql.SqlInterface
	redisClient Redis.RedisInterface
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImp{
		mysql:       mysql.NewSqlClient(),
		redisClient: Redis.NewRedisClient(),
	}
}

func (db *OrderRepositoryImp) PlaceOrder(order *Orders) (*Orders, error) {
	err := db.mysql.Insert(&mysql.Orders{
		OrderId:       order.OrderId,
		UserId:        order.UserId,
		Symbol:        order.Symbol,
		PricePerStock: order.PricePerStock,
		Quantity:      order.Quantity,
		TotalPrice:    order.TotalPrice,
		OrderType:     order.OrderType,
		OrderStatus: order.OrderStatus,
	})
	if err != nil {
		fmt.Printf("error in placing order repo")
		return nil, err
	}
	return order, nil
}

func (db *OrderRepositoryImp) CacheStockPrice(symbol, price string, exp int) error {
	return db.redisClient.Set(symbol, price, exp)
}

func (db *OrderRepositoryImp) GetCachedStockPrice(symbol string) (string, error) {
	price, err := db.redisClient.Get(symbol)
	if err != nil {
		return "", fmt.Errorf("key not found in cache")
	}
	return price, nil
}

func (db *OrderRepositoryImp) DeleteOrder(orderId string) error {
	return db.mysql.Delete(orderId)
}

func (db *OrderRepositoryImp) GetOrder(orderId string) (*mysql.Orders, error) {
	order, err := db.mysql.GetOne(orderId, "order_id")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("no record corrosponding to orderId : %s", orderId)
	}
	return order, nil
}

func (db *OrderRepositoryImp) GetOrders(userId string) ([]*mysql.Orders, error) {
	orders, err := db.mysql.GetAll(userId)
	if err != nil {
		return nil, err
	}

	result := make([]*mysql.Orders, len(orders))
	for i := range orders {
		result[i] = &orders[i]
	}

	return result, nil
}

func (db *OrderRepositoryImp)UpdateOrderStatus(orderId string,status string) (*mysql.Orders,error){
	order,err:=db.GetOrder(orderId)
	if err != nil {
		return nil, err
	}
	_,ok := AllowedTransitions[order.OrderStatus][status]
	if !ok{
		return nil,fmt.Errorf("invalid state change from %s to %s",order.OrderStatus,status)
	}
	err=db.mysql.Update(order)
	if err!=nil{
		return nil,err
	}
	return order,nil
}