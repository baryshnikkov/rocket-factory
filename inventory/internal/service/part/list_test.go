package part

import (
	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
	"github.com/baryshnikkov/rocket-factory/inventory/internal/testutils"

	"github.com/brianvoe/gofakeit/v7"
)

func (s *ServiceSuite) TestListPartsSuccess() {
	filter := testutils.CreatePartsFilter()

	part := testutils.CreatePart()

	expectedParts := []model.Part{part}

	s.inventoryRepository.On("ListParts", s.ctx, filter).Return(expectedParts, nil)

	res, err := s.service.ListParts(s.ctx, filter)
	s.NoError(err)
	s.Equal(expectedParts, res)
}

func (s *ServiceSuite) TestListPartsFail() {
	repoErr := gofakeit.Error()

	filter := testutils.CreatePartsFilter()

	s.inventoryRepository.On("ListParts", s.ctx, filter).Return([]model.Part{}, repoErr)

	res, err := s.service.ListParts(s.ctx, filter)
	s.Error(err)
	s.ErrorIs(err, repoErr)
	s.Empty(res)
}
