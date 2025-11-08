package order

import (
	"context"

	"github.com/baryshnikkov/rocket-factory/order/internal/model"
)

func (s *service) CancelOrder(ctx context.Context, userUUID string) error {
	order, err := s.orderRepository.GetOrder(ctx, userUUID)
	if err != nil {
		return err
	}

	switch order.Status {
	case model.OrderStatusPaid:
		return model.ErrOrderConflict
	case model.OrderStatusCancelled:
		return model.ErrOrderConflict
	case model.OrderStatusPendingPayment:
		status := model.OrderStatusCancelled
		err = s.orderRepository.UpdateOrder(ctx, order.UUID, model.OrderUpdateInfo{
			Status: &status,
		})
		if err != nil {
			return err
		}
		return nil
	default:
		return model.ErrOrderInternalError
	}
}
