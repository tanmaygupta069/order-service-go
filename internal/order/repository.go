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
	UpdateOrderStatus(order *mysql.Orders,status string) (*mysql.Orders,error)
	GetRandomPlacedOrder() (*mysql.Orders, error)
}

type OrderRepositoryImp struct {
	mysql       *mysql.SqlServiceImplementation[mysql.Orders]
	redisClient Redis.RedisInterface
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImp{
		mysql:       mysql.NewSqlClient[mysql.Orders](),
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
	filter:=map[string]interface{}{
		"order_id":orderId,
	}
	return db.mysql.Delete(filter)
}

func (db *OrderRepositoryImp) GetOrder(orderId string) (*mysql.Orders, error) {
	order, err := db.mysql.GetOne(map[string]interface{}{
		"order_id":orderId,
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("no record corrosponding to orderId : %s", orderId)
	}
	return order, nil
}

func (db *OrderRepositoryImp) GetOrders(userId string) ([]*mysql.Orders, error) {
	orders, err := db.mysql.GetAll(map[string]interface{}{
		"user_id":userId,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*mysql.Orders, len(orders))
	for i := range orders {
		result[i] = &orders[i]
	}

	return result, nil
}

func (db *OrderRepositoryImp)UpdateOrderStatus(order *mysql.Orders,status string) (*mysql.Orders,error){
	_,ok := AllowedTransitions[order.OrderStatus][status]
	if !ok{
		return nil,fmt.Errorf("invalid state change from %s to %s",order.OrderStatus,status)
	}
	order.OrderStatus=status
	err:=db.mysql.Update(order)
	if err!=nil{
		return nil,err
	}
	return order,nil
}

func (r *OrderRepositoryImp) GetRandomPlacedOrder() (*mysql.Orders, error) {
    order,err:=r.mysql.GetOneRandomly(map[string]interface{}{
		"order_status":"placed",
	})
	if err!=nil{
		return nil,err
	}
	return order,nil
}
