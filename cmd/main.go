package main

import (
	// "fmt"
	"log"
	"net"

	"github.com/tanmaygupta069/order-service-go/config"
	OrderPb "github.com/tanmaygupta069/order-service-go/generated/order"
	HoldingPb "github.com/tanmaygupta069/order-service-go/generated/holding"
	"github.com/tanmaygupta069/order-service-go/internal/holding"
	"github.com/tanmaygupta069/order-service-go/internal/order"

	// "github.com/tanmaygupta069/order-service-go/pkg/mysql"

	// Redis "github.com/tanmaygupta069/order-service-go/pkg/redis"
	"google.golang.org/grpc"

	// "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Redis.InitializeRedisClient()
	// mysql.InitializeSqlClient()
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
	orderController := order.NewOrderController()
	if err != nil {
		log.Fatalf("Failed to load TLS keys: %v", err)
	}
	holdingController := holding.NewHoldingController()
	if err != nil {
		log.Fatalf("Failed to load TLS keys: %v", err)
	}
	listener, err := net.Listen("tcp4", ":"+cfg.GrpcConfig.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// _, er := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	// if er != nil {
	// 	fmt.Printf("error in parsing certificate")
	// }
	grpcServer := grpc.NewServer()
	OrderPb.RegisterOrderServiceServer(grpcServer, orderController)
	HoldingPb.RegisterHoldingServiceServer(grpcServer,holdingController)
	reflection.Register(grpcServer)



	log.Printf("gRPC server is running on port %s", cfg.GrpcConfig.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
