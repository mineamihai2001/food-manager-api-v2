package dishes

import (
	"slices"

	"github.com/mineamihai2001/fm/internal/domain/entity"
	"github.com/mineamihai2001/fm/internal/domain/model"
	"github.com/mineamihai2001/fm/internal/domain/repo"
	"github.com/mineamihai2001/fm/internal/infrastructure/services"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type DishesService struct {
	dishesRepository          repo.IDishesRepository
	kitchensRepository        repo.IKitchensRepository
	ingredientsRepository     repo.IIngredientsRepository
	dishIngredientsRepository repo.IDishIngredientRepository
}

func NewDishesService(
	dishesRepository repo.IDishesRepository,
	kitchensRepository repo.IKitchensRepository,
	ingredientsRepository repo.IIngredientsRepository,
	dishIngredientsRepository repo.IDishIngredientRepository,
) *DishesService {
	return &DishesService{
		dishesRepository:          dishesRepository,
		kitchensRepository:        kitchensRepository,
		ingredientsRepository:     ingredientsRepository,
		dishIngredientsRepository: dishIngredientsRepository,
	}
}

func (s *DishesService) Create(
	kitchenId string,
	name string,
	duration int,
	rating int,
	images []string,
	steps []string,
	ingredientParts []model.IngredientPart,
) (*model.DishDetails, error) {
	// check if the kitchen exists
	_, err := s.kitchensRepository.GetById(kitchenId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil,
				services.NewServiceError(services.DocumentNotFound, "Kitchen with id %s not found", kitchenId)
		}
		return nil,
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	// create the Dish
	dish, err := s.dishesRepository.Create(entity.NewDish(kitchenId, name, duration, rating, images, steps))
	if err != nil {
		return nil,
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	var ingredientIds = make([]string, 0)
	for _, i := range ingredientParts {
		ingredientIds = append(ingredientIds, i.Id)
	}

	// check if Ingredients exists
	ingredients, err := s.ingredientsRepository.GetManyById(ingredientIds)
	if err != nil {
		return nil,
			services.NewServiceError(services.InternalServerError, err.Error())
	}
	if len(ingredientParts) != len(ingredients) {
		return nil,
			services.NewServiceError(services.DocumentNotFound, "Ingredients with id %s not found", ingredientIds)
	}

	ingredientsDetails := make([]model.IngredientDetails, 0)

	// create the Dish-Ingredients association
	for _, ingredientPart := range ingredientParts {
		_, err := s.dishIngredientsRepository.Create(
			entity.NewDishIngredient(dish.Id, ingredientPart.Id, ingredientPart.Unit, ingredientPart.Size),
		)
		if err != nil {
			return nil,
				services.NewServiceError(services.DocumentNotFound, "Error creating DishIngredient with id %s", ingredientPart.Id)
		}

		// get the details for the current ingredient
		index := slices.IndexFunc(ingredients, func(i entity.Ingredient) bool { return i.Id == ingredientPart.Id })
		if index == -1 {
			return nil,
				services.NewServiceError(services.InternalServerError, "Mismatched ingredient with ingredient part")
		}
		ingredient := ingredients[index]

		ingredientsDetails = append(ingredientsDetails,
			model.NewIngredientDetails(ingredientPart.Id, ingredient.Name, ingredientPart.Unit, ingredientPart.Size),
		)

	}

	dishDetails := model.NewDishDetails(dish.Id, kitchenId, name, duration, rating, images, steps, ingredientsDetails)

	return &dishDetails, nil
}

func (s *DishesService) GetById(id string) (*model.DishDetails, error) {
	// get the Dish
	dish, err := s.dishesRepository.GetById(id)

	if err != nil {
		return nil,
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	dishDetails, err := s.getDetailsFromDish(dish)
	if err != nil {
		return nil, err
	}

	return &dishDetails, nil
}

func (s *DishesService) GetAll(kitchenId string) (*[]entity.Dish, error) {
	res, err := s.dishesRepository.GetAll(kitchenId)

	if err != nil {
		return &[]entity.Dish{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *DishesService) Delete(id string) (bool, error) {
	res, err := s.dishesRepository.Delete(id)

	if err != nil {
		return false,
			services.NewServiceError(services.NotModified, err.Error())
	}

	return res, nil
}

func (s *DishesService) GetRandom(kitchenId string) (*entity.Dish, error) {
	res, err := s.dishesRepository.GetRandom(kitchenId)

	if err != nil {
		return nil,
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *DishesService) GetPage(page int, pageSize int, sort int, kitchenId string) ([]model.DishDetails, error) {
	maxPageSize := 100
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	dishes, err := s.dishesRepository.GetInterval(pageSize, page*pageSize, sort, kitchenId)

	if err != nil {
		return nil, services.NewServiceError(services.RepositoryError, err.Error())
	}

	dishDetailsList := make([]model.DishDetails, 0)
	for _, d := range dishes {
		dishDetails, err := s.getDetailsFromDish(d)
		if err != nil {
			log.Err(err).Msg("Error getting details from dish")
		}
		dishDetailsList = append(dishDetailsList, dishDetails)
	}

	return dishDetailsList, nil
}

// @private
// Creates a DishDetails object from a given Dish
func (s *DishesService) getDetailsFromDish(dish entity.Dish) (model.DishDetails, error) {
	// get the DishIngredient associations
	dishIngredients, err := s.dishIngredientsRepository.GetByDishId(dish.Id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.DishDetails{},
				services.NewServiceError(services.DocumentNotFound, "Dish with id %s not found", dish.Id)
		}
		return model.DishDetails{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	// get then ingredients
	ingredientIds := make([]string, 0)
	for _, di := range dishIngredients {
		ingredientIds = append(ingredientIds, di.IngredientId)
	}
	ingredients, err := s.ingredientsRepository.GetManyById(ingredientIds)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return model.DishDetails{},
				services.NewServiceError(services.InternalServerError, err.Error())
		}
		ingredients = []entity.Ingredient{}
	}

	ingredientsDetails := make([]model.IngredientDetails, 0)
	// create the Dish-Ingredients association
	for _, ingredient := range ingredients {
		// get the details for the current ingredient
		index := slices.IndexFunc(dishIngredients, func(i entity.DishIngredient) bool { return i.IngredientId == ingredient.Id })
		if index == -1 {
			return model.DishDetails{},
				services.NewServiceError(services.InternalServerError, "Mismatched ingredient with dish ingredient")
		}
		dishIngredient := dishIngredients[index]

		ingredientsDetails = append(ingredientsDetails,
			model.NewIngredientDetails(ingredient.Id, ingredient.Name, dishIngredient.Unit, dishIngredient.Size),
		)

	}

	return model.NewDishDetails(dish.Id, dish.KitchenId, dish.Name, dish.Duration, dish.Rating, dish.Images, dish.Steps, ingredientsDetails), nil
}
