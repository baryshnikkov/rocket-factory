package order

import (
	"context"
	grpcMock "github.com/baryshnikkov/rocket-factory/order/internal/client/grpc/mocks"
	"github.com/baryshnikkov/rocket-factory/order/internal/repository/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	ctx context.Context

	orderRepository *mocks.OrderRepository
	inventoryClient *grpcMock.InventoryClient
	paymentClient   *grpcMock.PaymentClient

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.orderRepository = mocks.NewOrderRepository(s.T())
	s.inventoryClient = grpcMock.NewInventoryClient(s.T())
	s.paymentClient = grpcMock.NewPaymentClient(s.T())

	s.service = NewService(
		s.orderRepository,
		s.inventoryClient,
		s.paymentClient,
	)
}

func (s *ServiceSuite) TearDownTest() {
	s.orderRepository.AssertExpectations(s.T())
	s.inventoryClient.AssertExpectations(s.T())
	s.paymentClient.AssertExpectations(s.T())
}

func TestServiceInventoryIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
