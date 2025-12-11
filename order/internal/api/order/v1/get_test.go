package v1

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/baryshnikkov/rocket-factory/order/internal/converter"
	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	orderV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/openapi/order/v1"
)

func (s *APISuite) TestGetOrderSuccess() {
	var (
		orderUUID = uuid.New()
		params    = orderV1.GetOrderParams{
			OrderUUID: orderUUID,
		}
		order = model.OrderDto{
			UUID:       orderUUID.String(),
			UserUUID:   uuid.New().String(),
			PartsUUIDs: []string{uuid.New().String(), uuid.New().String()},
			TotalPrice: 1500.75,
			Status:     model.OrderStatusPaid,
			CreatedAt:  time.Now(),
		}
		expectedResponse = &orderV1.GetOrderResponse{
			Data: converter.OrderDataToDTO(order),
		}
	)

	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(order, nil)

	res, err := s.api.GetOrder(s.ctx, params)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.GetOrderResponse{}, res)
	s.Require().Equal(expectedResponse.Data, res.(*orderV1.GetOrderResponse).Data)
}

func (s *APISuite) TestGetOrderNotFound() {
	var (
		orderUUID = uuid.New()
		params    = orderV1.GetOrderParams{
			OrderUUID: orderUUID,
		}
		expectedErr = model.ErrOrderNotFound
	)

	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.OrderDto{}, expectedErr)

	res, err := s.api.GetOrder(s.ctx, params)

	s.Require().Error(err)
	s.Require().Equal(codes.NotFound, status.Code(err))
	s.Require().Contains(err.Error(), "order by this UUID")
	s.Require().Nil(res)
}

func (s *APISuite) TestGetOrderContextDeadlineExceeded() {
	var (
		orderUUID = uuid.New()
		params    = orderV1.GetOrderParams{
			OrderUUID: orderUUID,
		}
		expectedErr = context.DeadlineExceeded
	)

	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.OrderDto{}, expectedErr)

	res, err := s.api.GetOrder(s.ctx, params)

	s.Require().Error(err)
	s.Require().Equal(codes.Unavailable, status.Code(err))
	s.Require().Contains(err.Error(), "Order service timeout")
	s.Require().Nil(res)
}

func (s *APISuite) TestGetOrderContextCanceled() {
	var (
		orderUUID = uuid.New()
		params    = orderV1.GetOrderParams{
			OrderUUID: orderUUID,
		}
		expectedErr = context.Canceled
	)

	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.OrderDto{}, expectedErr)

	res, err := s.api.GetOrder(s.ctx, params)

	s.Require().Error(err)
	s.Require().Equal(codes.Unavailable, status.Code(err))
	s.Require().Contains(err.Error(), "Order service timeout")
	s.Require().Nil(res)
}

func (s *APISuite) TestGetOrderInternalError() {
	var (
		orderUUID = uuid.New()
		params    = orderV1.GetOrderParams{
			OrderUUID: orderUUID,
		}
		expectedErr = model.ErrOrderInternalError
	)

	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.OrderDto{}, expectedErr)

	res, err := s.api.GetOrder(s.ctx, params)

	s.Require().Error(err)
	s.Require().Equal(codes.Internal, status.Code(err))
	s.Require().Contains(err.Error(), "Order service internal error")
	s.Require().Nil(res)
}

func (s *APISuite) TestGetOrderUnknownError() {
	var (
		orderUUID = uuid.New()
		params    = orderV1.GetOrderParams{
			OrderUUID: orderUUID,
		}
		expectedErr = errors.New("unknown error")
	)

	s.orderService.On("GetOrder", s.ctx, orderUUID.String()).Return(model.OrderDto{}, expectedErr)

	res, err := s.api.GetOrder(s.ctx, params)

	s.Require().Error(err)
	s.Require().Equal(expectedErr, err)
	s.Require().Nil(res)
}
