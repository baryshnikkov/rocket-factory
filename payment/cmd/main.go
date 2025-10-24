package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/baryshnikkov/rocket-factory/payment/internal/interceptor"
	paymentV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/proto/payment/v1"
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
			grpc.UnaryServerInterceptor(interceptor.LoggerInterceptor()),
		),
	)

	service := NewPaymentService()

	paymentV1.RegisterPaymentServiceServer(grpcServer, service)

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

type PaymentService struct {
	paymentV1.UnimplementedPaymentServiceServer
}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) PayOrder(_ context.Context, _ *paymentV1.PayOrderRequest) (*paymentV1.PayOrderResponse, error) {
	UUID := uuid.New().String()

	log.Printf("✅Оплата прошла успешно, transaction_uuid: %v\n", UUID)

	return &paymentV1.PayOrderResponse{
		TransactionUuid: UUID,
	}, nil
}
