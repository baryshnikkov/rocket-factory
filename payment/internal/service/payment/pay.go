package payment

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/baryshnikkov/rocket-factory/payment/internal/model"
)

func (s *service) PayOrder(ctx context.Context, req model.PayOrderRequest) (model.PayOrderResponse, error) {
	log.Printf(`
💳 [Order Paid]
• 🆔 Order UUID: %s
• 👤 User UUID: %s
• 💰 Payment Method: %s
`, req.OrderUUID, req.UserUUID, req.PaymentMethod,
	)

	//  if err != nil {
	//	  return model.PayOrderResponse{}, model.ErrPaymentInternalError
	//  }

	UUID := uuid.New().String()
	log.Printf("✅Оплата прошла успешно, transaction_uuid: %v\n", UUID)

	return model.PayOrderResponse{
		TransactionUUID: UUID,
	}, nil
}
