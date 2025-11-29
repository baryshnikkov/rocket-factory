package order

import (
	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	"github.com/brianvoe/gofakeit/v7"
)

func (s *ServiceSuite) TestCreateOrderSuccess() {
	userUUID := gofakeit.UUID()
	partsUUIDs := []string{gofakeit.UUID(), gofakeit.UUID()}
	partsList := []model.Part{
		{UUID: partsUUIDs[0], Price: 1000.0},
		{UUID: partsUUIDs[1], Price: 500.0},
	}
	orderInfo := model.OrderCreationInfo{
		OrderUUID:  gofakeit.UUID(),
		TotalPrice: 1500.0,
	}

	s.inventoryClient.On("ListParts", s.ctx, model.PartsFilter{Uuids: partsUUIDs}).Return(partsList, nil)
	s.orderRepository.On("CreateOrder", s.ctx, userUUID, partsList).Return(orderInfo, nil)

	res, err := s.service.CreateOrder(s.ctx, userUUID, partsUUIDs)

	s.NoError(err)
	s.Equal(orderInfo.OrderUUID, res.OrderUUID)
	s.Equal(orderInfo.TotalPrice, res.TotalPrice)
}

func (s *ServiceSuite) TestCreateOrderInventoryClientError() {
	userUUID := gofakeit.UUID()
	partsUUIDs := []string{gofakeit.UUID()}
	expectedErr := model.ErrPartsNotFound

	s.inventoryClient.On("ListParts", s.ctx, model.PartsFilter{Uuids: partsUUIDs}).Return([]model.Part{}, expectedErr)

	res, err := s.service.CreateOrder(s.ctx, userUUID, partsUUIDs)

	s.Error(err)
	s.ErrorIs(err, expectedErr)
	s.Equal(model.OrderCreationInfo{}, res)
}

func (s *ServiceSuite) TestCreateOrderPartsNotFound() {
	userUUID := gofakeit.UUID()
	partsUUIDs := []string{gofakeit.UUID(), gofakeit.UUID()}
	partsList := []model.Part{
		{UUID: partsUUIDs[0], Price: 1000.0},
		// Второй части нет в ответе
	}

	s.inventoryClient.On("ListParts", s.ctx, model.PartsFilter{Uuids: partsUUIDs}).Return(partsList, nil)

	res, err := s.service.CreateOrder(s.ctx, userUUID, partsUUIDs)

	s.Error(err)
	s.ErrorIs(err, model.ErrOrderConflict)
	s.Equal(model.OrderCreationInfo{}, res)
}

func (s *ServiceSuite) TestCreateOrderRepositoryError() {
	userUUID := gofakeit.UUID()
	partsUUIDs := []string{gofakeit.UUID()}
	partsList := []model.Part{
		{UUID: partsUUIDs[0], Price: 1000.0},
	}
	expectedErr := model.ErrOrderInternalError

	s.inventoryClient.On("ListParts", s.ctx, model.PartsFilter{Uuids: partsUUIDs}).Return(partsList, nil)
	s.orderRepository.On("CreateOrder", s.ctx, userUUID, partsList).Return(model.OrderCreationInfo{}, expectedErr)

	res, err := s.service.CreateOrder(s.ctx, userUUID, partsUUIDs)

	s.Error(err)
	s.ErrorIs(err, expectedErr)
	s.Equal(model.OrderCreationInfo{}, res)
}

func (s *ServiceSuite) TestCreateOrderEmptyParts() {
	userUUID := gofakeit.UUID()
	var partsUUIDs []string
	orderInfo := model.OrderCreationInfo{
		OrderUUID:  gofakeit.UUID(),
		TotalPrice: 0.0,
	}

	s.inventoryClient.On("ListParts", s.ctx, model.PartsFilter{Uuids: partsUUIDs}).Return([]model.Part{}, nil)
	s.orderRepository.On("CreateOrder", s.ctx, userUUID, []model.Part{}).Return(orderInfo, nil)

	res, err := s.service.CreateOrder(s.ctx, userUUID, partsUUIDs)

	s.NoError(err)
	s.Equal(orderInfo.OrderUUID, res.OrderUUID)
	s.Equal(0.0, res.TotalPrice)
}
