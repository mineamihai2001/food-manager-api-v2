package ingredients

import (
	"strings"

	"github.com/mineamihai2001/fm/internal/domain/entity"
	"github.com/mineamihai2001/fm/internal/domain/repo"
	"github.com/mineamihai2001/fm/internal/infrastructure/services"
)

type IngredientsService struct {
	repository repo.IIngredientsRepository
}

func NewIngredientsService(repo repo.IIngredientsRepository) *IngredientsService {
	return &IngredientsService{
		repository: repo,
	}
}

func (s *IngredientsService) Create(name string) (*entity.Ingredient, error) {
	res, err := s.repository.Create(entity.NewIngredient(strings.ToLower(name)))

	if err != nil {
		return &entity.Ingredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *IngredientsService) CreateMany(names []string) ([]entity.Ingredient, error) {
	ingredients := make([]entity.Ingredient, len(names))
	for i, name := range names {
		ingredients[i] = entity.NewIngredient(name)
	}
	res, err := s.repository.CreateMany(ingredients)

	if err != nil {
		return []entity.Ingredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return res, nil
}

func (s *IngredientsService) GetById(id string) (*entity.Ingredient, error) {
	res, err := s.repository.GetById(id)

	if err != nil {
		return &entity.Ingredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *IngredientsService) GetManyById(ids []string) ([]entity.Ingredient, error) {
	res, err := s.repository.GetManyById(ids)

	if err != nil {
		return []entity.Ingredient{},
			services.NewServiceError(services.DocumentNotFound, err.Error())
	}

	return res, nil
}

func (s *IngredientsService) GetAll() ([]entity.Ingredient, error) {
	res, err := s.repository.GetAll()

	if err != nil {
		return []entity.Ingredient{},
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

func (s *IngredientsService) GetPage(page int, pageSize int, sort int) ([]entity.Ingredient, error) {
	maxPageSize := 100
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	res, err := s.repository.GetInterval(pageSize, page*pageSize, sort)

	if err != nil {
		return []entity.Ingredient{}, services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}

func (s *IngredientsService) GetByName(name string) ([]entity.Ingredient, error) {
	res, err := s.repository.GetByName(name)

	if err != nil {
		return []entity.Ingredient{}, services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}
