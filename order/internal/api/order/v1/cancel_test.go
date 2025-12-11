package v1

import (
	"errors"

	"github.com/google/uuid"

	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	orderV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/openapi/order/v1"
)

func (s *APISuite) TestCancelOrderSuccess() {
	var (
		uuidData          = uuid.New()
		cancelOrderParams = orderV1.CancelOrderParams{
			OrderUUID: uuidData,
		}
	)

	s.orderService.On("CancelOrder", s.ctx, uuidData.String()).Return(nil)

	res, err := s.api.CancelOrder(s.ctx, cancelOrderParams)

	s.Require().NoError(err)
	s.Require().Nil(res)
}

func (s *APISuite) TestCancelOrderNotFound() {
	var (
		uuidData          = uuid.New()
		cancelOrderParams = orderV1.CancelOrderParams{
			OrderUUID: uuidData,
		}

		expectedErr      = model.ErrOrderNotFound
		expectedResponse = &orderV1.NotFoundError{
			Code:    404,
			Message: "Order by this UUID `" + uuidData.String() + "` not found",
		}
	)

	s.orderService.On("CancelOrder", s.ctx, uuidData.String()).Return(expectedErr)

	res, err := s.api.CancelOrder(s.ctx, cancelOrderParams)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.NotFoundError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.NotFoundError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.NotFoundError).Message)
}

func (s *APISuite) TestCancelOrderAlreadyPaid() {
	var (
		uuidData          = uuid.New()
		cancelOrderParams = orderV1.CancelOrderParams{
			OrderUUID: uuidData,
		}

		expectedErr      = model.ErrOrderAlreadyPaid
		expectedResponse = &orderV1.ConflictError{
			Code:    409,
			Message: "Заказ уже оплачен и не может быть отменён",
		}
	)

	s.orderService.On("CancelOrder", s.ctx, uuidData.String()).Return(expectedErr)

	res, err := s.api.CancelOrder(s.ctx, cancelOrderParams)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.ConflictError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.ConflictError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.ConflictError).Message)
}

func (s *APISuite) TestCancelOrderAlreadyCancelled() {
	var (
		uuidData          = uuid.New()
		cancelOrderParams = orderV1.CancelOrderParams{
			OrderUUID: uuidData,
		}

		expectedErr      = model.ErrOrderAlreadyCancelled
		expectedResponse = &orderV1.ConflictError{
			Code:    409,
			Message: "Заказ уже отменён",
		}
	)

	s.orderService.On("CancelOrder", s.ctx, uuidData.String()).Return(expectedErr)

	res, err := s.api.CancelOrder(s.ctx, cancelOrderParams)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.ConflictError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.ConflictError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.ConflictError).Message)
}

func (s *APISuite) TestCancelOrderInternalError() {
	var (
		uuidData          = uuid.New()
		cancelOrderParams = orderV1.CancelOrderParams{
			OrderUUID: uuidData,
		}

		expectedErr      = errors.New("some internal error")
		expectedResponse = &orderV1.InternalServerError{
			Code:    500,
			Message: "Внутренняя ошибка сервера",
		}
	)

	s.orderService.On("CancelOrder", s.ctx, uuidData.String()).Return(expectedErr)

	res, err := s.api.CancelOrder(s.ctx, cancelOrderParams)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.InternalServerError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.InternalServerError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.InternalServerError).Message)
}
