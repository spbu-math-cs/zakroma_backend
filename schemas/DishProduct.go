package schemas

type DishProduct struct {
	ProductId         int     `json:"product-id" example:"0"`
	Name              string  `json:"name" example:"Яблоко"`
	Amount            float32 `json:"amount" example:"1"`
	UnitOfMeasurement string  `json:"unit-of-measurement" example:"piece"`
}
