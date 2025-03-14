package holding

import (
	"context"
	"fmt"
	"net/http"

	holdingPb "github.com/tanmaygupta069/order-service-go/generated/holding"
	common "github.com/tanmaygupta069/order-service-go/generated/common"
	"github.com/tanmaygupta069/order-service-go/pkg/auth"
	"google.golang.org/grpc/metadata"
)


type HoldingController struct{
	holdingService HoldingService
	auth auth.AuthPackage
	holdingPb.UnimplementedHoldingServiceServer
}

func NewHoldingController()*HoldingController{
	return &HoldingController{
		holdingService: NewHoldingService(),
		auth: auth.NewAuthPackage(),
	}
}


func (s *HoldingController)GetCurrentHoldings(ctx context.Context,req *holdingPb.CurrentHoldingsRequest)(*holdingPb.CurrentHoldingsResponse,error){
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}
	token, err := s.auth.GetTokenFromMetadata(md)
	if err != nil {
		return &holdingPb.CurrentHoldingsResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	email, err := s.auth.ExtractUserIDFromToken(token)
	if err != nil {
		return &holdingPb.CurrentHoldingsResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	res,er := s.holdingService.GetHoldings(email)
	if er!=nil{
		return &holdingPb.CurrentHoldingsResponse{
			Response: &common.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	holdings:=make([]*holdingPb.Holding,0)

	for _,holding :=range res{
		holdings = append(holdings,&holdingPb.Holding{
			Symbol: holding.Symbol,
			Quantity: holding.Quantity,
			TotalPrice: holding.TotalPrice,
		})
	}

	return &holdingPb.CurrentHoldingsResponse{
		Response: &common.Response{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Holdings: holdings,
	}, nil 
}
