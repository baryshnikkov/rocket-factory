package service

import (
	"context"

	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
)

type InventoryService interface {
	GetPart(ctx context.Context, orderUUID string) (model.Part, error)
	ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error)
}
