package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/baryshnikkov/rocket-factory/inventory/internal/service/mocks"
)

type APISuite struct {
	suite.Suite
	ctx context.Context

	inventoryService *mocks.InventoryService
	api              *api
}

func (s *APISuite) SetupTest() {
	s.ctx = context.Background()
	s.inventoryService = mocks.NewInventoryService(s.T())
	s.api = NewAPI(s.inventoryService)
}

func (s *APISuite) TearDownSuite() {}

func TestAPIInventoryIntegration(t *testing.T) {
	suite.Run(t, new(APISuite))
}
