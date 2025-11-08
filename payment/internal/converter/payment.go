package converter

import (
	"github.com/baryshnikkov/rocket-factory/payment/internal/model"
	paymentV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/proto/payment/v1"
)

func PayOrderRequestToModel(payOrder *paymentV1.PayOrderRequest) model.PayOrderRequest {
	return model.PayOrderRequest{
		OrderUUID:     payOrder.GetOrderUuid(),
		UserUUID:      payOrder.GetUserUuid(),
		PaymentMethod: paymentMethodToModel(payOrder.GetPaymentMethod()),
	}
}

func PayOrderResponseToProto(payOrder model.PayOrderResponse) *paymentV1.PayOrderResponse {
	return &paymentV1.PayOrderResponse{
		TransactionUuid: payOrder.TransactionUUID,
	}
}

func paymentMethodToModel(payMethod paymentV1.PaymentMethod) model.PaymentMethod {
	switch payMethod {
	case paymentV1.PaymentMethod_PAYMENT_METHOD_CARD:
		return model.Card
	case paymentV1.PaymentMethod_PAYMENT_METHOD_SBP:
		return model.SBP
	case paymentV1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD:
		return model.CreditCard
	case paymentV1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY:
		return model.InvestorMoney
	default:
		return model.Unspecified
	}
}
