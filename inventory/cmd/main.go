package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	inventoryV1API "github.com/baryshnikkov/rocket-factory/inventory/internal/api/inventory/v1"
	"github.com/baryshnikkov/rocket-factory/inventory/internal/interceptor"
	inventoryRepository "github.com/baryshnikkov/rocket-factory/inventory/internal/repository/part"
	inventoryService "github.com/baryshnikkov/rocket-factory/inventory/internal/service/part"
	inventoryV1Proto "github.com/baryshnikkov/rocket-factory/shared/pkg/proto/inventory/v1"
)

const grpcPort = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}

	defer func() {
		if err := lis.Close(); err != nil {
			log.Printf("failed to close listener: %v", err)
		}
	}()

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.LoggerInterceptor(),
			interceptor.Validate(),
		),
	)

	repository := inventoryRepository.NewRepository()
	repository.InitParts()
	service := inventoryService.NewService(repository)
	api := inventoryV1API.NewAPI(service)

	inventoryV1Proto.RegisterInventoryServiceServer(grpcServer, api)

	reflection.Register(grpcServer)

	go func() {
		log.Printf("inventory gRPC сервер запущен на порту %d\n", grpcPort)
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑Shutting down  inventory gRPC server...")
	grpcServer.GracefulStop()
	log.Println("✅ Server stopped")
}
