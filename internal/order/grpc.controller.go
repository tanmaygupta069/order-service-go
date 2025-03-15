package order

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	OrderPb "github.com/tanmaygupta069/order-service-go/generated/order"
	common "github.com/tanmaygupta069/order-service-go/generated/common"
	"github.com/tanmaygupta069/order-service-go/internal/pkg/auth"
	"google.golang.org/grpc/metadata"
)

type OrderController struct {
	service OrderService
	auth auth.AuthPackage
	OrderPb.UnimplementedOrderServiceServer
}

func NewOrderController() *OrderController {
	return &OrderController{
		service: NewOrderService(),
		auth: auth.NewAuthPackage(),
	}
}

func (s *OrderController) PlaceOrder(ctx context.Context, req *OrderPb.OrderRequest) (*OrderPb.OrderResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}
	token, err := s.auth.GetTokenFromMetadata(md)
	if err != nil {
		return &OrderPb.OrderResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	req.OrderType=strings.ToUpper(req.OrderType)
	if req.OrderType == "" || req.Symbol == "" {
		return &OrderPb.OrderResponse{
			Response: &common.Response{
				Code:    http.StatusBadRequest,
				Message: "symbol,ordertype,quantity can't be empty",
			},
		}, nil
	}
	if req.Quantity <= 0 {
		return &OrderPb.OrderResponse{
			Response: &common.Response{
				Code:    http.StatusBadRequest,
				Message: "quantity must be greater than zero",
			},
		}, nil
	}

	if req.Symbol == "" {
		return &OrderPb.OrderResponse{
			Response: &common.Response{
				Code:    http.StatusBadRequest,
				Message: "symbol must have atleast 1 letter",
			},
		}, nil
	}

	if strings.ToUpper(req.OrderType) != "BUY" && strings.ToUpper(req.OrderType) != "SELL" {
		return &OrderPb.OrderResponse{
			Response: &common.Response{
				Code:    http.StatusBadRequest,
				Message: "order type must be either buy or sell",
			},
		}, nil
	}

	email, err := s.auth.ExtractUserIDFromToken(token)
	if err != nil {
		return &OrderPb.OrderResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	if req.OrderType == "sell" || req.OrderType == "SELL"{
		ok,err:=s.service.CheckStockQuantity(email,req.Symbol,req.Quantity);
		if err!=nil{
			return &OrderPb.OrderResponse{
				Response: &common.Response{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			}, nil
		}
		if ok==false && err==nil{
			return &OrderPb.OrderResponse{
				Response: &common.Response{
					Code:    http.StatusBadRequest,
					Message: "can't sell less than your current quantity",
				},
			}, nil
		}
	} 
	stockPrice, err := s.service.GetStockPrice(strings.ToUpper(req.Symbol))
	if err != nil {
		return &OrderPb.OrderResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}
	order := Orders{
		OrderId:       s.service.GenerateOrderId(),
		UserId:        email,
		Symbol:        strings.ToUpper(req.Symbol),
		PricePerStock: stockPrice,
		Quantity:      req.Quantity,
		TotalPrice:    stockPrice * float64(req.Quantity),
		OrderType:     req.OrderType,
		OrderStatus:   "placed",
	}
	res, err := s.service.PlaceOrder(&order)
	if err != nil {
		return &OrderPb.OrderResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, err
	}
	return &OrderPb.OrderResponse{
		Order: &OrderPb.Order{
			OrderId:       res.OrderId,
			Symbol:        strings.ToUpper(req.Symbol),
			Quantity:      res.Quantity,
			PricePerStock: res.PricePerStock,
			TotalPrice:    res.TotalPrice,
			OrderType:     res.OrderType,
			OrderStatus:   res.OrderStatus,
		},
		Response: &common.Response{
			Code:    http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
		},
	}, nil
}

func (s *OrderController) CancelOrder(ctx context.Context, req *OrderPb.CancelOrderRequest) (*OrderPb.CancelOrderResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}
	token, err := s.auth.GetTokenFromMetadata(md)
	if err != nil {
		return &OrderPb.CancelOrderResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	if req.OrderId == "" {
		return &OrderPb.CancelOrderResponse{
			Response: &common.Response{
				Code:    http.StatusBadRequest,
				Message: "orderId can't be empty",
			},
		}, nil
	}
	fmt.Print(req.OrderId)

	if !IsValidUUID(req.OrderId) {
		return &OrderPb.CancelOrderResponse{
			Response: &common.Response{
				Code:    http.StatusBadRequest,
				Message: "not a valid format for orderId",
			},
		}, nil
	}
	email, err := s.auth.ExtractUserIDFromToken(token)
	if err != nil {
		return &OrderPb.CancelOrderResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	valid, err := s.service.IDORCheck(email, req.OrderId)
	if valid == false && err == nil {
		return &OrderPb.CancelOrderResponse{
			Response: &common.Response{
				Code:    http.StatusUnauthorized,
				Message: "can't cancel order which is not yours",
			},
		}, nil
	} else if err != nil {
		return &OrderPb.CancelOrderResponse{
			Response: &common.Response{
				Code:    http.StatusNotFound,
				Message: "you have no such order",
			},
		}, nil
	}

	order, err := s.service.CancelOrder(req.OrderId)
	if err != nil {
		return &OrderPb.CancelOrderResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	return &OrderPb.CancelOrderResponse{
		Order: &OrderPb.Order{
			OrderId:       order.OrderId,
			Symbol:        order.Symbol,
			Quantity:      order.Quantity,
			PricePerStock: order.PricePerStock,
			TotalPrice:    order.TotalPrice,
			OrderType:     order.OrderType,
			OrderStatus: order.OrderStatus,
		},
		Response: &common.Response{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
	}, nil
}

func (s *OrderController) GetOrderHistory(ctx context.Context, req *OrderPb.OrderHistoryRequest) (*OrderPb.OrderHistoryResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	token, err := s.auth.GetTokenFromMetadata(md)
	if err != nil {
		return &OrderPb.OrderHistoryResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	email, err := s.auth.ExtractUserIDFromToken(token)
	if err != nil {
		return &OrderPb.OrderHistoryResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	res, err := s.service.GetOrderHistory(email)
	if err != nil {
		return &OrderPb.OrderHistoryResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}
	orders := make([]*OrderPb.Order, 0)

	for _, order := range res {
		orders = append(orders, &OrderPb.Order{
			OrderId:       order.OrderId,
			Symbol:        order.Symbol,
			Quantity:      order.Quantity,
			PricePerStock: order.PricePerStock,
			TotalPrice:    order.TotalPrice,
			OrderType:     order.OrderType,
			OrderStatus: order.OrderStatus,
		})
	}

	return &OrderPb.OrderHistoryResponse{
		Orders: orders,
		Response: &common.Response{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
	}, nil
}

func (s *OrderController) GetCurrentPrice(ctx context.Context, req *OrderPb.GetCurrentPriceRequest) (*OrderPb.GetCurrentPriceResponse, error) {

	if req.Symbol == "" {
		return &OrderPb.GetCurrentPriceResponse{
			Response: &common.Response{
				Code:    http.StatusBadRequest,
				Message: "symbol can't be empty",
			},
		}, nil
	}
	if len(req.Symbol) <= 0 {
		return &OrderPb.GetCurrentPriceResponse{
			Response: &common.Response{
				Code:    http.StatusBadRequest,
				Message: "symbol must have atleast 1 letter",
			},
		}, nil
	}

	req.Symbol=strings.ToUpper(req.Symbol)
	price, err := s.service.GetStockPrice(req.Symbol)
	if err != nil {
		return &OrderPb.GetCurrentPriceResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}
	return &OrderPb.GetCurrentPriceResponse{
		Price: price,
		Response: &common.Response{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
	}, nil
}

func (s *OrderController)CompleteOrder(ctx context.Context,req *OrderPb.CompleteOrderRequest)(*OrderPb.CompleteOrderResponse,error){
	if req.OrderId == "" {
		return &OrderPb.CompleteOrderResponse{
			Response: &common.Response{
				Code:    http.StatusBadRequest,
				Message: "orderId can't be empty",
			},
		}, nil
	}

	if !IsValidUUID(req.OrderId) {
		return &OrderPb.CompleteOrderResponse{
			Response: &common.Response{
				Code:    http.StatusBadRequest,
				Message: "not a valid format for orderId",
			},
		}, nil
	}

	order, err := s.service.CompleteOrder(req.OrderId)
	if err != nil {
		return &OrderPb.CompleteOrderResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	return &OrderPb.CompleteOrderResponse{
		Response: &common.Response{
			Code: http.StatusOK,
			Message: "ordered completed",
		},
		Order: &OrderPb.Order{
			OrderId: order.OrderId,
			Symbol: order.Symbol,
			Quantity: order.Quantity,
			PricePerStock: order.PricePerStock,
			TotalPrice: order.TotalPrice,
			OrderType: order.OrderType,
			OrderStatus: order.OrderStatus,
		},
	},nil
}