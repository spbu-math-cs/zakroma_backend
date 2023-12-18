package schemas

type Group struct {
	Id            int    `json:"id"`
	Hash          string `json:"hash"`
	Name          string `json:"name"`
	CurrentDietId int    `json:"current-diet"`
}
