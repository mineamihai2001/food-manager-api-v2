package domain

type Kitchen struct {
	Id   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
}

func NewKitchen(name string) Kitchen {
	return Kitchen{
		Name: name,
	}
}
