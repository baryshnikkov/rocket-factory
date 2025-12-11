package part

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/baryshnikkov/rocket-factory/inventory/internal/model"
	"github.com/baryshnikkov/rocket-factory/inventory/internal/testutils"
)

func (s *RepositorySuite) TestGetPartSuccess() {
	// Arrange
	repoPart := testutils.CreateRepoPart()

	uuid := repoPart.UUID

	s.repo.data[uuid] = repoPart

	// Act
	res, err := s.repo.GetPart(s.ctx, uuid)

	// Assert
	s.Require().NoError(err)
	s.Require().Equal(uuid, res.UUID)
	s.Require().Equal("Main Engine", res.Name)
}

func (s *RepositorySuite) TestGetPartNotFound() {
	// Arrange - НЕ добавляем данные в repository
	uuid := gofakeit.UUID()

	// Act - пытаемся получить несуществующую деталь
	res, err := s.repo.GetPart(s.ctx, uuid)

	// Assert - проверяем что получили ошибку "не найдено"
	s.Require().Error(err)
	s.Require().Equal(model.ErrPartNotFound, err)
	s.Require().Equal(model.Part{}, res)
}

func (s *RepositorySuite) TestGetPartEmptyUUID() {
	// Act
	res, err := s.repo.GetPart(s.ctx, "")

	// Assert
	s.Require().Error(err)
	s.Require().Equal(model.ErrPartNotFound, err)
	s.Require().Equal(model.Part{}, res)
}

func (s *RepositorySuite) TestGetPartNilContext() {
	// Arrange
	repoPart := testutils.CreateRepoPart()
	s.repo.data[repoPart.UUID] = repoPart

	// Act - передаем nil контекст (метод его игнорирует, но тестируем)
	res, err := s.repo.GetPart(nil, repoPart.UUID)

	// Assert - должен работать даже с nil контекстом
	s.Require().NoError(err)
	s.Require().Equal(repoPart.UUID, res.UUID)
}
