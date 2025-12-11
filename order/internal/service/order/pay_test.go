package order

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"

	"github.com/baryshnikkov/rocket-factory/order/internal/model"
)

func (s *ServiceSuite) TestPayOrderSuccess() {
	order := model.OrderDto{
		UUID:     gofakeit.UUID(),
		UserUUID: gofakeit.UUID(),
		Status:   model.OrderStatusPendingPayment,
	}
	paymentMethod := "CARD"
	transUUID := gofakeit.UUID()

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)
	s.paymentClient.On("PayOrder", s.ctx, order.UserUUID, order.UUID, paymentMethod).Return(transUUID, nil)
	s.orderRepository.On("UpdateOrder", s.ctx, order.UUID, model.OrderUpdateInfo{
		Status:          lo.ToPtr(model.OrderStatusPaid),
		PaymentMethod:   lo.ToPtr(model.PaymentMethod(paymentMethod)),
		TransactionUUID: lo.ToPtr(transUUID),
	}).Return(nil)

	result, err := s.service.PayOrder(s.ctx, order.UUID, paymentMethod)

	s.NoError(err)
	s.Equal(transUUID, result)
}

func (s *ServiceSuite) TestPayOrderAlreadyPaid() {
	order := model.OrderDto{
		UUID:   gofakeit.UUID(),
		Status: model.OrderStatusPaid,
	}
	paymentMethod := "CARD"

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)

	result, err := s.service.PayOrder(s.ctx, order.UUID, paymentMethod)

	s.Error(err)
	s.ErrorIs(err, model.ErrPaymentConflict)
	s.Equal("", result)
}

func (s *ServiceSuite) TestPayOrderAlreadyCancelled() {
	order := model.OrderDto{
		UUID:   gofakeit.UUID(),
		Status: model.OrderStatusCancelled,
	}
	paymentMethod := "SBP"

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)

	result, err := s.service.PayOrder(s.ctx, order.UUID, paymentMethod)

	s.Error(err)
	s.ErrorIs(err, model.ErrPaymentConflict)
	s.Equal("", result)
}

func (s *ServiceSuite) TestPayOrderGetOrderNotFound() {
	uuid := gofakeit.UUID()
	paymentMethod := "CARD"
	expectedErr := model.ErrOrderNotFound

	s.orderRepository.On("GetOrder", s.ctx, uuid).Return(model.OrderDto{}, expectedErr)

	result, err := s.service.PayOrder(s.ctx, uuid, paymentMethod)

	s.Error(err)
	s.ErrorIs(err, expectedErr)
	s.Equal("", result)
}

func (s *ServiceSuite) TestPayOrderPaymentClientError() {
	order := model.OrderDto{
		UUID:     gofakeit.UUID(),
		UserUUID: gofakeit.UUID(),
		Status:   model.OrderStatusPendingPayment,
	}
	paymentMethod := "CREDIT_CARD"
	expectedErr := model.ErrPaymentInternalError

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)
	s.paymentClient.On("PayOrder", s.ctx, order.UserUUID, order.UUID, paymentMethod).Return("", expectedErr)

	result, err := s.service.PayOrder(s.ctx, order.UUID, paymentMethod)

	s.Error(err)
	s.ErrorIs(err, expectedErr)
	s.Equal("", result)
}

func (s *ServiceSuite) TestPayOrderUpdateOrderError() {
	order := model.OrderDto{
		UUID:     gofakeit.UUID(),
		UserUUID: gofakeit.UUID(),
		Status:   model.OrderStatusPendingPayment,
	}
	paymentMethod := "INVESTOR_MONEY"
	transUUID := gofakeit.UUID()
	expectedErr := model.ErrOrderInternalError

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)
	s.paymentClient.On("PayOrder", s.ctx, order.UserUUID, order.UUID, paymentMethod).Return(transUUID, nil)
	s.orderRepository.On("UpdateOrder", s.ctx, order.UUID, model.OrderUpdateInfo{
		Status:          lo.ToPtr(model.OrderStatusPaid),
		PaymentMethod:   lo.ToPtr(model.PaymentMethod(paymentMethod)),
		TransactionUUID: lo.ToPtr(transUUID),
	}).Return(expectedErr)

	result, err := s.service.PayOrder(s.ctx, order.UUID, paymentMethod)

	s.Error(err)
	s.ErrorIs(err, expectedErr)
	s.Equal("", result)
}

func (s *ServiceSuite) TestPayOrderUnknownStatus() {
	order := model.OrderDto{
		UUID:   gofakeit.UUID(),
		Status: model.OrderStatus("UNKNOWN_STATUS"),
	}
	paymentMethod := "CARD"

	s.orderRepository.On("GetOrder", s.ctx, order.UUID).Return(order, nil)

	result, err := s.service.PayOrder(s.ctx, order.UUID, paymentMethod)

	s.Error(err)
	s.ErrorIs(err, model.ErrPaymentInternalError)
	s.Equal("", result)
}

func (s *ServiceSuite) TestCanPayOrderPendingPayment() {
	order := model.OrderDto{
		Status: model.OrderStatusPendingPayment,
	}

	err, shouldReturn := canPayOrder(order)

	s.NoError(err)
	s.False(shouldReturn)
}

func (s *ServiceSuite) TestCanPayOrderAlreadyPaid() {
	order := model.OrderDto{
		Status: model.OrderStatusPaid,
	}

	err, shouldReturn := canPayOrder(order)

	s.Error(err)
	s.ErrorIs(err, model.ErrPaymentConflict)
	s.True(shouldReturn)
}

func (s *ServiceSuite) TestCanPayOrderAlreadyCancelled() {
	order := model.OrderDto{
		Status: model.OrderStatusCancelled,
	}

	err, shouldReturn := canPayOrder(order)

	s.Error(err)
	s.ErrorIs(err, model.ErrPaymentConflict)
	s.True(shouldReturn)
}

func (s *ServiceSuite) TestCanPayOrderUnknownStatus() {
	order := model.OrderDto{
		Status: model.OrderStatus("UNKNOWN"),
	}

	err, shouldReturn := canPayOrder(order)

	s.Error(err)
	s.ErrorIs(err, model.ErrPaymentInternalError)
	s.True(shouldReturn)
}
