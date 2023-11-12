package schemas

type Meal struct {
	Id     int        `json:"id"`
	Name   string     `json:"name"`
	Index  int        `json:"index"`
	Dishes []MealDish `json:"dishes"`
}
