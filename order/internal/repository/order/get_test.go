package order

import (
	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	repoModel "github.com/baryshnikkov/rocket-factory/order/internal/repository/model"
	"github.com/brianvoe/gofakeit/v7"
	"time"
)

func (s *RepositorySuite) TestGetOrderSuccess() {
	// Arrange
	repoOrder := repoModel.OrderDto{
		UUID:       gofakeit.UUID(),
		UserUUID:   gofakeit.UUID(),
		PartsUUIDs: []string{gofakeit.UUID(), gofakeit.UUID()},
		TotalPrice: 1500.75,
		Status:     repoModel.OrderStatusPendingPayment,
		CreatedAt:  time.Now(),
	}
	s.repo.data[repoOrder.UUID] = repoOrder

	// Act
	res, err := s.repo.GetOrder(s.ctx, repoOrder.UUID)

	// Assert
	s.Require().NoError(err)
	s.Require().Equal(repoOrder.UUID, res.UUID)
	s.Require().Equal(repoOrder.UserUUID, res.UserUUID)
	s.Require().Equal(repoOrder.PartsUUIDs, res.PartsUUIDs)
	s.Require().Equal(repoOrder.TotalPrice, res.TotalPrice)
	s.Require().Equal(model.OrderStatusPendingPayment, res.Status)
}

func (s *RepositorySuite) TestGetOrderNotFound() {
	// Arrange - НЕ добавляем данные в repository
	uuid := gofakeit.UUID()

	// Act - пытаемся получить несуществующий заказ
	res, err := s.repo.GetOrder(s.ctx, uuid)

	// Assert - проверяем что получили ошибку "не найдено"
	s.Require().Error(err)
	s.Require().Equal(model.ErrOrderNotFound, err)
	s.Require().Equal(model.OrderDto{}, res)
}

func (s *RepositorySuite) TestGetOrderEmptyUUID() {
	// Act
	res, err := s.repo.GetOrder(s.ctx, "")

	// Assert
	s.Require().Error(err)
	s.Require().Equal(model.ErrOrderNotFound, err)
	s.Require().Equal(model.OrderDto{}, res)
}

func (s *RepositorySuite) TestGetOrderNilContext() {
	// Arrange
	repoOrder := repoModel.OrderDto{
		UUID:       gofakeit.UUID(),
		UserUUID:   gofakeit.UUID(),
		PartsUUIDs: []string{gofakeit.UUID()},
		TotalPrice: 1000.0,
		Status:     repoModel.OrderStatusPaid,
		CreatedAt:  time.Now(),
	}
	s.repo.data[repoOrder.UUID] = repoOrder

	// Act - передаем nil контекст
	res, err := s.repo.GetOrder(nil, repoOrder.UUID)

	// Assert - должен работать даже с nil контекстом
	s.Require().NoError(err)
	s.Require().Equal(repoOrder.UUID, res.UUID)
	s.Require().Equal(model.OrderStatusPaid, res.Status)
}
