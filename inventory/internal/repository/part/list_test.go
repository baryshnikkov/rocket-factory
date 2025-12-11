package part

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
	repoModel "github.com/baryshnikkov/rocket-factory/inventory/internal/repository/model"
	"github.com/baryshnikkov/rocket-factory/inventory/internal/testutils"
)

func (s *RepositorySuite) TestListPartsSuccess() {
	// Arrange
	part1 := testutils.CreateRepoPartWithUUID("uuid-1")
	part2 := testutils.CreateRepoPartWithUUID("uuid-2")
	part3 := testutils.CreateRepoPartWithUUID("uuid-3")

	s.repo.data["uuid-1"] = part1
	s.repo.data["uuid-2"] = part2
	s.repo.data["uuid-3"] = part3

	filter := model.PartsFilter{} // Пустой фильтр - все детали

	// Act
	result, err := s.repo.ListParts(s.ctx, filter)

	// Assert
	s.Require().NoError(err)
	s.Require().Len(result, 3)
}

func (s *RepositorySuite) TestListPartsFilterByUUIDs() {
	// Arrange
	part1 := testutils.CreateRepoPartWithUUID("uuid-1")
	part2 := testutils.CreateRepoPartWithUUID("uuid-2")
	part3 := testutils.CreateRepoPartWithUUID("uuid-3")

	s.repo.data["uuid-1"] = part1
	s.repo.data["uuid-2"] = part2
	s.repo.data["uuid-3"] = part3

	filter := model.PartsFilter{
		UUIDs: []string{"uuid-1", "uuid-3"},
	}

	// Act
	result, err := s.repo.ListParts(s.ctx, filter)

	// Assert
	s.Require().NoError(err)
	s.Require().Len(result, 2)

	// Собираем UUIDs из результата для проверки
	resultUUIDs := make([]string, 0, len(result))
	for _, part := range result {
		resultUUIDs = append(resultUUIDs, part.UUID)
	}

	// Проверяем что в результате есть нужные UUIDs (порядок не важен)
	s.Require().ElementsMatch([]string{"uuid-1", "uuid-3"}, resultUUIDs)
}

func (s *RepositorySuite) TestListPartsFilterByCategory() {
	// Arrange
	enginePart := testutils.CreateRepoPartWithUUID("engine-uuid")
	enginePart.Category = repoModel.CategoryEngine

	fuelPart := testutils.CreateRepoPartWithUUID("fuel-uuid")
	fuelPart.Category = repoModel.CategoryFuel

	s.repo.data["engine-uuid"] = enginePart
	s.repo.data["fuel-uuid"] = fuelPart

	filter := model.PartsFilter{
		Categories: []model.Category{model.CategoryEngine},
	}

	// Act
	result, err := s.repo.ListParts(s.ctx, filter)

	// Assert
	s.Require().NoError(err)
	s.Require().Len(result, 1)
	s.Require().Equal("engine-uuid", result[0].UUID)
	s.Require().Equal(model.CategoryEngine, result[0].Category)
}

func (s *RepositorySuite) TestListPartsFilterByCountry() {
	// Arrange
	usPart := testutils.CreateRepoPartWithUUID("us-uuid")
	usPart.Manufacturer.Country = "USA"

	germanPart := testutils.CreateRepoPartWithUUID("german-uuid")
	germanPart.Manufacturer.Country = "Germany"

	s.repo.data["us-uuid"] = usPart
	s.repo.data["german-uuid"] = germanPart

	filter := model.PartsFilter{
		ManufacturerCountries: []string{"USA"},
	}

	// Act
	result, err := s.repo.ListParts(s.ctx, filter)

	// Assert
	s.Require().NoError(err)
	s.Require().Len(result, 1)
	s.Require().Equal("us-uuid", result[0].UUID)
	s.Require().Equal("USA", result[0].Manufacturer.Country)
}

func (s *RepositorySuite) TestListPartsFilterByTags() {
	// Arrange
	enginePart := testutils.CreateRepoPartWithUUID("engine-uuid")
	enginePart.Tags = []string{"engine", "propulsion", "power"}

	fuelPart := testutils.CreateRepoPartWithUUID("fuel-uuid")
	fuelPart.Tags = []string{"fuel", "liquid", "combustion"}

	s.repo.data["engine-uuid"] = enginePart
	s.repo.data["fuel-uuid"] = fuelPart

	filter := model.PartsFilter{
		Tags: []string{"power", "combustion"}, // Ищем по любому из тегов
	}

	// Act
	result, err := s.repo.ListParts(s.ctx, filter)

	// Assert
	s.Require().NoError(err)
	s.Require().Len(result, 2) // Обе детали должны подойти
}

func (s *RepositorySuite) TestListPartsNotFound() {
	// Arrange - repository пустой после SetupTest
	filter := model.PartsFilter{
		UUIDs: []string{"non-existent-uuid"},
	}

	// Act
	result, err := s.repo.ListParts(s.ctx, filter)

	// Assert
	s.Require().Error(err)
	s.Require().Equal(model.ErrPartsNotFound, err)
	s.Require().Nil(result)
}

func (s *RepositorySuite) TestListPartsMultipleFilters() {
	// Arrange
	usEnginePart := testutils.CreateRepoPartWithUUID("us-engine")
	usEnginePart.Category = repoModel.CategoryEngine
	usEnginePart.Manufacturer.Country = "USA"
	usEnginePart.Tags = []string{"engine", "power"}

	germanEnginePart := testutils.CreateRepoPartWithUUID("german-engine")
	germanEnginePart.Category = repoModel.CategoryEngine
	germanEnginePart.Manufacturer.Country = "Germany"
	germanEnginePart.Tags = []string{"engine", "efficient"}

	usFuelPart := testutils.CreateRepoPartWithUUID("us-fuel")
	usFuelPart.Category = repoModel.CategoryFuel
	usFuelPart.Manufacturer.Country = "USA"
	usFuelPart.Tags = []string{"fuel", "liquid"}

	s.repo.data["us-engine"] = usEnginePart
	s.repo.data["german-engine"] = germanEnginePart
	s.repo.data["us-fuel"] = usFuelPart

	// Фильтр: USA + Engine
	filter := model.PartsFilter{
		Categories:            []model.Category{model.CategoryEngine},
		ManufacturerCountries: []string{"USA"},
	}

	// Act
	result, err := s.repo.ListParts(s.ctx, filter)

	// Assert
	s.Require().NoError(err)
	s.Require().Len(result, 1)
	s.Require().Equal("us-engine", result[0].UUID)
	s.Require().Equal(model.CategoryEngine, result[0].Category)
	s.Require().Equal("USA", result[0].Manufacturer.Country)
}

func (s *RepositorySuite) TestListPartsConcurrentAccess() {
	// Arrange
	for i := 0; i < 10; i++ {
		part := testutils.CreateRepoPartWithUUID(gofakeit.UUID())
		s.repo.data[part.UUID] = part
	}

	// Act - запускаем несколько горутин
	done := make(chan bool, 5)
	for i := 0; i < 5; i++ {
		go func() {
			result, err := s.repo.ListParts(s.ctx, model.PartsFilter{})
			s.Require().NoError(err)
			s.Require().Len(result, 10)
			done <- true
		}()
	}

	// Assert - ждем завершения всех горутин
	for i := 0; i < 5; i++ {
		<-done
	}
}
