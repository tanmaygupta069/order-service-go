package Redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tanmaygupta069/order-service-go/config"
)

var once sync.Once

var redisClient *redis.Client

type RedisInterface interface {
	Get(key string) (string, error)
	Set(key string, value string, exp int) error
	Delete(key string) (int64, error)
	Exists(key string) (int64, error)
}

type RedisServiceImplementation struct {
	redisClient *redis.Client
}

func NewRedisClient() RedisInterface {
	return &RedisServiceImplementation{
		redisClient: GetRedisClient(),
	}
}

func InitializeRedisClient() {
	once.Do(func() {
		cfg, er := config.GetConfig()
		if er != nil {
			fmt.Println("an error occured.")
		}
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.RedisConfig.Host, cfg.RedisConfig.Port),
			Password: cfg.RedisConfig.Password, // no password set
			DB:       cfg.RedisConfig.Db,       // use default DB
		})

		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			fmt.Printf("error in pinging redis client : %v", err)
		}
		redisClient = client
	})
}

func (r *RedisServiceImplementation) Get(key string) (string, error) {
	val := redisClient.Get(context.Background(), key).Val()
	if val == "" || val == redis.Nil.Error() {
		return "", fmt.Errorf("key not found in cache")
	}
	return val, nil
}

func (r *RedisServiceImplementation) Set(key string, value string, exp int) error {
	err := redisClient.Set(context.Background(), key, value, time.Duration(exp)*time.Minute).Err()
	if err != nil {
		return err
	}
	return redisClient.Expire(context.Background(), key, time.Duration(exp)*time.Minute).Err()
}

func (r *RedisServiceImplementation) Delete(key string) (int64, error) {
	return redisClient.Del(context.Background(), key).Result()
}

func (r *RedisServiceImplementation) Exists(key string) (int64, error) {
	return redisClient.Exists(context.Background(), key).Result()
}

func (r *RedisServiceImplementation) Decrement(key string, field string) (int64, error) {
	return redisClient.HIncrBy(context.Background(), key, field, -1).Result()
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		InitializeRedisClient()
	}
	return redisClient
}
