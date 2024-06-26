package schemas

type Diet struct {
	Id         int    `json:"id" example:"0"`
	Hash       string `json:"hash" example:"d154c7edfc68899a46e017394f5427f3d4a8a5c659570e1e47be9dbdf7e4d1c6"` // 64 символа
	Name       string `json:"name" example:"Чит мил"`
	IsPersonal bool   `json:"is_personal"`
	DayDiets   []int  `json:"day-diets"`
}
