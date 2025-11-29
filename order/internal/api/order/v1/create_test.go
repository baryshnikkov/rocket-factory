package v1

import (
	"errors"
	"github.com/baryshnikkov/rocket-factory/order/internal/converter"
	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	orderV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/openapi/order/v1"
	"github.com/google/uuid"
	"net/http"
)

func (s *APISuite) TestCreateOrderSuccess() {
	var (
		userUUID  = uuid.New()
		partUUIDs = []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}

		req = &orderV1.CreateOrderRequest{
			UserUUID: converter.StringToUUID(userUUID.String()),
			PartUuids: []uuid.UUID{
				converter.StringToUUID(partUUIDs[0].String()),
				converter.StringToUUID(partUUIDs[1].String()),
				converter.StringToUUID(partUUIDs[2].String()),
			},
		}

		orderInfo = model.OrderCreationInfo{
			OrderUUID:  uuid.New().String(),
			TotalPrice: 1500.75,
		}

		expectedResponse = &orderV1.CreateOrderResponse{
			OrderUUID:  converter.StringToUUID(orderInfo.OrderUUID),
			TotalPrice: orderInfo.TotalPrice,
		}
	)

	partUUIDStrings := make([]string, len(partUUIDs))
	for i, u := range partUUIDs {
		partUUIDStrings[i] = u.String()
	}

	s.orderService.On("CreateOrder", s.ctx, userUUID.String(), partUUIDStrings).Return(orderInfo, nil)

	res, err := s.api.CreateOrder(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.CreateOrderResponse{}, res)
	s.Require().Equal(expectedResponse.OrderUUID, res.(*orderV1.CreateOrderResponse).OrderUUID)
	s.Require().Equal(expectedResponse.TotalPrice, res.(*orderV1.CreateOrderResponse).TotalPrice)
}

func (s *APISuite) TestCreateOrderPartsNotFound() {
	var (
		userUUID  = uuid.New()
		partUUIDs = []uuid.UUID{uuid.New(), uuid.New()}

		req = &orderV1.CreateOrderRequest{
			UserUUID: converter.StringToUUID(userUUID.String()),
			PartUuids: []uuid.UUID{
				converter.StringToUUID(partUUIDs[0].String()),
				converter.StringToUUID(partUUIDs[1].String()),
			},
		}

		expectedErr      = model.ErrPartsNotFound
		expectedResponse = orderV1.NotFoundError{
			Code:    http.StatusNotFound,
			Message: "Одна или несколько частей не найдены",
		}
	)

	partUUIDStrings := make([]string, len(partUUIDs))
	for i, u := range partUUIDs {
		partUUIDStrings[i] = u.String()
	}

	s.orderService.On("CreateOrder", s.ctx, userUUID.String(), partUUIDStrings).Return(model.OrderCreationInfo{}, expectedErr)

	res, err := s.api.CreateOrder(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().IsType(&orderV1.NotFoundError{}, res)
	s.Require().Equal(expectedResponse.Code, res.(*orderV1.NotFoundError).Code)
	s.Require().Equal(expectedResponse.Message, res.(*orderV1.NotFoundError).Message)
}

func (s *APISuite) TestCreateOrderInternalError() {
	var (
		userUUID  = uuid.New()
		partUUIDs = []uuid.UUID{uuid.New()}

		req = &orderV1.CreateOrderRequest{
			UserUUID: converter.StringToUUID(userUUID.String()),
			PartUuids: []uuid.UUID{
				converter.StringToUUID(partUUIDs[0].String()),
			},
		}

		expectedErr = errors.New("database error")
	)

	partUUIDStrings := make([]string, len(partUUIDs))
	for i, u := range partUUIDs {
		partUUIDStrings[i] = u.String()
	}

	s.orderService.On("CreateOrder", s.ctx, userUUID.String(), partUUIDStrings).Return(model.OrderCreationInfo{}, expectedErr)

	res, err := s.api.CreateOrder(s.ctx, req)

	s.Require().Error(err)
	s.Require().Equal(expectedErr, err)
	s.Require().Nil(res)
}
