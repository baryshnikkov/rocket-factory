package order

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	repoModel "github.com/baryshnikkov/rocket-factory/order/internal/repository/model"
)

func (r *repository) CreateOrder(ctx context.Context, userUUID string, parts []model.Part) (info model.OrderCreationInfo, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	orderUUID := uuid.NewString()

	var partUUIDs []string
	var totalPrice float64
	for _, part := range parts {
		partUUIDs = append(partUUIDs, part.UUID)
		totalPrice += part.Price
	}

	order := repoModel.OrderDto{
		UUID:       orderUUID,
		UserUUID:   userUUID,
		PartsUUIDs: partUUIDs,
		TotalPrice: totalPrice,
		Status:     repoModel.OrderStatusPendingPayment,
		CreatedAt:  time.Now(),
	}

	r.data[orderUUID] = order

	log.Printf(`
💳 [Order Created]
• 🆔 Order UUID: %s
• 👤 User UUID: %s
• 💰 Part UUIDs: %v
• 💰 Total Price: %f
• 💰 Status: %s
• 💰 CreatedAt: %v
`, order.UUID, order.UserUUID, order.PartsUUIDs, order.TotalPrice, order.Status, order.CreatedAt,
	)

	return model.OrderCreationInfo{
		OrderUUID:  orderUUID,
		TotalPrice: totalPrice,
	}, nil
}
