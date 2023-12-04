package schemas

type Diet struct {
	Id       int       `json:"id"`
	Hash     string    `json:"hash"`
	Name     string    `json:"name"`
	DayDiets []DayDiet `json:"day-diets"`
}
