package v1

import (
	"github.com/baryshnikkov/rocket-factory/inventory/internal/converter"
	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
	"github.com/baryshnikkov/rocket-factory/inventory/internal/testutils"
	inventoryV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/proto/inventory/v1"
	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *APISuite) TestGetPartSuccess() {
	var (
		uuid      = gofakeit.UUID()
		modelPart = testutils.CreatePart()

		req = &inventoryV1.GetPartRequest{
			Uuid: uuid,
		}

		expectedProtoPart = &inventoryV1.GetPartResponse{
			Part: converter.PartToProto(modelPart),
		}
	)

	s.inventoryService.On("GetPart", s.ctx, uuid).Return(modelPart, nil)

	res, err := s.api.GetPart(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Equal(expectedProtoPart.Part, res.GetPart())
}

func (s *APISuite) TestGetFail() {
	var (
		uuid = gofakeit.UUID()

		req = &inventoryV1.GetPartRequest{
			Uuid: uuid,
		}

		expectedErr = model.ErrPartNotFound
	)

	s.inventoryService.On("GetPart", s.ctx, uuid).Return(model.Part{}, expectedErr)

	res, err := s.api.GetPart(s.ctx, req)

	s.Require().Error(err)
	s.Require().Equal(codes.NotFound, status.Code(err))
	s.Require().Contains(err.Error(), "part with UUID")
	s.Require().Empty(res)
}
