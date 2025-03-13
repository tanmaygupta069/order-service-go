package mysql

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tanmaygupta069/order-service/config"
)

var db *gorm.DB

var once sync.Once

type SqlInterface interface {
	Insert(order *Orders) error
	Delete(OrderId string) error
	GetOne(key string,value string) (*Orders, error)
	GetAll(userId string) ([]Orders, error)
}

type SqlServiceImplementation struct {
	db *gorm.DB
}

func NewSqlClient() *SqlServiceImplementation {
	return &SqlServiceImplementation{
		db:GetSqlClient(),
	}
}

func InitializeSqlClient() {
	once.Do(func() {
		// fmt.Print("in initialize")
		cfg, er := config.GetConfig()
		if er != nil {
			fmt.Println("error occured in sql client init")
		}
		dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.MySqlConfig.User, cfg.MySqlConfig.Password, cfg.MySqlConfig.Host, cfg.MySqlConfig.Port, cfg.MySqlConfig.Database)
		d, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
		if err != nil {
			fmt.Println("error occured while connecting to mysql")
		}
		sqlDB, err := d.DB()
		if err != nil {
			fmt.Println("Failed to get DB instance:", err)
			return
		}
		if err := sqlDB.Ping(); err != nil {
			fmt.Println("Failed to ping DB:", err)
			return
		}

		if err := d.AutoMigrate(&Orders{}); err != nil {
			fmt.Println("Failed to auto-migrate:", err)
			return
		}

		fmt.Println("DB connection successful")
		db = d
	})
}

func GetSqlClient() *gorm.DB {
	if db == nil {
		InitializeSqlClient()
	}
	return db
}

func (s *SqlServiceImplementation) GetOne(key string,value string) (*Orders, error) {
	var order Orders
	err := db.Where(fmt.Sprintf("%s = ?", value),key).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *SqlServiceImplementation) GetAll(userId string) ([]Orders, error) {
	var orders []Orders
	if err := db.Where("user_id = ?",userId).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *SqlServiceImplementation) Insert(order *Orders) error {
	return db.Create(order).Error
}

func (s *SqlServiceImplementation) Delete(OrderId string) error {
	return db.Delete(&Orders{}, "order_id = ?",OrderId).Error
}
