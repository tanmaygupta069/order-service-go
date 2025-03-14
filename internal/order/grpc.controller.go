package order

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	pb "github.com/tanmaygupta069/order-service-go/generated"
	"google.golang.org/grpc/metadata"
)

type OrderController interface {
	PlaceOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error)
	CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error)
	GetOrderHistory(ctx context.Context, req *pb.OrderHistoryRequest) (*pb.OrderHistoryResponse, error)
	GetCurrentPrice(ctx context.Context, req *pb.GetCurrentPriceRequest) (*pb.GetCurrentPriceResponse, error)
	pb.UnimplementedOrderServiceServer
}

type OrderControllerImp struct {
	service OrderService
	pb.UnimplementedOrderServiceServer
}

func NewOrderController() *OrderControllerImp {
	return &OrderControllerImp{
		service: NewOrderService(),
	}
}

func (s *OrderControllerImp) PlaceOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}
	token, err := s.service.GetTokenFromMetadata(md)
	if err != nil {
		return &pb.OrderResponse{
			Response: &pb.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	if req.OrderType == "" || req.Symbol == "" {
		return &pb.OrderResponse{
			Response: &pb.Response{
				Code:    http.StatusBadRequest,
				Message: "symbol,ordertype,quantity can't be empty",
			},
		}, nil
	}
	if req.Quantity <= 0 {
		return &pb.OrderResponse{
			Response: &pb.Response{
				Code:    http.StatusBadRequest,
				Message: "quantity must be greater than zero",
			},
		}, nil
	}

	if req.Symbol == "" {
		return &pb.OrderResponse{
			Response: &pb.Response{
				Code:    http.StatusBadRequest,
				Message: "symbol must have atleast 1 letter",
			},
		}, nil
	}

	if strings.ToUpper(req.OrderType) != "BUY" && strings.ToUpper(req.OrderType) != "SELL" {
		return &pb.OrderResponse{
			Response: &pb.Response{
				Code:    http.StatusBadRequest,
				Message: "order type must be either buy or sell",
			},
		}, nil
	}

	email, err := s.service.ExtractUserIDFromToken(token)
	if err != nil {
		return &pb.OrderResponse{
			Response: &pb.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	stockPrice, err := s.service.GetStockPrice(strings.ToUpper(req.Symbol))
	if err != nil {
		return &pb.OrderResponse{
			Response: &pb.Response{
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
		return &pb.OrderResponse{
			Response: &pb.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, err
	}
	return &pb.OrderResponse{
		Order: &pb.Order{
			OrderId:       res.OrderId,
			Symbol:        strings.ToUpper(req.Symbol),
			Quantity:      res.Quantity,
			PricePerStock: res.PricePerStock,
			TotalPrice:    res.TotalPrice,
			OrderType:     res.OrderType,
			OrderStatus:   res.OrderStatus,
		},
		Response: &pb.Response{
			Code:    http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
		},
	}, nil
}

func (s *OrderControllerImp) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}
	token, err := s.service.GetTokenFromMetadata(md)
	if err != nil {
		return &pb.CancelOrderResponse{
			Response: &pb.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	if req.OrderId == "" {
		return &pb.CancelOrderResponse{
			Response: &pb.Response{
				Code:    http.StatusBadRequest,
				Message: "orderId can't be empty",
			},
		}, nil
	}

	if !IsValidUUID(req.OrderId) {
		return &pb.CancelOrderResponse{
			Response: &pb.Response{
				Code:    http.StatusBadRequest,
				Message: "not a valid format for orderId",
			},
		}, nil
	}
	email, err := s.service.ExtractUserIDFromToken(token)
	if err != nil {
		return &pb.CancelOrderResponse{
			Response: &pb.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	valid, err := s.service.IDORCheck(email, req.OrderId)
	if valid == false && err == nil {
		return &pb.CancelOrderResponse{
			Response: &pb.Response{
				Code:    http.StatusUnauthorized,
				Message: "can't cancel order which is not yours",
			},
		}, nil
	} else if err != nil {
		return &pb.CancelOrderResponse{
			Response: &pb.Response{
				Code:    http.StatusNotFound,
				Message: "you have no such order",
			},
		}, nil
	}

	order, err := s.service.CancelOrder(req.OrderId)
	if err != nil {
		return &pb.CancelOrderResponse{
			Response: &pb.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	return &pb.CancelOrderResponse{
		Order: &pb.Order{
			OrderId:       order.OrderId,
			Symbol:        order.Symbol,
			Quantity:      order.Quantity,
			PricePerStock: order.PricePerStock,
			TotalPrice:    order.TotalPrice,
			OrderType:     order.OrderType,
			OrderStatus: order.OrderStatus,
		},
		Response: &pb.Response{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
	}, nil
}

func (s *OrderControllerImp) GetOrderHistory(ctx context.Context, req *pb.OrderHistoryRequest) (*pb.OrderHistoryResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	token, err := s.service.GetTokenFromMetadata(md)
	if err != nil {
		return &pb.OrderHistoryResponse{
			Response: &pb.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	email, err := s.service.ExtractUserIDFromToken(token)
	if err != nil {
		return &pb.OrderHistoryResponse{
			Response: &pb.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	res, err := s.service.GetOrderHistory(email)
	if err != nil {
		return &pb.OrderHistoryResponse{
			Response: &pb.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}
	orders := make([]*pb.Order, 0)

	for _, order := range res {
		orders = append(orders, &pb.Order{
			OrderId:       order.OrderId,
			Symbol:        order.Symbol,
			Quantity:      order.Quantity,
			PricePerStock: order.PricePerStock,
			TotalPrice:    order.TotalPrice,
			OrderType:     order.OrderType,
			OrderStatus: order.OrderStatus,
		})
	}

	return &pb.OrderHistoryResponse{
		Orders: orders,
		Response: &pb.Response{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
	}, nil
}

func (s *OrderControllerImp) GetCurrentPrice(ctx context.Context, req *pb.GetCurrentPriceRequest) (*pb.GetCurrentPriceResponse, error) {

	if req.Symbol == "" {
		return &pb.GetCurrentPriceResponse{
			Response: &pb.Response{
				Code:    http.StatusBadRequest,
				Message: "symbol can't be empty",
			},
		}, nil
	}
	if len(req.Symbol) <= 0 {
		return &pb.GetCurrentPriceResponse{
			Response: &pb.Response{
				Code:    http.StatusBadRequest,
				Message: "symbol must have atleast 1 letter",
			},
		}, nil
	}
	price, err := s.service.GetStockPrice(req.Symbol)
	if err != nil {
		return &pb.GetCurrentPriceResponse{
			Response: &pb.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}
	return &pb.GetCurrentPriceResponse{
		Price: price,
		Response: &pb.Response{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
	}, nil
}
