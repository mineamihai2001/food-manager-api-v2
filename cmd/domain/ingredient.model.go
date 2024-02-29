package domain

type Ingredient struct {
	Id   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
}

func NewIngredient(name string) Ingredient {
	return Ingredient{
		Name: name,
	}
}
