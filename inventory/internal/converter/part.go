package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
	inventoryV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/proto/inventory/v1"
)

func PartToProto(part model.Part) *inventoryV1.Part {
	var updatedAt *timestamppb.Timestamp
	if part.UpdatedAt != nil {
		updatedAt = timestamppb.New(*part.UpdatedAt)
	}

	return &inventoryV1.Part{
		Uuid:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      categoryToProto(part.Category),
		Dimensions:    dimensionsToProto(part.Dimensions),
		Manufacturer:  manufacturerToProto(part.Manufacturer),
		Tags:          part.Tags,
		Metadata:      metadataToProto(part.Metadata),
		CreatedAt:     timestamppb.New(part.CreatedAt),
		UpdatedAt:     updatedAt,
	}
}

func categoryToProto(category model.Category) inventoryV1.Category {
	switch category {
	case model.CategoryEngine:
		return inventoryV1.Category_CATEGORY_ENGINE
	case model.CategoryFuel:
		return inventoryV1.Category_CATEGORY_FUEL
	case model.CategoryPorthole:
		return inventoryV1.Category_CATEGORY_PORTHOLE
	case model.CategoryWing:
		return inventoryV1.Category_CATEGORY_WING
	default:
		return inventoryV1.Category_CATEGORY_UNSPECIFIED
	}
}

func dimensionsToProto(dims model.Dimensions) *inventoryV1.Dimensions {
	return &inventoryV1.Dimensions{
		Length: dims.Length,
		Width:  dims.Width,
		Height: dims.Height,
		Weight: dims.Weight,
	}
}

func manufacturerToProto(man model.Manufacturer) *inventoryV1.Manufacturer {
	return &inventoryV1.Manufacturer{
		Name:    man.Name,
		Country: man.Country,
		Website: man.Website,
	}
}

func metadataToProto(meta model.Metadata) map[string]*inventoryV1.Value {
	var val *inventoryV1.Value
	switch {
	case meta.StringValue != nil:
		val = &inventoryV1.Value{
			Kind: &inventoryV1.Value_StringValue{StringValue: *meta.StringValue},
		}
	case meta.Int64Value != nil:
		val = &inventoryV1.Value{
			Kind: &inventoryV1.Value_Int64Value{Int64Value: *meta.Int64Value},
		}
	case meta.DoubleValue != nil:
		val = &inventoryV1.Value{
			Kind: &inventoryV1.Value_DoubleValue{DoubleValue: *meta.DoubleValue},
		}
	case meta.BoolValue != nil:
		val = &inventoryV1.Value{
			Kind: &inventoryV1.Value_BoolValue{BoolValue: *meta.BoolValue},
		}
	default:
		val = &inventoryV1.Value{}
	}
	return map[string]*inventoryV1.Value{"value": val}
}

func PartsFilterToModel(filter *inventoryV1.PartsFilter) model.PartsFilter {
	partsUUIDs := make([]string, 0, len(filter.Uuids))
	if len(filter.Uuids) > 0 {
		partsUUIDs = filter.Uuids
	}

	partsNames := make([]string, 0, len(filter.Names))
	if len(filter.Names) > 0 {
		partsUUIDs = filter.Names
	}

	partsTags := make([]string, 0, len(filter.Tags))
	if len(filter.Tags) > 0 {
		partsUUIDs = filter.Tags
	}

	partsManufacturerCountries := make([]string, 0, len(filter.ManufacturerCountries))
	if len(filter.ManufacturerCountries) > 0 {
		partsUUIDs = filter.ManufacturerCountries
	}

	partsCategories := make([]model.Category, 0, len(filter.Categories))
	if len(filter.Categories) > 0 {
		for _, c := range filter.Categories {
			partsCategories = append(partsCategories, model.Category(c.String()))
		}
	}

	return model.PartsFilter{
		UUIDs:                 partsUUIDs,
		Names:                 partsNames,
		Categories:            partsCategories,
		ManufacturerCountries: partsManufacturerCountries,
		Tags:                  partsTags,
	}
}

func PartsToProto(parts []model.Part) []*inventoryV1.Part {
	result := make([]*inventoryV1.Part, 0, len(parts))
	for _, part := range parts {
		result = append(result, PartToProto(part))
	}
	return result
}
