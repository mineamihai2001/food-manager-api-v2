package dishes

import (
	"github.com/mineamihai2001/fm/cmd/domain"
	"github.com/mineamihai2001/fm/cmd/repository"
	"github.com/mineamihai2001/fm/cmd/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type DishesService struct {
	dishesRepository      *repository.DishesRepository
	kitchensRepository    *repository.KitchensRepository
	ingredientsRepository *repository.IngredientsRepository
}

func New() *DishesService {
	return &DishesService{
		dishesRepository:      repository.NewDishesRepository(),
		kitchensRepository:    repository.NewKitchensRepository(),
		ingredientsRepository: repository.NewIngredientsRepository(),
	}
}

func (s *DishesService) Create(kitchenId string, name string, ingredientIds []string, duration int, rating int, images []string) (*domain.Dish, error) {
	_, err := s.kitchensRepository.GetById(kitchenId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &domain.Dish{},
				services.NewServiceError(services.DocumentNotFound, "Kitchen with id %s not found", kitchenId)
		}
		return &domain.Dish{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	_, err = s.ingredientsRepository.GetManyById(ingredientIds)
	if err != nil {
		return &domain.Dish{},
			services.NewServiceError(services.DocumentNotFound, "Ingredients not found, %s", err.Error())
	}

	res, err := s.dishesRepository.Create(domain.NewDish(kitchenId, name, ingredientIds, duration, rating, images))

	if err != nil {
		return &domain.Dish{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *DishesService) GetById(id string) (*domain.Dish, error) {
	res, err := s.dishesRepository.GetById(id)

	if err != nil {
		return &domain.Dish{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *DishesService) GetAll() (*[]domain.Dish, error) {
	res, err := s.dishesRepository.GetAll()

	if err != nil {
		return &[]domain.Dish{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *DishesService) Delete(id string) (bool, error) {
	res, err := s.dishesRepository.Delete(id)

	if err != nil {
		return false,
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return res, nil
}

func (s *DishesService) GetRandom() (*domain.Dish, error) {
	res, err := s.dishesRepository.GetRandom()

	if err != nil {
		return nil,
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *DishesService) GetPage(page int, pageSize int, sort int) ([]domain.Dish, error) {
	maxPageSize := 100
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	res, err := s.dishesRepository.GetInterval(pageSize, page*pageSize, sort)

	if err != nil {
		return []domain.Dish{}, services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}

func (s *DishesService) GetDetailsById(id string) (*DishDetails, error) {
	res, err := s.dishesRepository.GetById(id)

	if err != nil {
		return &DishDetails{},
			services.NewServiceError(services.RepositoryError, err.Error())
	}

	return s.Details(res)
}

func (s *DishesService) Details(dish domain.Dish) (*DishDetails, error) {
	ingredients, err := s.ingredientsRepository.GetManyById(dish.IngredientIds)
	if err != nil {
		return nil, services.NewServiceError(services.RepositoryError, err.Error())
	}

	kitchen, err := s.kitchensRepository.GetById(dish.KitchenId)
	if err != nil {
		return nil, services.NewServiceError(services.RepositoryError, err.Error())
	}

	return NewDishDetails(kitchen, dish.Name, ingredients, dish.Duration, dish.Rating, dish.Images), nil
}
