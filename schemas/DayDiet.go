package schemas

type DayDiet struct {
	Id          int      `json:"id" example:"0"`
	Index       int      `json:"index" example:"0"`
	Name        string   `json:"name" example:"Кета"`
	MealsAmount int      `json:"meals-amount" example:"1"`
	Meals       []string `json:"meals"`
}
