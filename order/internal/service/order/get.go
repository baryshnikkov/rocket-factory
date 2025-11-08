package order

import (
	"context"

	"github.com/baryshnikkov/rocket-factory/order/internal/model"
)

func (s *service) GetOrder(ctx context.Context, orderUUID string) (order model.OrderDto, err error) {
	outOrder, err := s.orderRepository.GetOrder(ctx, orderUUID)
	if err != nil {
		return model.OrderDto{}, err
	}

	return outOrder, nil
}
