package order

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/google/uuid"
	"github.com/tanmaygupta069/order-service-go/config"
	"github.com/tanmaygupta069/order-service-go/internal/holding"
	"github.com/tanmaygupta069/order-service-go/internal/pkg/mysql"
)

var cfg, _ = config.GetConfig()

type OrderService interface {
	PlaceOrder(order *Orders) (*Orders, error)
	DeleteOrder(orderId string) (*mysql.Orders, error)
	GenerateOrderId() string
	GetStockPrice(symbol string) (float64, error)
	IDORCheck(userid, orderId string) (bool, error)
	GetOrderHistory(userId string) ([]*mysql.Orders, error)
	CancelOrder(orderId string) (*mysql.Orders, error)
	CompleteRandomOrders() error
	CompleteOrder(orderId string)(*mysql.Orders,error)
	CheckStockQuantity(userId string,symbol string,quantity int32)(bool,error)
}

type OrderServiceImp struct {
	repo OrderRepository
	holdingService holding.HoldingService
}

func NewOrderService() OrderService {
	return &OrderServiceImp{
		repo: NewOrderRepository(),
		holdingService: holding.NewHoldingService(),
	}
}

func (r *OrderServiceImp) PlaceOrder(order *Orders) (*Orders, error) {
	return r.repo.PlaceOrder(order)
}

func (r *OrderServiceImp) GenerateOrderId() string {
	orderID := uuid.New()
	return orderID.String()
}

func (r *OrderServiceImp) GetStockPrice(symbol string) (float64, error) {

	if price, err := r.repo.GetCachedStockPrice(symbol); price != "" || err == nil {
		return strconv.ParseFloat(price, 64)
	}
	url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", symbol, cfg.StockApiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return 0.00, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var stockResp StockResponse
	err = json.Unmarshal(body, &stockResp)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return 0.00, err
	}
	if stockResp.C <= 0.00 {
		randomFloat := 20 + rand.Float64()*(500-20)
		randomFloat = float64(int(randomFloat*100)) / 100
		stockResp.C = randomFloat
	}
	simulatedPrice := SimulatePrice(stockResp.C)
	r.repo.CacheStockPrice(symbol, strconv.FormatFloat(simulatedPrice, 'f', 2, 64), 1)
	simulatedPrice = math.Round(simulatedPrice*100) / 100

	return simulatedPrice, nil
}

func (r *OrderServiceImp) DeleteOrder(orderId string) (*mysql.Orders, error) {
	order, err := r.repo.GetOrder(orderId)
	if err != nil {
		return nil, fmt.Errorf("error getting order : %v", err)
	}
	if err = r.repo.DeleteOrder(orderId); err != nil {
		return nil, fmt.Errorf("error deleting order : %v", err)
	}
	return order, nil
}

func (r *OrderServiceImp) IDORCheck(userid, orderId string) (bool, error) {
	order, err := r.repo.GetOrder(orderId)
	if err != nil {
		return false, err
	}
	if userid != order.UserId {
		return false, nil
	}
	return true, nil
}

func (r *OrderServiceImp) GetOrderHistory(userId string) ([]*mysql.Orders, error) {
	return r.repo.GetOrders(userId)
}

func (r *OrderServiceImp) CancelOrder(orderId string) (*mysql.Orders, error) {
	order, err := r.repo.GetOrder(orderId)
	if err != nil {
		return nil, err
	}
	return r.repo.UpdateOrderStatus(order, "cancelled")
}

func (r *OrderServiceImp) CompleteRandomOrders() error {
	order, err := r.repo.GetRandomPlacedOrder()
	if err != nil {
		return err
	}
	if rand.Intn(2) == 0 {
		fmt.Printf("Completing order %s\n", order.OrderId)
		_, err = r.repo.UpdateOrderStatus(order, "completed")
		if err != nil {
			return err
		}
		holding:=&mysql.Holdings{
			UserId: order.UserId,
			Symbol: order.Symbol,
			Quantity: order.Quantity,
			TotalPrice: order.TotalPrice,
		}
		r.holdingService.UpdateHoldings(holding,order.OrderType)
	}
	return nil
}

func (r *OrderServiceImp)CompleteOrder(orderId string)(*mysql.Orders,error){
	if r.repo == nil {
		return nil, fmt.Errorf("repo is nil")
	}
	if r.holdingService == nil {
		return nil, fmt.Errorf("holdingService is nil")
	}

	order, err := r.repo.GetOrder(orderId)
	if err != nil {
		return nil, err
	}
	updatedorder,er:=r.repo.UpdateOrderStatus(order, "completed")
	if er!=nil{
		return nil,er
	}
	if order == nil{
		return nil,fmt.Errorf("order nil in complete order")
	}
	holding:=&mysql.Holdings{
		UserId: updatedorder.UserId,
		Symbol: updatedorder.Symbol,
		Quantity: updatedorder.Quantity,
		TotalPrice: updatedorder.TotalPrice,
	}
	return order,r.holdingService.UpdateHoldings(holding,updatedorder.OrderType)
}

func (r *OrderServiceImp)CheckStockQuantity(userId string,symbol string,quantity int32)(bool,error){
	holding,err:=r.holdingService.GetHolding(userId,symbol)
	if err!=nil{
		return false,err
	}
	if holding.Quantity < quantity{
		return false,nil
	}
	return true,nil
}