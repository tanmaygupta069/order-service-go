package order

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tanmaygupta069/order-service/config"
	"github.com/tanmaygupta069/order-service/pkg/mysql"
)

var cfg, _ = config.GetConfig()

type OrderService interface {
	PlaceOrder(order *Orders) (*Orders, error)
	DeleteOrder(orderId string)(*mysql.Orders,error) 
	GenerateOrderId() string
	ExtractUserIDFromToken(tokenString string) (string, error)
	GetStockPrice(symbol string) (float64, error)
	IDORCheck(userid, orderId string) bool
	GetOrderHistory(userId string)([]*mysql.Orders,error)
}

type OrderServiceImp struct {
	repo OrderRepository
}

func NewOrderService() OrderService {
	return &OrderServiceImp{
		repo: NewOrderRepository(),
	}
}

func (r *OrderServiceImp) PlaceOrder(order *Orders) (*Orders, error) {
	return r.repo.PlaceOrder(order)
}

func (r *OrderServiceImp) GenerateOrderId() string {
	orderID := uuid.New()
	return orderID.String()
}

func (r *OrderServiceImp) ExtractUserIDFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token signing method is as expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JwtSecret), nil
	})

	if err != nil {
		return "", fmt.Errorf("error parsing token: %v", err)
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email, ok := claims["email"].(string)
		if !ok {
			return "", fmt.Errorf("email not found in token")
		}
		return email, nil
	}

	return "", fmt.Errorf("invalid token")
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

	simulatedPrice := SimulatePrice(stockResp.C)
	r.repo.CacheStockPrice(symbol, strconv.FormatFloat(simulatedPrice, 'f', 2, 64), 1)
	simulatedPrice = math.Round(simulatedPrice*100) / 100

	return simulatedPrice, nil
}

func (r *OrderServiceImp) DeleteOrder(orderId string) (*mysql.Orders,error) {
	order, err := r.repo.GetOrder(orderId)
	if err != nil {
		return nil,fmt.Errorf("error getting order : %v",err)
	}
	if err = r.repo.DeleteOrder(orderId);err!=nil{
		return nil,fmt.Errorf("error deleting order : %v",err)
	}
	return order,nil
}

func (r *OrderServiceImp) IDORCheck(userid, orderId string) bool {
	order, err := r.repo.GetOrder(orderId)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if userid != order.UserId {
		return false
	}
	return true
}

func (r *OrderServiceImp)GetOrderHistory(userId string)([]*mysql.Orders,error){
	return r.repo.GetOrders(userId)
}