package testutils

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"

	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
	repoModel "github.com/baryshnikkov/rocket-factory/inventory/internal/repository/model"
)

//nolint:dupl
func CreatePart() model.Part {
	return model.Part{
		UUID:          gofakeit.UUID(),
		Name:          "Main Engine",
		Description:   "Primary propulsion unit",
		Price:         gofakeit.Float64Range(100, 10000),
		StockQuantity: int64(gofakeit.Number(1, 100)),
		Category:      model.CategoryEngine,
		Dimensions: model.Dimensions{
			Length: gofakeit.Float64Range(0.1, 10),
			Width:  gofakeit.Float64Range(0.1, 10),
			Height: gofakeit.Float64Range(0.1, 10),
			Weight: gofakeit.Float64Range(0.1, 10),
		},
		Manufacturer: model.Manufacturer{
			Name:    gofakeit.Name(),
			Country: gofakeit.Country(),
			Website: gofakeit.URL(),
		},
		Tags: []string{gofakeit.EmojiTag(), gofakeit.EmojiTag()},
		Metadata: model.Metadata{
			StringValue: lo.ToPtr(gofakeit.Word()),
		},
		CreatedAt: gofakeit.Date(),
		UpdatedAt: lo.ToPtr(gofakeit.Date()),
	}
}

//nolint:dupl
func CreateRepoPart() repoModel.Part {
	return repoModel.Part{
		UUID:          gofakeit.UUID(),
		Name:          "Main Engine",
		Description:   "Primary propulsion unit",
		Price:         gofakeit.Float64Range(100, 10000),
		StockQuantity: int64(gofakeit.Number(1, 100)),
		Category:      repoModel.CategoryEngine,
		Dimensions: repoModel.Dimensions{
			Length: gofakeit.Float64Range(0.1, 10),
			Width:  gofakeit.Float64Range(0.1, 10),
			Height: gofakeit.Float64Range(0.1, 10),
			Weight: gofakeit.Float64Range(0.1, 10),
		},
		Manufacturer: repoModel.Manufacturer{
			Name:    gofakeit.Name(),
			Country: gofakeit.Country(),
			Website: gofakeit.URL(),
		},
		Tags: []string{gofakeit.EmojiTag(), gofakeit.EmojiTag()},
		Metadata: repoModel.Metadata{
			StringValue: lo.ToPtr(gofakeit.Word()),
		},
		CreatedAt: gofakeit.Date(),
		UpdatedAt: lo.ToPtr(gofakeit.Date()),
	}
}

func CreateRepoPartWithUUID(uuid string) repoModel.Part {
	part := CreateRepoPart()
	part.UUID = uuid
	return part
}

func CreatePartsFilter() model.PartsFilter {
	partsUUIDs := []string{gofakeit.UUID(), gofakeit.UUID()}
	partsNames := []string{gofakeit.Name(), gofakeit.Name()}
	partsCategories := []model.Category{"UNKNOWN", "ENGINE", "FUEL", "PORTHOLE", "WING"}
	manufactureCountries := []string{gofakeit.Country(), gofakeit.Country()}
	tags := []string{gofakeit.Word(), gofakeit.Word()}

	return model.PartsFilter{
		UUIDs:                 partsUUIDs,
		Names:                 partsNames,
		Categories:            partsCategories,
		ManufacturerCountries: manufactureCountries,
		Tags:                  tags,
	}
}
