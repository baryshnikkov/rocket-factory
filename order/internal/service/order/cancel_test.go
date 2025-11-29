package order

import (
	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"
)

func (s *ServiceSuite) TestCancelOrderSuccess() {
	order := model.OrderDto{
		UUID:   gofakeit.UUID(),
		Status: model.OrderStatusPendingPayment,
	}

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)
	s.orderRepository.On("UpdateOrder", s.ctx, order.UUID, model.OrderUpdateInfo{
		Status: lo.ToPtr(model.OrderStatusCancelled),
	}).Return(nil)

	err := s.service.CancelOrder(s.ctx, order.UUID)

	s.NoError(err)
}

func (s *ServiceSuite) TestCancelOrderAlreadyPaid() {
	order := model.OrderDto{
		UUID:   gofakeit.UUID(),
		Status: model.OrderStatusPaid,
	}

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)

	err := s.service.CancelOrder(s.ctx, order.UUID)

	s.Error(err)
	s.ErrorIs(err, model.ErrOrderConflict)
}

func (s *ServiceSuite) TestCancelOrderAlreadyCancelled() {
	order := model.OrderDto{
		UUID:   gofakeit.UUID(),
		Status: model.OrderStatusCancelled,
	}

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)

	err := s.service.CancelOrder(s.ctx, order.UUID)

	s.Error(err)
	s.ErrorIs(err, model.ErrOrderConflict)
}

func (s *ServiceSuite) TestCancelOrderGetOrderNotFound() {
	uuid := gofakeit.UUID()
	expectedErr := model.ErrOrderNotFound

	s.orderRepository.On("GetOrder", s.ctx, uuid).Return(model.OrderDto{}, expectedErr)

	err := s.service.CancelOrder(s.ctx, uuid)

	s.Error(err)
	s.ErrorIs(err, expectedErr)
}

func (s *ServiceSuite) TestCancelOrderUpdateOrderError() {
	order := model.OrderDto{
		UUID:   gofakeit.UUID(),
		Status: model.OrderStatusPendingPayment,
	}
	expectedErr := model.ErrOrderInternalError

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)
	s.orderRepository.On("UpdateOrder", s.ctx, order.UUID, model.OrderUpdateInfo{
		Status: lo.ToPtr(model.OrderStatusCancelled),
	}).Return(expectedErr)

	err := s.service.CancelOrder(s.ctx, order.UUID)

	s.Error(err)
	s.ErrorIs(err, expectedErr)
}

func (s *ServiceSuite) TestCancelOrderUnknownStatus() {
	order := model.OrderDto{
		UUID:   gofakeit.UUID(),
		Status: model.OrderStatus("UNKNOWN_STATUS"),
	}

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)

	err := s.service.CancelOrder(s.ctx, order.UUID)

	s.Error(err)
	s.ErrorIs(err, model.ErrOrderInternalError)
}
