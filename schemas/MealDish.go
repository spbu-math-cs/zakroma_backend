package schemas

type MealDish struct {
	Dish     Dish    `json:"dish"`
	Portions float32 `json:"portions"`
}
