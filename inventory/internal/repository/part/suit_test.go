package part

import (
	"context"
	"sync"
	"testing"

	repoModel "github.com/baryshnikkov/rocket-factory/inventory/internal/repository/model"
	"github.com/stretchr/testify/suite"
)

type RepositorySuite struct {
	suite.Suite
	ctx context.Context

	repo *repository
}

func (s *RepositorySuite) SetupTest() {
	s.ctx = context.Background()
	s.repo = &repository{
		data: make(map[string]repoModel.Part), // ← реальные данные
		mu:   sync.RWMutex{},
	}
}

func (s *RepositorySuite) TearDownSuite() {}

func TestRepositoryInventoryIntegration(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
