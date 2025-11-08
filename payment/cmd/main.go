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

	paymentV1API "github.com/baryshnikkov/rocket-factory/payment/internal/api/payment/v1"
	"github.com/baryshnikkov/rocket-factory/payment/internal/interceptor"
	paymentService "github.com/baryshnikkov/rocket-factory/payment/internal/service/payment"
	paymentV1Proto "github.com/baryshnikkov/rocket-factory/shared/pkg/proto/payment/v1"
)

const grpcPort = 50052

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v", err)
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
		),
	)

	service := paymentService.NewService()
	api := paymentV1API.NewAPI(service)

	paymentV1Proto.RegisterPaymentServiceServer(grpcServer, api)

	reflection.Register(grpcServer)

	go func() {
		log.Printf("payment gRPC сервер запущен на порту %d\n", grpcPort)
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑Shutting down payment gRPC server...")
	grpcServer.GracefulStop()
	log.Println("✅ Server stopped")
}
