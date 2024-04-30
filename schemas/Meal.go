package schemas

type Meal struct {
	Id           int    `json:"id"`
	Hash         string `json:"hash"`
	Name         string `json:"name"`
	Index        int    `json:"index"`
	DishesAmount int    `json:"dishes-amount"`
	Dishes       []int  `json:"dishes"`
}
