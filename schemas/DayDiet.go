package schemas

type DayDiet struct {
	Id          int    `json:"id"`
	Index       int    `json:"index"`
	Name        int    `json:"name"`
	MealsAmount int    `json:"meals-amount"`
	Meals       []Meal `json:"meals"`
}
