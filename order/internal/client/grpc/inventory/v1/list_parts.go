package v1

import (
	"context"

	"github.com/baryshnikkov/rocket-factory/order/internal/client/converter"
	"github.com/baryshnikkov/rocket-factory/order/internal/model"
	inventoryV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/proto/inventory/v1"
)

func (c *client) ListParts(ctx context.Context, filter model.PartsFilter) (parts []model.Part, error error) {
	partsFilter := &inventoryV1.ListPartsRequest{
		Filter: converter.PartsFilterToProto(filter),
	}

	partsList, err := c.generatedClient.ListParts(ctx, partsFilter)
	if err != nil {
		return nil, err
	}

	return converter.PartListToModel(partsList.Parts), nil
}
