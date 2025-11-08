package v1

import (
	"context"

	paymentV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/proto/payment/v1"
)

func (c *client) PayOrder(ctx context.Context, userUUID, orderUUID, paymentMethod string) (transactionUUID string, err error) {
	res, err := c.generatedClient.PayOrder(ctx, &paymentV1.PayOrderRequest{
		OrderUuid:     orderUUID,
		UserUuid:      userUUID,
		PaymentMethod: paymentV1.PaymentMethod(paymentV1.PaymentMethod_value[paymentMethod]),
	})
	if err != nil {
		return "", err
	}
	return res.TransactionUuid, nil
}
