package payment

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"

	"github.com/baryshnikkov/rocket-factory/payment/internal/model"
)

func (s *ServiceSuite) TestPayOrderSuccess() {
	req := model.PayOrderRequest{
		OrderUUID:     gofakeit.UUID(),
		UserUUID:      gofakeit.UUID(),
		PaymentMethod: model.Card,
	}

	res, err := s.service.PayOrder(s.ctx, req)

	s.NoError(err)
	s.NotEmpty(res.TransactionUUID)

	// Проверяем что UUID валидный
	_, err = uuid.Parse(res.TransactionUUID)
	s.NoError(err)
}

func (s *ServiceSuite) TestPayOrderSBP() {
	req := model.PayOrderRequest{
		OrderUUID:     gofakeit.UUID(),
		UserUUID:      gofakeit.UUID(),
		PaymentMethod: model.SBP,
	}

	res, err := s.service.PayOrder(s.ctx, req)

	s.NoError(err)
	s.NotEmpty(res.TransactionUUID)
}

func (s *ServiceSuite) TestPayOrderCreditCard() {
	req := model.PayOrderRequest{
		OrderUUID:     gofakeit.UUID(),
		UserUUID:      gofakeit.UUID(),
		PaymentMethod: model.CreditCard,
	}

	res, err := s.service.PayOrder(s.ctx, req)

	s.NoError(err)
	s.NotEmpty(res.TransactionUUID)
}

func (s *ServiceSuite) TestPayOrderInvestorMoney() {
	req := model.PayOrderRequest{
		OrderUUID:     gofakeit.UUID(),
		UserUUID:      gofakeit.UUID(),
		PaymentMethod: model.InvestorMoney,
	}

	res, err := s.service.PayOrder(s.ctx, req)

	s.NoError(err)
	s.NotEmpty(res.TransactionUUID)
}

func (s *ServiceSuite) TestPayOrderUnspecifiedPaymentMethod() {
	req := model.PayOrderRequest{
		OrderUUID:     gofakeit.UUID(),
		UserUUID:      gofakeit.UUID(),
		PaymentMethod: model.Unspecified,
	}

	res, err := s.service.PayOrder(s.ctx, req)

	s.NoError(err)
	s.NotEmpty(res.TransactionUUID)
}

func (s *ServiceSuite) TestPayOrderEmptyUUIDs() {
	req := model.PayOrderRequest{
		OrderUUID:     "",
		UserUUID:      "",
		PaymentMethod: model.Card,
	}

	res, err := s.service.PayOrder(s.ctx, req)

	s.NoError(err)
	s.NotEmpty(res.TransactionUUID)
}

func (s *ServiceSuite) TestPayOrderNilContext() {
	req := model.PayOrderRequest{
		OrderUUID:     gofakeit.UUID(),
		UserUUID:      gofakeit.UUID(),
		PaymentMethod: model.Card,
	}

	res, err := s.service.PayOrder(nil, req)
	s.NoError(err)
	s.NotEmpty(res.TransactionUUID)
}
