package holding

import (
	"errors"

	"github.com/tanmaygupta069/order-service-go/pkg/mysql"
	Redis "github.com/tanmaygupta069/order-service-go/pkg/redis"
	"gorm.io/gorm"
)

type HoldingRepository interface {
	UpdateHoldings(holding *mysql.Holdings) error
	GetHoldings(holding *mysql.Holdings)(*mysql.Holdings,error)
	InsertHolding(holding *mysql.Holdings)(error)
}

type HoldingRepositoryImp struct{
	mysql *mysql.SqlServiceImplementation[mysql.Holdings]
	redis Redis.RedisInterface
}

func NewHoldingRepository()HoldingRepository{
	return &HoldingRepositoryImp{
		mysql : mysql.NewSqlClient[mysql.Holdings](),
		redis:  Redis.NewRedisClient(),
	}
}

func (db *HoldingRepositoryImp)UpdateHoldings(holding *mysql.Holdings)error{
	return db.mysql.Update(holding)
}

func (db *HoldingRepositoryImp)GetHoldings(holding *mysql.Holdings)(*mysql.Holdings,error){
	existingHolding, err := db.mysql.GetOne(map[string]interface{}{
		"symbol":holding.Symbol,
		"user_id":holding.UserId,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil,err
		}
	}
	return existingHolding,nil
}

func (db *HoldingRepositoryImp)InsertHolding(holding *mysql.Holdings)(error){
	return db.mysql.Insert(holding)
}