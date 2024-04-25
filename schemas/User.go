package schemas

type User struct {
	Id        int    `json:"id"`
	Hash      string `json:"hash"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	BirthDate string `json:"birth-date"` // В формате YYYY-MM-DD
}
