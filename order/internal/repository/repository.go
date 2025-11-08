package repository

import (
	"context"

	"github.com/baryshnikkov/rocket-factory/order/internal/model"
)

type OrderRepository interface {
	GetOrder(ctx context.Context, UUID string) (order model.OrderDto, err error)
	CreateOrder(ctx context.Context, userUUID string, parts []model.Part) (info model.OrderCreationInfo, err error)
	UpdateOrder(ctx context.Context, orderUUID string, orderUpdateInfo model.OrderUpdateInfo) error
}
