package order

import (
	"context"
	"time"

	"github.com/samber/lo"

	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	repoModel "github.com/baryshnikkov/rocket-factory/order/internal/repository/model"
)

func (r *repository) UpdateOrder(ctx context.Context, orderUUID string, orderUpdateInfo model.OrderUpdateInfo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, ok := r.data[orderUUID]
	if !ok {
		return model.ErrOrderNotFound
	}

	if orderUpdateInfo.TotalPrice != nil {
		order.TotalPrice = *orderUpdateInfo.TotalPrice
	}

	if orderUpdateInfo.TransactionUUID != nil {
		order.TransactionUUID = orderUpdateInfo.TransactionUUID
	}

	if orderUpdateInfo.PaymentMethod != nil {
		order.PaymentMethod = lo.ToPtr(repoModel.PaymentMethod(lo.FromPtr(orderUpdateInfo.PaymentMethod)))
	}

	if orderUpdateInfo.Status != nil {
		order.Status = repoModel.OrderStatus(lo.FromPtr(orderUpdateInfo.Status))
	}

	order.UpdatedAt = lo.ToPtr(time.Now())

	r.data[orderUUID] = order

	return nil
}
