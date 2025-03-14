package holding

import "github.com/tanmaygupta069/order-service-go/pkg/mysql"

type HoldingService interface {
	UpdateHoldings(holding *mysql.Holdings,orderType string)error
}

type HoldingServiceImp struct {
	repo HoldingRepository
}

func NewHoldingService() HoldingService {
	return &HoldingServiceImp{
		repo: NewHoldingRepository(),
	}
}

func (r *HoldingServiceImp)UpdateHoldings(holding *mysql.Holdings,orderType string)error{
	exsistingHolding,err:=r.repo.GetHoldings(holding);
	if err!=nil{
		return r.repo.InsertHolding(holding)
	}

	if orderType=="BUY"{
		holding.Quantity+=exsistingHolding.Quantity
		holding.TotalPrice+=exsistingHolding.TotalPrice
	}else if orderType=="SELL"{
		holding.Quantity-=exsistingHolding.Quantity
		holding.TotalPrice-=exsistingHolding.TotalPrice
	}
	return r.repo.UpdateHoldings(holding)
}
