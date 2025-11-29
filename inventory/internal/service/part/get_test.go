package part

import (
	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
	"github.com/baryshnikkov/rocket-factory/inventory/internal/testutils"
	"github.com/brianvoe/gofakeit/v7"
)

func (s *ServiceSuite) TestGetPartSuccess() {
	part := testutils.CreatePart()
	uuid := part.UUID

	s.inventoryRepository.On("GetPart", s.ctx, uuid).Return(part, nil)

	res, err := s.service.GetPart(s.ctx, uuid)
	s.NoError(err)
	s.Equal(part, res)
}

func (s *ServiceSuite) TestGetPartFail() {
	var (
		repoErr = gofakeit.Error()
		uuid    = gofakeit.UUID()
	)

	s.inventoryRepository.On("GetPart", s.ctx, uuid).Return(model.Part{}, repoErr)

	res, err := s.service.GetPart(s.ctx, uuid)
	s.Error(err)
	s.ErrorIs(err, repoErr)
	s.Empty(res)
}
