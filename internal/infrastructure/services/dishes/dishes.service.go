package dishes

import (
	"github.com/mineamihai2001/fm/internal/domain/model"
	"github.com/mineamihai2001/fm/internal/domain/repo"
	"github.com/mineamihai2001/fm/internal/infrastructure/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type DishesService struct {
	dishesRepository      repo.IDishesRepository
	kitchensRepository    repo.IKitchensRepository
	ingredientsRepository repo.IIngredientsRepository
}

func NewDishesService(
	dishesRepository repo.IDishesRepository,
	kitchensRepository repo.IKitchensRepository,
	ingredientsRepository repo.IIngredientsRepository,
) *DishesService {
	return &DishesService{
		dishesRepository:      dishesRepository,
		kitchensRepository:    kitchensRepository,
		ingredientsRepository: ingredientsRepository,
	}
}

func (s *DishesService) Create(
	kitchenId string,
	name string,
	ingredientIds []string,
	duration int,
	rating int,
	images []string,
) (*model.Dish, error) {
	_, err := s.kitchensRepository.GetById(kitchenId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &model.Dish{},
				services.NewServiceError(services.DocumentNotFound, "Kitchen with id %s not found", kitchenId)
		}
		return &model.Dish{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	_, err = s.ingredientsRepository.GetManyById(ingredientIds)
	if err != nil {
		return &model.Dish{},
			services.NewServiceError(services.DocumentNotFound, "Ingredients not found, %s", err.Error())
	}

	res, err := s.dishesRepository.Create(model.NewDish(kitchenId, name, ingredientIds, duration, rating, images))

	if err != nil {
		return &model.Dish{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *DishesService) GetById(id string) (*model.Dish, error) {
	res, err := s.dishesRepository.GetById(id)

	if err != nil {
		return &model.Dish{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *DishesService) GetAll(kitchenId string) (*[]model.Dish, error) {
	res, err := s.dishesRepository.GetAll(kitchenId)

	if err != nil {
		return &[]model.Dish{},
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

func (s *DishesService) GetRandom(kitchenId string) (*model.Dish, error) {
	res, err := s.dishesRepository.GetRandom(kitchenId)

	if err != nil {
		return nil,
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *DishesService) GetPage(page int, pageSize int, sort int, kitchenId string) ([]model.Dish, error) {
	maxPageSize := 100
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	res, err := s.dishesRepository.GetInterval(pageSize, page*pageSize, sort, kitchenId)

	if err != nil {
		return []model.Dish{}, services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}

func (s *DishesService) GetByIngredientIds(ids []string) ([]model.Dish, error) {
	res, err := s.dishesRepository.GetByIngredientIds(ids)

	if err != nil {
		return []model.Dish{},
			services.NewServiceError(services.RepositoryError, err.Error())
	}

	return res, nil
}

// func (s *DishesService) GetDetailsById(id string) (*DishDetails, error) {
// 	res, err := s.dishesRepository.GetById(id)

// 	if err != nil {
// 		return &DishDetails{},
// 			services.NewServiceError(services.RepositoryError, err.Error())
// 	}

// 	return s.Details(res)
// }

// func (s *DishesService) GetDetailsByIngredientIds(ids []string) ([]*DishDetails, error) {
// 	res, err := s.dishesRepository.GetByIngredientIds(ids)

// 	if err != nil {
// 		return []*DishDetails{},
// 			services.NewServiceError(services.RepositoryError, err.Error())
// 	}

// 	detailsResult := make([]*DishDetails, len(res))
// 	for i, doc := range res {
// 		details, err := s.Details(doc)
// 		if err != nil {
// 			return []*DishDetails{},
// 				services.NewServiceError(services.RepositoryError, err.Error())
// 		}
// 		detailsResult[i] = details
// 	}

// 	return detailsResult, nil
// }

// func (s *DishesService) Details(dish model.Dish) (*DishDetails, error) {
// 	ingredients, err := s.ingredientsRepository.GetManyById(dish.IngredientIds)
// 	if err != nil {
// 		return nil, services.NewServiceError(services.RepositoryError, err.Error())
// 	}

// 	kitchen, err := s.kitchensRepository.GetById(dish.KitchenId)
// 	if err != nil {
// 		return nil, services.NewServiceError(services.RepositoryError, err.Error())
// 	}

// 	return NewDishDetails(kitchen, dish.Name, ingredients, dish.Duration, dish.Rating, dish.Images), nil
// }
