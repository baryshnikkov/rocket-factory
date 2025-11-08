package repository

import (
	"context"

	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
)

type InventoryRepository interface {
	GetPart(ctx context.Context, UUID string) (model.Part, error)
	ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error)
	InitParts()
}
