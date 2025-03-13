package config

type Config struct{
	ServerConfig ServerConfig
	RateLimiterConfig RateLimiterConfig
	MySqlConfig MySqlConfig
	RedisConfig RedisConfig
	GrpcConfig GrpcConfig 
	JwtSecret string
	StockApiKey string 
}

type GrpcConfig struct{
	Port string
	User string
}
 
type ServerConfig struct{
	Port string
}

type RedisConfig struct{
	Port string
	User string
	Db int
	Password string
	Host string
}

type RateLimiterConfig struct{
	RateLimit int
	BucketSize int
}

type MySqlConfig struct{
	Port int
    User string
    Password string
    Database string
	Host string
}