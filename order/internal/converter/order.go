package converter

import (
	"log"

	"github.com/google/uuid"

	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	orderV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/openapi/order/v1"
)

func StringToUUID(s string) uuid.UUID {
	u, err := uuid.Parse(s)
	if err != nil {
		log.Printf("Failed to parse UUID: %v", err)
	}

	return u
}

func UUIDsToStrings(arr []uuid.UUID) []string {
	uuids := make([]string, len(arr))
	for i, s := range arr {
		uuids[i] = s.String()
	}
	return uuids
}

func OrderDataToDTO(order model.OrderDto) orderV1.OrderDto {
	var transactionUUID orderV1.OptUUID
	if order.TransactionUUID != nil {
		transactionUUID = orderV1.OptUUID{Value: StringToUUID(*order.TransactionUUID)}
	}

	var paymentMethod orderV1.OptPaymentMethod
	if order.PaymentMethod != nil {
		paymentMethod = orderV1.OptPaymentMethod{Value: paymentMethodToOpt(*order.PaymentMethod)}
	}

	var updatedAt orderV1.OptDateTime
	if order.UpdatedAt != nil {
		updatedAt = orderV1.OptDateTime{Value: *order.UpdatedAt}
	}

	createdAt := orderV1.OptDateTime{Value: order.CreatedAt}

	return orderV1.OrderDto{
		OrderUUID:       StringToUUID(order.UUID),
		UserUUID:        StringToUUID(order.UserUUID),
		PartUuids:       stringsToUUIDs(order.PartsUUIDs),
		TotalPrice:      order.TotalPrice,
		TransactionUUID: transactionUUID,
		PaymentMethod:   paymentMethod,
		Status:          orderV1.OrderStatus(order.Status),
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}

func paymentMethodToOpt(paymentMethod model.PaymentMethod) orderV1.PaymentMethod {
	return orderV1.PaymentMethod(paymentMethod)
}

func stringsToUUIDs(arr []string) []uuid.UUID {
	uuids := make([]uuid.UUID, len(arr))
	for i, s := range arr {
		uuids[i] = StringToUUID(s)
	}
	return uuids
}
