package schemas

type Product struct {
	Id                int     `json:"id" example:"2"`
	Name              string  `json:"name" example:"Мандарин"`
	Calories          float32 `json:"calories" example:"33"`
	Proteins          float32 `json:"proteins" example:"0.8"`
	Fats              float32 `json:"fats" example:"0.2"`
	Carbs             float32 `json:"carbs" example:"7.5"`
	UnitOfMeasurement string  `json:"unit-of-measurement" example:"piece"`
}
