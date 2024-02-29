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