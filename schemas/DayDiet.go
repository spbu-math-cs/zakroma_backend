package schemas

type DayDiet struct {
	Id          int      `json:"id"`
	Index       int      `json:"index"`
	Name        string   `json:"name"`
	MealsAmount int      `json:"meals-amount"`
	Meals       []string `json:"meals"`
}
