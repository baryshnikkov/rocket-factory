package order

import (
	"context"

	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	"github.com/baryshnikkov/rocket-factory/order/internal/repository/converter"
)

func (r *repository) GetOrder(_ context.Context, uuid string) (order model.OrderDto, err error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	outOrder, ok := r.data[uuid]
	if !ok {
		return model.OrderDto{}, model.ErrOrderNotFound
	}

	return converter.OrderDataToModel(outOrder), nil
}
