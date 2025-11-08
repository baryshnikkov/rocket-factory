package converter

import (
	"github.com/samber/lo"

	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	repoModel "github.com/baryshnikkov/rocket-factory/order/internal/repository/model"
)

func OrderDataToModel(order repoModel.OrderDto) model.OrderDto {
	return model.OrderDto{
		UUID:            order.UUID,
		UserUUID:        order.UserUUID,
		PartsUUIDs:      order.PartsUUIDs,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   lo.ToPtr(model.PaymentMethod(lo.FromPtr(order.PaymentMethod))),
		Status:          model.OrderStatus(order.Status),
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}
}
