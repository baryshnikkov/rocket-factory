package v1

import (
	"errors"
	"github.com/baryshnikkov/rocket-factory/order/internal/converter"
	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	orderV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/openapi/order/v1"
	"github.com/google/uuid"
	"net/http"
)

func (s *APISuite) TestPayOrderSuccess() {
	var (
		orderUUID = uuid.New()
		params    = orderV1.PayOrderParams{
			OrderUUID: orderUUID,
		}
		req = &orderV1.PayOrderRequest{
			PaymentMethod: orderV1.PaymentMethodCARD,
		}
		transUUID        = uuid.New().String()
		expectedResponse = &orderV1.PayOrderResponse{
			TransactionUUID: converter.StringToUUID(transUUID),
		}
	)

	s.orderService.On("PayOrder", s.ctx, orderUUID.String(), string(req.GetPaymentMethod())).Return(transUUID, nil)

	res, err := s.api.PayOrder(s.ctx, req, params)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.PayOrderResponse{}, res)
	s.Require().Equal(expectedResponse.TransactionUUID, res.(*orderV1.PayOrderResponse).TransactionUUID)
}

func (s *APISuite) TestPayOrderPaymentConflict() {
	var (
		orderUUID = uuid.New()
		params    = orderV1.PayOrderParams{
			OrderUUID: orderUUID,
		}
		req = &orderV1.PayOrderRequest{
			PaymentMethod: orderV1.PaymentMethodSBP,
		}
		expectedErr      = model.ErrPaymentConflict
		expectedResponse = &orderV1.ConflictError{
			Code:    http.StatusBadRequest,
			Message: "Заказ уже оплачен или отменен",
		}
	)

	s.orderService.On("PayOrder", s.ctx, orderUUID.String(), string(req.GetPaymentMethod())).Return("", expectedErr)

	res, err := s.api.PayOrder(s.ctx, req, params)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.ConflictError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.ConflictError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.ConflictError).Message)
}

func (s *APISuite) TestPayOrderPaymentNotFound() {
	var (
		orderUUID = uuid.New()
		params    = orderV1.PayOrderParams{
			OrderUUID: orderUUID,
		}
		req = &orderV1.PayOrderRequest{
			PaymentMethod: orderV1.PaymentMethodCREDITCARD,
		}
		expectedErr      = model.ErrPaymentNotFound
		expectedResponse = &orderV1.NotFoundError{
			Code:    http.StatusBadRequest,
			Message: "Заказ не найден или не существует",
		}
	)

	s.orderService.On("PayOrder", s.ctx, orderUUID.String(), string(req.GetPaymentMethod())).Return("", expectedErr)

	res, err := s.api.PayOrder(s.ctx, req, params)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.NotFoundError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.NotFoundError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.NotFoundError).Message)
}

func (s *APISuite) TestPayOrderInternalError() {
	var (
		orderUUID = uuid.New()
		params    = orderV1.PayOrderParams{
			OrderUUID: orderUUID,
		}
		req = &orderV1.PayOrderRequest{
			PaymentMethod: orderV1.PaymentMethodINVESTORMONEY,
		}
		expectedErr      = errors.New("some internal error")
		expectedResponse = &orderV1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "Ошибка сервера при проведении платежа",
		}
	)

	s.orderService.On("PayOrder", s.ctx, orderUUID.String(), string(req.GetPaymentMethod())).Return("", expectedErr)

	res, err := s.api.PayOrder(s.ctx, req, params)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.InternalServerError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.InternalServerError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.InternalServerError).Message)
}
