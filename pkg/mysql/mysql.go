package mysql

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tanmaygupta069/order-service-go/config"
)

var db *gorm.DB

var once sync.Once

type SqlServiceImplementation[T any] struct {
	db *gorm.DB
}

func NewSqlClient[T any]()*SqlServiceImplementation[T] {
	return &SqlServiceImplementation[T]{
		db: GetSqlClient(),
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
		if err := d.AutoMigrate(&Orders{}, &Holdings{}); err != nil {
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

func (s *SqlServiceImplementation[T]) GetOne(filters map[string]interface{}) (*T, error) {
	var entity T
	query := s.db
	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	if err := query.First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// ✅ Get all records using a specific filter
func (s *SqlServiceImplementation[T]) GetAll(filters map[string]interface{}) ([]T, error) {
	var entities []T
	query := s.db
	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	if err := query.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// ✅ Update a record
func (s *SqlServiceImplementation[T]) Update(data *T) error {
	return s.db.Save(data).Error
}

// ✅ Get one record randomly based on a filter
func (s *SqlServiceImplementation[T]) GetOneRandomly(filters map[string]interface{}) (*T, error) {
	var entity T
	query := s.db
	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	err := query.Order("RAND()").Limit(1).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (s *SqlServiceImplementation[T]) Insert(data *T) error {
	return s.db.Create(data).Error
}

func (s *SqlServiceImplementation[T]) Delete(filters map[string]interface{}) error {
	var entity T
	query := s.db
	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	err:=query.Delete(&entity).Error
	return err
}
