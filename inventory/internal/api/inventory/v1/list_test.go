package v1

import (
	"github.com/baryshnikkov/rocket-factory/inventory/internal/converter"
	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
	"github.com/baryshnikkov/rocket-factory/inventory/internal/testutils"
	inventoryV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/proto/inventory/v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *APISuite) TestListPartsSuccess() {
	var (
		modelParts = []model.Part{
			testutils.CreatePart(),
			testutils.CreatePart(),
		}

		req = &inventoryV1.ListPartsRequest{
			Filter: &inventoryV1.PartsFilter{
				Categories: []inventoryV1.Category{
					inventoryV1.Category_CATEGORY_ENGINE,
					inventoryV1.Category_CATEGORY_FUEL,
				},
			},
		}

		expectedResponse = &inventoryV1.ListPartsResponse{
			Parts: converter.PartsToProto(modelParts),
		}
	)

	s.inventoryService.On("ListParts", s.ctx, mock.AnythingOfType("model.PartsFilter")).
		Return(modelParts, nil)

	res, err := s.api.ListParts(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Len(res.Parts, 2)
	s.Require().Equal(expectedResponse.Parts, res.Parts)
}

func (s *APISuite) TestListPartsEmptyFilter() {
	var (
		modelParts = []model.Part{
			testutils.CreatePart(),
		}

		req = &inventoryV1.ListPartsRequest{
			Filter: nil, // Пустой фильтр
		}
	)

	s.inventoryService.On("ListParts", s.ctx, model.PartsFilter{}).
		Return(modelParts, nil)

	res, err := s.api.ListParts(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Len(res.Parts, 1)
}

func (s *APISuite) TestListPartsNotFound() {
	var (
		req = &inventoryV1.ListPartsRequest{
			Filter: &inventoryV1.PartsFilter{
				Uuids: []string{gofakeit.UUID()},
			},
		}
	)

	s.inventoryService.On("ListParts", s.ctx, mock.AnythingOfType("model.PartsFilter")).
		Return([]model.Part{}, model.ErrPartsNotFound)

	res, err := s.api.ListParts(s.ctx, req)

	s.Require().Error(err)
	s.Require().Equal(codes.NotFound, status.Code(err))
	s.Require().Contains(err.Error(), "parts not found")
	s.Require().Nil(res)
}
