package schemas

type Diet struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	DayDiets []DayDiet `json:"day-diets"`
}
