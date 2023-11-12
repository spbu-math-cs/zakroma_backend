package schemas

type DayDiet struct {
	Id    int    `json:"id"`
	Index int    `json:"index"`
	Meals []Meal `json:"meals"`
}
