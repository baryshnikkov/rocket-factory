package part

import (
	"sync"

	def "github.com/baryshnikkov/rocket-factory/inventory/internal/repository"
	repoModel "github.com/baryshnikkov/rocket-factory/inventory/internal/repository/model"
)

var _ def.InventoryRepository = (*repository)(nil)

type repository struct {
	mu   sync.RWMutex
	data map[string]repoModel.Part
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]repoModel.Part),
	}
}
