package holding

import (
	"fmt"
	"github.com/tanmaygupta069/order-service-go/pkg/mysql"
	"google.golang.org/grpc/metadata"
)

type HoldingService interface {
	UpdateHoldings(holding *mysql.Holdings, orderType string) error
	GetHolding(userId string, symbol string) (*mysql.Holdings, error)
	GetHoldings(userId string)([]*mysql.Holdings,error)
	GetTokenFromMetadata(md metadata.MD) (string, error)
}

type HoldingServiceImp struct {
	repo HoldingRepository
}

func NewHoldingService() HoldingService {
	return &HoldingServiceImp{
		repo: NewHoldingRepository(),
	}
}

func (r *HoldingServiceImp) UpdateHoldings(holding *mysql.Holdings, orderType string) error {
	exsistingHolding, err := r.repo.GetHolding(holding)
	if err != nil {
		return r.repo.InsertHolding(holding)
	}
	if orderType == "BUY" {
		holding.Quantity += exsistingHolding.Quantity
		holding.TotalPrice += exsistingHolding.TotalPrice
	} else if orderType == "SELL" {
		if exsistingHolding.Quantity < holding.Quantity {
			return fmt.Errorf("can't sell,number of holding for %s is smaller than holdings to be sold", holding.Symbol)
		}
		exsistingHolding.Quantity -= holding.Quantity
		exsistingHolding.TotalPrice -= holding.TotalPrice
	}
	return r.repo.UpdateHoldings(holding)
}

func (r *HoldingServiceImp) GetHolding(userId string, symbol string) (*mysql.Holdings, error) {
	holding, err := r.repo.GetHolding(&mysql.Holdings{
		UserId: userId,
		Symbol: symbol,
	})
	if err != nil {
		return nil, err
	}
	return holding, nil
}

func (r *HoldingServiceImp) GetTokenFromMetadata(md metadata.MD) (string, error) {
	token := md.Get("Authorization")
	if len(token) == 0 {
		return "", fmt.Errorf("no token found")
	}
	return token[0], nil
}
func (r *HoldingServiceImp)GetHoldings(userId string)([]*mysql.Holdings,error){
	return r.repo.GetHoldings(&mysql.Holdings{
		UserId: userId,
	})
}