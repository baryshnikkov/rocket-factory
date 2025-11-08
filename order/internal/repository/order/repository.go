package order

import (
	"sync"

	def "github.com/baryshnikkov/rocket-factory/order/internal/repository"
	repoModel "github.com/baryshnikkov/rocket-factory/order/internal/repository/model"
)

var _ def.OrderRepository = (*repository)(nil)

type repository struct {
	mu   sync.RWMutex
	data map[string]repoModel.OrderDto
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]repoModel.OrderDto),
	}
}
