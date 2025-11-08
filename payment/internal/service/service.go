package service

import (
	"context"

	"github.com/baryshnikkov/rocket-factory/payment/internal/model"
)

type PaymentService interface {
	PayOrder(ctx context.Context, req model.PayOrderRequest) (model.PayOrderResponse, error)
}
