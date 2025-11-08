package part

import (
	"context"

	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
)

func (s *service) GetPart(ctx context.Context, orderUUID string) (model.Part, error) {
	part, err := s.inventoryRepository.GetPart(ctx, orderUUID)
	if err != nil {
		return model.Part{}, err
	}

	return part, nil
}
