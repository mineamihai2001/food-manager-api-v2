package ingredients

import (
	"github.com/mineamihai2001/fm/cmd/domain"
	"github.com/mineamihai2001/fm/cmd/repository"
	"github.com/mineamihai2001/fm/cmd/services"
)

type IngredientsService struct {
	repository *repository.IngredientsRepository
}

func New() *IngredientsService {
	return &IngredientsService{
		repository: repository.NewIngredientsRepository(),
	}
}

func (s *IngredientsService) Create(name string) (*domain.Ingredient, error) {
	res, err := s.repository.Create(domain.NewIngredient(name))

	if err != nil {
		return &domain.Ingredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *IngredientsService) GetById(id string) (*domain.Ingredient, error) {
	res, err := s.repository.GetById(id)

	if err != nil {
		return &domain.Ingredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *IngredientsService) GetManyById(ids []string) ([]domain.Ingredient, error) {
	res, err := s.repository.GetManyById(ids)

	if err != nil {
		return []domain.Ingredient{},
			services.NewServiceError(services.DocumentNotFound, err.Error())
	}

	return res, nil
}

func (s *IngredientsService) GetAll() ([]domain.Ingredient, error) {
	res, err := s.repository.GetAll()

	if err != nil {
		return []domain.Ingredient{},
			services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}

func (s *IngredientsService) Delete(id string) (bool, error) {
	res, err := s.repository.Delete(id)

	if err != nil {
		return false,
			services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}

func (s *IngredientsService) DeleteMany(ids []string) (int, error) {
	res, err := s.repository.DeleteMany(ids)

	if err != nil {
		return 0,
			services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}

func (s *IngredientsService) GetPage(page int, pageSize int, sort int) ([]domain.Ingredient, error) {
	maxPageSize := 100
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	res, err := s.repository.GetInterval(pageSize, page*pageSize, sort)

	if err != nil {
		return []domain.Ingredient{}, services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}

func (s *IngredientsService) GetByName(name string) ([]domain.Ingredient, error) {
	res, err := s.repository.GetByName(name)

	if err != nil {
		return []domain.Ingredient{}, services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}
