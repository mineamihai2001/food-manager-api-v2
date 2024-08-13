package entity

type Unit string

type DishIngredient struct {
	Id           string `bson:"_id,omitempty" json:"id"`
	DishId       string `bson:"dishId" json:"dishId"`
	IngredientId string `bson:"ingredientId" json:"ingredientId"`
	Unit         Unit   `bson:"unit" json:"unit"`
	Size         int    `bson:"size" json:"size"`
}

func NewDishIngredient(dishId string, ingredientId string, unit Unit, size int) DishIngredient {
	return DishIngredient{
		DishId:       dishId,
		IngredientId: ingredientId,
		Unit:         unit,
		Size:         size,
	}
}
