package order

import (
	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	"github.com/brianvoe/gofakeit/v7"
	"time"
)

func (s *ServiceSuite) TestGetOrderSuccess() {
	order := model.OrderDto{
		UUID:       gofakeit.UUID(),
		UserUUID:   gofakeit.UUID(),
		PartsUUIDs: []string{gofakeit.UUID(), gofakeit.UUID()},
		TotalPrice: 1500.75,
		Status:     model.OrderStatusPaid,
		CreatedAt:  time.Now(),
	}

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)

	res, err := s.service.GetOrder(s.ctx, order.UUID)

	s.NoError(err)
	s.Equal(order, res)
}

func (s *ServiceSuite) TestGetOrderNotFound() {
	uuid := gofakeit.UUID()
	expectedErr := model.ErrOrderNotFound

	s.orderRepository.On("GetOrder", s.ctx, uuid).Return(model.OrderDto{}, expectedErr)

	res, err := s.service.GetOrder(s.ctx, uuid)

	s.Error(err)
	s.ErrorIs(err, expectedErr)
	s.Equal(model.OrderDto{}, res)
}

func (s *ServiceSuite) TestGetOrderRepositoryError() {
	uuid := gofakeit.UUID()
	expectedErr := model.ErrOrderInternalError

	s.orderRepository.On("GetOrder", s.ctx, uuid).Return(model.OrderDto{}, expectedErr)

	res, err := s.service.GetOrder(s.ctx, uuid)

	s.Error(err)
	s.ErrorIs(err, expectedErr)
	s.Equal(model.OrderDto{}, res)
}

func (s *ServiceSuite) TestGetOrderEmptyUUID() {
	uuid := ""
	expectedErr := model.ErrOrderNotFound

	s.orderRepository.On("GetOrder", s.ctx, uuid).Return(model.OrderDto{}, expectedErr)

	res, err := s.service.GetOrder(s.ctx, uuid)

	s.Error(err)
	s.ErrorIs(err, expectedErr)
	s.Equal(model.OrderDto{}, res)
}
