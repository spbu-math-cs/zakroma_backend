package schemas

type User struct {
	Id        int    `json:"id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	BirthDate string `json:"birth-date"`
}
