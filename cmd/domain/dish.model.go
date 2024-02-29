package domain

type Dish struct {
	Id            string   `bson:"_id,omitempty" json:"id"`
	KitchenId     string   `bson:"kitchenId" json:"kitchenId"`
	Name          string   `bson:"name" json:"name"`
	IngredientIds []string `bson:"ingredients" json:"ingredients"`
	Duration      int      `bson:"duration" json:"duration"`
	Rating        int      `bson:"rating" json:"rating"`
	Images        []string `bson:"images" json:"images"`
}

func NewDish(kitchenId string, name string, ingredientIds []string, duration int, rating int, images []string) Dish {
	return Dish{
		KitchenId:     kitchenId,
		Name:          name,
		IngredientIds: ingredientIds,
		Duration:      duration,
		Rating:        rating,
		Images:        images,
	}
}
