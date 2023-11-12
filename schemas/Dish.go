package schemas

type Dish struct {
	Id       int           `json:"id"`
	Name     string        `json:"name"`
	Calories float32       `json:"calories"`
	Proteins float32       `json:"proteins"`
	Fats     float32       `json:"fats"`
	Carbs    float32       `json:"carbs"`
	Products []DishProduct `json:"products"`
	Recipe   string        `json:"recipe"`
}
