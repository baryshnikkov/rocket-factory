package part

import (
	"context"

	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
	repoConverter "github.com/baryshnikkov/rocket-factory/inventory/internal/repository/converter"
)

func (r *repository) GetPart(_ context.Context, orderUUID string) (model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.data[orderUUID]
	if !ok {
		return model.Part{}, model.ErrPartNotFound
	}

	return repoConverter.PartToModel(part), nil
}
