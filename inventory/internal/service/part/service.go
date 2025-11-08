package part

import (
	"github.com/baryshnikkov/rocket-factory/inventory/internal/repository"
	def "github.com/baryshnikkov/rocket-factory/inventory/internal/service"
)

var _ def.InventoryService = (*service)(nil)

type service struct {
	inventoryRepository repository.InventoryRepository
}

func NewService(inventoryRepository repository.InventoryRepository) *service {
	return &service{
		inventoryRepository: inventoryRepository,
	}
}
