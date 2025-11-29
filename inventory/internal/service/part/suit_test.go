package part

import (
	"context"
	"github.com/baryshnikkov/rocket-factory/inventory/internal/repository/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	ctx context.Context

	inventoryRepository *mocks.InventoryRepository

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.inventoryRepository = mocks.NewInventoryRepository(s.T())

	s.service = NewService(
		s.inventoryRepository,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceInventoryIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
