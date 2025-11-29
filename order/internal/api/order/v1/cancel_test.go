package v1

import (
	"errors"
	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	orderV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/openapi/order/v1"
	"github.com/google/uuid"
)

func (s *APISuite) TestCancelOrderSuccess() {
	var (
		uuid              = uuid.New()
		cancelOrderParams = orderV1.CancelOrderParams{
			OrderUUID: uuid,
		}
	)

	s.orderService.On("CancelOrder", s.ctx, uuid.String()).Return(nil)

	res, err := s.api.CancelOrder(s.ctx, cancelOrderParams)

	s.Require().NoError(err)
	s.Require().Nil(res)
}

func (s *APISuite) TestCancelOrderNotFound() {
	var (
		uuid              = uuid.New()
		cancelOrderParams = orderV1.CancelOrderParams{
			OrderUUID: uuid,
		}

		expectedErr      = model.ErrOrderNotFound
		expectedResponse = &orderV1.NotFoundError{
			Code:    404,
			Message: "Order by this UUID `" + uuid.String() + "` not found",
		}
	)

	s.orderService.On("CancelOrder", s.ctx, uuid.String()).Return(expectedErr)

	res, err := s.api.CancelOrder(s.ctx, cancelOrderParams)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.NotFoundError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.NotFoundError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.NotFoundError).Message)
}

func (s *APISuite) TestCancelOrderAlreadyPaid() {
	var (
		uuid              = uuid.New()
		cancelOrderParams = orderV1.CancelOrderParams{
			OrderUUID: uuid,
		}

		expectedErr      = model.ErrOrderAlreadyPaid
		expectedResponse = &orderV1.ConflictError{
			Code:    409,
			Message: "Заказ уже оплачен и не может быть отменён",
		}
	)

	s.orderService.On("CancelOrder", s.ctx, uuid.String()).Return(expectedErr)

	res, err := s.api.CancelOrder(s.ctx, cancelOrderParams)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.ConflictError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.ConflictError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.ConflictError).Message)
}

func (s *APISuite) TestCancelOrderAlreadyCancelled() {
	var (
		uuid              = uuid.New()
		cancelOrderParams = orderV1.CancelOrderParams{
			OrderUUID: uuid,
		}

		expectedErr      = model.ErrOrderAlreadyCancelled
		expectedResponse = &orderV1.ConflictError{
			Code:    409,
			Message: "Заказ уже отменён",
		}
	)

	s.orderService.On("CancelOrder", s.ctx, uuid.String()).Return(expectedErr)

	res, err := s.api.CancelOrder(s.ctx, cancelOrderParams)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.ConflictError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.ConflictError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.ConflictError).Message)
}

func (s *APISuite) TestCancelOrderInternalError() {
	var (
		uuid              = uuid.New()
		cancelOrderParams = orderV1.CancelOrderParams{
			OrderUUID: uuid,
		}

		expectedErr      = errors.New("some internal error")
		expectedResponse = &orderV1.InternalServerError{
			Code:    500,
			Message: "Внутренняя ошибка сервера",
		}
	)

	s.orderService.On("CancelOrder", s.ctx, uuid.String()).Return(expectedErr)

	res, err := s.api.CancelOrder(s.ctx, cancelOrderParams)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.InternalServerError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.InternalServerError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.InternalServerError).Message)
}
