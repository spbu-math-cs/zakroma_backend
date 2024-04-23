package schemas

type Meal struct {
	Id           int        `json:"id" example:"0"`
	Hash         string     `json:"hash" example:"d154c7edfc68899a46e017394f5427f3d4a8a5c659570e1e47be9dbdf7e4d1c6"`
	Name         string     `json:"name" example:"Завтрак"`
	Index        int        `json:"index" example:"0"`
	DishesAmount int        `json:"dishes-amount" example:"1"`
	Dishes       []MealDish `json:"dishes"`
}
