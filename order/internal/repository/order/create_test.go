package order

import (
	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	repoModel "github.com/baryshnikkov/rocket-factory/order/internal/repository/model"
	"github.com/brianvoe/gofakeit/v7"
	"sync"
)

func (s *RepositorySuite) TestCreateOrderSuccess() {
	// Arrange
	userUUID := gofakeit.UUID()
	parts := []model.Part{
		{
			UUID:  gofakeit.UUID(),
			Name:  "Engine",
			Price: 1000.50,
		},
		{
			UUID:  gofakeit.UUID(),
			Name:  "Wing",
			Price: 500.25,
		},
	}

	expectedTotalPrice := 1000.50 + 500.25

	// Act
	res, err := s.repo.CreateOrder(s.ctx, userUUID, parts)

	// Assert
	s.Require().NoError(err)
	s.Require().NotEmpty(res.OrderUUID)
	s.Require().Equal(expectedTotalPrice, res.TotalPrice)

	// Проверяем что заказ сохранился в репозитории
	savedOrder, exists := s.repo.data[res.OrderUUID]
	s.Require().True(exists)
	s.Require().Equal(res.OrderUUID, savedOrder.UUID)
	s.Require().Equal(userUUID, savedOrder.UserUUID)
	s.Require().Equal(expectedTotalPrice, savedOrder.TotalPrice)
	s.Require().Equal(repoModel.OrderStatusPendingPayment, savedOrder.Status)
	s.Require().Len(savedOrder.PartsUUIDs, 2)
	s.Require().Contains(savedOrder.PartsUUIDs, parts[0].UUID)
	s.Require().Contains(savedOrder.PartsUUIDs, parts[1].UUID)
}

func (s *RepositorySuite) TestCreateOrderEmptyParts() {
	// Arrange
	userUUID := gofakeit.UUID()
	var parts []model.Part

	// Act
	res, err := s.repo.CreateOrder(s.ctx, userUUID, parts)

	// Assert
	s.Require().NoError(err)
	s.Require().NotEmpty(res.OrderUUID)
	s.Require().Equal(0.0, res.TotalPrice)

	savedOrder, exists := s.repo.data[res.OrderUUID]
	s.Require().True(exists)
	s.Require().Equal(res.OrderUUID, savedOrder.UUID)
	s.Require().Equal(userUUID, savedOrder.UserUUID)
	s.Require().Equal(0.0, savedOrder.TotalPrice)
	s.Require().Empty(savedOrder.PartsUUIDs)
}

func (s *RepositorySuite) TestCreateOrderNilContext() {
	// Arrange
	userUUID := gofakeit.UUID()
	parts := []model.Part{
		{
			UUID:  gofakeit.UUID(),
			Name:  "Engine",
			Price: 1000.0,
		},
	}

	// Act - передаем nil контекст
	res, err := s.repo.CreateOrder(nil, userUUID, parts)

	// Assert - должен работать даже с nil контекстом
	s.Require().NoError(err)
	s.Require().NotEmpty(res.OrderUUID)
	s.Require().Equal(1000.0, res.TotalPrice)

	savedOrder, exists := s.repo.data[res.OrderUUID]
	s.Require().True(exists)
	s.Require().Equal(res.OrderUUID, savedOrder.UUID)
}

func (s *RepositorySuite) TestCreateOrderEmptyUserUUID() {
	// Arrange
	userUUID := ""
	parts := []model.Part{
		{
			UUID:  gofakeit.UUID(),
			Name:  "Engine",
			Price: 1000.0,
		},
	}

	// Act
	res, err := s.repo.CreateOrder(s.ctx, userUUID, parts)

	// Assert
	s.Require().NoError(err)
	s.Require().NotEmpty(res.OrderUUID)
	s.Require().Equal(1000.0, res.TotalPrice)

	savedOrder, exists := s.repo.data[res.OrderUUID]
	s.Require().True(exists)
	s.Require().Equal("", savedOrder.UserUUID)
}

func (s *RepositorySuite) TestCreateOrderConcurrent() {
	// Arrange
	userUUID := gofakeit.UUID()
	parts := []model.Part{
		{
			UUID:  gofakeit.UUID(),
			Name:  "Engine",
			Price: 1000.0,
		},
	}

	// Act - запускаем несколько горутин для создания заказов
	var wg sync.WaitGroup
	orderUUIDs := make([]string, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			res, err := s.repo.CreateOrder(s.ctx, userUUID, parts)
			s.Require().NoError(err)
			orderUUIDs[index] = res.OrderUUID
		}(i)
	}
	wg.Wait()

	// Assert - проверяем что все заказы создались с уникальными UUID
	s.Require().Len(s.repo.data, 10)

	uniqueUUIDs := make(map[string]bool)
	for _, uuid := range orderUUIDs {
		uniqueUUIDs[uuid] = true
		_, exists := s.repo.data[uuid]
		s.Require().True(exists)
	}
	s.Require().Len(uniqueUUIDs, 10) // Все UUID должны быть уникальными
}
