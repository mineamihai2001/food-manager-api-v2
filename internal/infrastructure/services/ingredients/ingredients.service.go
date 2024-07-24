package ingredients

import (
	"github.com/mineamihai2001/fm/internal/domain/model"
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

func (s *IngredientsService) Create(name string) (*model.Ingredient, error) {
	res, err := s.repository.Create(model.NewIngredient(name))

	if err != nil {
		return &model.Ingredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *IngredientsService) CreateMany(names []string) ([]model.Ingredient, error) {
	ingredients := make([]model.Ingredient, len(names))
	for i, name := range names {
		ingredients[i] = model.NewIngredient(name)
	}
	res, err := s.repository.CreateMany(ingredients)

	if err != nil {
		return []model.Ingredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return res, nil
}

func (s *IngredientsService) GetById(id string) (*model.Ingredient, error) {
	res, err := s.repository.GetById(id)

	if err != nil {
		return &model.Ingredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *IngredientsService) GetManyById(ids []string) ([]model.Ingredient, error) {
	res, err := s.repository.GetManyById(ids)

	if err != nil {
		return []model.Ingredient{},
			services.NewServiceError(services.DocumentNotFound, err.Error())
	}

	return res, nil
}

func (s *IngredientsService) GetAll() ([]model.Ingredient, error) {
	res, err := s.repository.GetAll()

	if err != nil {
		return []model.Ingredient{},
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

func (s *IngredientsService) GetPage(page int, pageSize int, sort int) ([]model.Ingredient, error) {
	maxPageSize := 100
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	res, err := s.repository.GetInterval(pageSize, page*pageSize, sort)

	if err != nil {
		return []model.Ingredient{}, services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}

func (s *IngredientsService) GetByName(name string) ([]model.Ingredient, error) {
	res, err := s.repository.GetByName(name)

	if err != nil {
		return []model.Ingredient{}, services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}
