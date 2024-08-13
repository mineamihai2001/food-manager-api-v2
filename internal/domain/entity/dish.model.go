package entity

type Dish struct {
	Id        string   `bson:"_id,omitempty" json:"id"`
	KitchenId string   `bson:"kitchenId" json:"kitchenId"`
	Name      string   `bson:"name" json:"name"`
	Duration  int      `bson:"duration" json:"duration"`
	Rating    int      `bson:"rating" json:"rating"`
	Images    []string `bson:"images" json:"images"`
	Steps     []string `bson:"steps" json:"steps"`
}

func NewDish(kitchenId string, name string, duration int, rating int, images []string, steps []string) Dish {
	return Dish{
		KitchenId: kitchenId,
		Name:      name,
		Duration:  duration,
		Rating:    rating,
		Images:    images,
		Steps:     steps,
	}
}
