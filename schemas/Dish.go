package schemas

type Dish struct {
	Id        int           `json:"id" example:"0"`
	Hash      string        `json:"hash" example:"0ec197ebfcbcf8768574c4faf15d197fac861c323981f46a80236274efd7d4b7"` // Длина 64 символа
	Name      string        `json:"name" example:"Фруктовый салат с йогуртом"`
	Calories  float32       `json:"calories" example:"119.9"`
	Proteins  float32       `json:"proteins" example:"3.5"`
	Fats      float32       `json:"fats" example:"1.6"`
	Carbs     float32       `json:"carbs" example:"23.5"`
	Products  []DishProduct `json:"products" extensions:"x-nullable"`
	ImagePath string        `json:"image-path" example:"https://site.ru/18351.jpg"`
	Recipe    string        `json:"recipe" example:"Шаг 1: Киви почистить и порезать кубиками."`
}
