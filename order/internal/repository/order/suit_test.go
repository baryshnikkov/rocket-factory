package order

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"

	repoModel "github.com/baryshnikkov/rocket-factory/order/internal/repository/model"
)

type RepositorySuite struct {
	suite.Suite
	ctx context.Context

	repo *repository
}

func (s *RepositorySuite) SetupTest() {
	s.ctx = context.Background()
	s.repo = &repository{
		data: make(map[string]repoModel.OrderDto), // ← реальные данные
		mu:   sync.RWMutex{},
	}
}

func (s *RepositorySuite) TearDownSuite() {}

func TestRepositoryOrderIntegration(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
