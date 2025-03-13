package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetConfig()(*Config,error){
	err:=godotenv.Load()
	if err!=nil{
		fmt.Printf("No .env file found : %v",err)
		return nil,err
	}
	config:=&Config{
		ServerConfig: ServerConfig{
			Port: getEnv("PORT"),
		},
		RateLimiterConfig: RateLimiterConfig{
			RateLimit:   getEnvInt("RATE_LIMIT",2),
			BucketSize: getEnvInt("BUCKET_SIZE",10),
		},
		MySqlConfig: MySqlConfig{
			Port: getEnvInt("MYSQL_PORT",3306),
			User: getEnv("MYSQL_USER"),
			Password: getEnv("MYSQL_PASS"),
			Database: getEnv("MYSQL_DB"),
			Host: getEnv("MYSQL_HOST"),
		},
		RedisConfig:RedisConfig{
			Port: getEnv("REDIS_PORT"),
			User: getEnv("REDIS_USER"),
			Db: getEnvInt("REDIS_DB",0),
			Password: getEnv("REDIS_PASS"),
			Host: getEnv("REDIS_HOST"),
		},
		GrpcConfig: GrpcConfig{
			Port: getEnv("ORDER_SERVICE_PORT"),
			User: getEnv("ORDER_SERVICE_USER"),
		},
		JwtSecret : getEnv("JWT_SECRET"),
		StockApiKey: getEnv("STOCK_API_KEY"),
	}
	return config,nil
}

func getEnv(key string)string{
	if val,exisit := os.LookupEnv(key);exisit{
		return val
	}
	fmt.Printf("\n%s not found in .env\n",key)
	return "";
}

func getEnvInt(key string,Default int)int{
	if val,exisit := os.LookupEnv(key);exisit{
		res,_ := strconv.Atoi(val)
		return res
	}
	fmt.Printf("\n%s not found in .env\n",key)
	return Default;
}