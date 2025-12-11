package order

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"

	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	repoModel "github.com/baryshnikkov/rocket-factory/order/internal/repository/model"
)

func (s *RepositorySuite) TestUpdateOrderSuccess() {
	// Arrange
	repoOrder := repoModel.OrderDto{
		UUID:       gofakeit.UUID(),
		UserUUID:   gofakeit.UUID(),
		PartsUUIDs: []string{gofakeit.UUID()},
		TotalPrice: 1000.0,
		Status:     repoModel.OrderStatusPendingPayment,
		CreatedAt:  time.Now(),
	}
	s.repo.data[repoOrder.UUID] = repoOrder

	newTotalPrice := 1500.0
	newTransactionUUID := gofakeit.UUID()
	newPaymentMethod := model.Card
	newStatus := model.OrderStatusPaid

	updateInfo := model.OrderUpdateInfo{
		TotalPrice:      &newTotalPrice,
		TransactionUUID: &newTransactionUUID,
		PaymentMethod:   &newPaymentMethod,
		Status:          &newStatus,
	}

	// Act
	err := s.repo.UpdateOrder(s.ctx, repoOrder.UUID, updateInfo)

	// Assert
	s.Require().NoError(err)

	updatedOrder := s.repo.data[repoOrder.UUID]
	s.Require().Equal(newTotalPrice, updatedOrder.TotalPrice)
	s.Require().Equal(newTransactionUUID, *updatedOrder.TransactionUUID)
	s.Require().Equal(repoModel.PaymentMethod(newPaymentMethod), *updatedOrder.PaymentMethod)
	s.Require().Equal(repoModel.OrderStatus(newStatus), updatedOrder.Status)
	s.Require().NotNil(updatedOrder.UpdatedAt)
}

func (s *RepositorySuite) TestUpdateOrderPartialFields() {
	// Arrange
	repoOrder := repoModel.OrderDto{
		UUID:       gofakeit.UUID(),
		UserUUID:   gofakeit.UUID(),
		PartsUUIDs: []string{gofakeit.UUID()},
		TotalPrice: 1000.0,
		Status:     repoModel.OrderStatusPendingPayment,
		CreatedAt:  time.Now(),
	}
	s.repo.data[repoOrder.UUID] = repoOrder

	newStatus := model.OrderStatusCancelled
	updateInfo := model.OrderUpdateInfo{
		Status: &newStatus,
	}

	// Act
	err := s.repo.UpdateOrder(s.ctx, repoOrder.UUID, updateInfo)

	// Assert
	s.Require().NoError(err)

	updatedOrder := s.repo.data[repoOrder.UUID]
	s.Require().Equal(repoModel.OrderStatus(newStatus), updatedOrder.Status)
	s.Require().Equal(1000.0, updatedOrder.TotalPrice) // осталось прежним
	s.Require().NotNil(updatedOrder.UpdatedAt)
}

func (s *RepositorySuite) TestUpdateOrderNotFound() {
	// Arrange
	uuid := gofakeit.UUID()
	updateInfo := model.OrderUpdateInfo{
		Status: lo.ToPtr(model.OrderStatusPaid),
	}

	// Act
	err := s.repo.UpdateOrder(s.ctx, uuid, updateInfo)

	// Assert
	s.Require().Error(err)
	s.Require().Equal(model.ErrOrderNotFound, err)
}

func (s *RepositorySuite) TestUpdateOrderEmptyUUID() {
	// Arrange
	updateInfo := model.OrderUpdateInfo{
		Status: lo.ToPtr(model.OrderStatusPaid),
	}

	// Act
	err := s.repo.UpdateOrder(s.ctx, "", updateInfo)

	// Assert
	s.Require().Error(err)
	s.Require().Equal(model.ErrOrderNotFound, err)
}

func (s *RepositorySuite) TestUpdateOrderNilContext() {
	// Arrange
	repoOrder := repoModel.OrderDto{
		UUID:       gofakeit.UUID(),
		UserUUID:   gofakeit.UUID(),
		PartsUUIDs: []string{gofakeit.UUID()},
		TotalPrice: 1000.0,
		Status:     repoModel.OrderStatusPendingPayment,
		CreatedAt:  time.Now(),
	}
	s.repo.data[repoOrder.UUID] = repoOrder

	newStatus := model.OrderStatusPaid
	updateInfo := model.OrderUpdateInfo{
		Status: &newStatus,
	}

	// Act - передаем nil контекст
	err := s.repo.UpdateOrder(nil, repoOrder.UUID, updateInfo)

	// Assert - должен работать даже с nil контекстом
	s.Require().NoError(err)

	updatedOrder := s.repo.data[repoOrder.UUID]
	s.Require().Equal(repoModel.OrderStatus(newStatus), updatedOrder.Status)
	s.Require().NotNil(updatedOrder.UpdatedAt)
}

func (s *RepositorySuite) TestUpdateOrderEmptyUpdateInfo() {
	// Arrange
	repoOrder := repoModel.OrderDto{
		UUID:       gofakeit.UUID(),
		UserUUID:   gofakeit.UUID(),
		PartsUUIDs: []string{gofakeit.UUID()},
		TotalPrice: 1000.0,
		Status:     repoModel.OrderStatusPendingPayment,
		CreatedAt:  time.Now(),
	}
	s.repo.data[repoOrder.UUID] = repoOrder

	updateInfo := model.OrderUpdateInfo{}

	// Act
	err := s.repo.UpdateOrder(s.ctx, repoOrder.UUID, updateInfo)

	// Assert
	s.Require().NoError(err)

	updatedOrder := s.repo.data[repoOrder.UUID]
	s.Require().Equal(1000.0, updatedOrder.TotalPrice)                          // не изменилось
	s.Require().Equal(repoModel.OrderStatusPendingPayment, updatedOrder.Status) // не изменилось
	s.Require().NotNil(updatedOrder.UpdatedAt)                                  // UpdatedAt все равно обновился
}
