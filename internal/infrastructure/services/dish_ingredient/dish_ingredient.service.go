package dish_ingredient

import (
	"github.com/mineamihai2001/fm/internal/domain/entity"
	"github.com/mineamihai2001/fm/internal/domain/repo"
	"github.com/mineamihai2001/fm/internal/infrastructure/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type DishIngredientService struct {
	dishIngredientRepository repo.IDishIngredientRepository
	dishesRepository         repo.IDishesRepository
	ingredientsRepository    repo.IIngredientsRepository
}

func NewDishIngredientService(
	dishIngredientRepository repo.IDishIngredientRepository,
	dishesRepository repo.IDishesRepository,
	ingredientsRepository repo.IIngredientsRepository,
) *DishIngredientService {
	return &DishIngredientService{
		dishIngredientRepository,
		dishesRepository,
		ingredientsRepository,
	}
}

func (s *DishIngredientService) GetById(id string) (entity.DishIngredient, error) {
	res, err := s.dishIngredientRepository.GetById(id)

	if err != nil {
		return entity.DishIngredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return res, nil
}

func (s *DishIngredientService) Delete(id string) (bool, error) {
	res, err := s.dishIngredientRepository.Delete(id)

	if err != nil {
		return false,
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return res, nil
}

func (s *DishIngredientService) Create(dishId string, ingredientId string, unit entity.Unit, size int) (entity.DishIngredient, error) {
	_, err := s.dishesRepository.GetById(dishId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.DishIngredient{},
				services.NewServiceError(services.DocumentNotFound, "Dish with id %s not found", dishId)
		}
		return entity.DishIngredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	_, e := s.ingredientsRepository.GetById(ingredientId)
	if e != nil {
		if e == mongo.ErrNoDocuments {
			return entity.DishIngredient{},
				services.NewServiceError(services.DocumentNotFound, "Ingredient with id %s not found", dishId)
		}
		return entity.DishIngredient{},
			services.NewServiceError(services.InternalServerError, e.Error())
	}

	// check to see if this association already exists
	count, err := s.dishIngredientRepository.CountDishIngredient(dishId, ingredientId)
	if err != nil {
		return entity.DishIngredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	if count != 0 {
		return entity.DishIngredient{},
			services.NewServiceError(services.Conflict, "Dish ingredient already exists")
	}

	res, err := s.dishIngredientRepository.Create(entity.NewDishIngredient(dishId, ingredientId, unit, size))

	if err != nil {
		return entity.DishIngredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return res, nil
}

func (s *DishIngredientService) GetByDishId(dishId string) ([]entity.DishIngredient, error) {
	res, err := s.dishIngredientRepository.GetByDishId(dishId)

	if err != nil {
		return []entity.DishIngredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return res, nil
}

func (s *DishIngredientService) GetByIngredientId(ingredientId string) ([]entity.DishIngredient, error) {
	res, err := s.dishIngredientRepository.GetByDishId(ingredientId)

	if err != nil {
		return []entity.DishIngredient{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return res, nil
}
