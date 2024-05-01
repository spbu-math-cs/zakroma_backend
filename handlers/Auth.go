package handlers

import (
	"net/http"
	"zakroma_backend/middleware"
	"zakroma_backend/schemas"
	"zakroma_backend/stores"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var tokens []string

func generateToken(username string) (string, error) {
	token, _ := middleware.GenerateJWT(username)
	tokens = append(tokens, token)

	return token, nil
}

func Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}

type authResponseBody struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InBldHJpdmFoMUB5YW5kZXgucnUiLCJleHAiOjE3MTY0MDQzNDZ9.XomHB-q6M7mfXp9mryTs01NGpzb0JpJaeZR71ZGcTbY"`
}

// Login godoc
//
// @Tags auth
// @Accept json
// @Produce json
// @Param data body handlers.Login.RequestBody true "Тело запроса"
// @Success 200 {object} authResponseBody
// @Router /auth/login [post]
func Login(c *gin.Context) {
	type RequestBody struct {
		Email    string `json:"email" example:"example@gmail.com"`
		Password string `json:"password" example:"qwerty"`
	}

	var body RequestBody
	err := c.BindJSON(&body)

	user := schemas.User{
		Email:    body.Email,
		Password: body.Password,
	}

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	user.Hash, err = stores.Login(user)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	session := sessions.Default(c)
	session.Set("hash", user.Hash)
	session.Set("group", user.Hash)
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := generateToken(user.Email)

	c.JSON(http.StatusOK, authResponseBody{
		Token: token,
	})
}

// Register godoc
//
// @Tags auth
// @Accept json
// @Produce json
// @Param data body handlers.Register.RequestBody true "Тело запроса"
// @Success 200 {object} authResponseBody
// @Router /auth/register [post]
func Register(c *gin.Context) {
	type RequestBody struct {
		Password  string `json:"password" example:"qwerty"`
		Email     string `json:"email" example:"example@gmail.com"`
		Name      string `json:"name" example:"Ivan"`
		Surname   string `json:"surname" example:"Ivanov"`
		BirthDate string `json:"birth-date" example:"1970-00-00"` // В формате YYYY-MM-DD
	}

	var body RequestBody

	err := c.BindJSON(&body)

	user := schemas.User{
		Email:     body.Email,
		Password:  body.Password,
		Name:      body.Name,
		Surname:   body.Surname,
		BirthDate: body.BirthDate,
	}

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	user.Hash, err = stores.Register(user)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	session := sessions.Default(c)
	session.Set("hash", user.Hash)
	session.Set("group", user.Hash)
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := generateToken(user.Email)

	c.JSON(http.StatusOK, authResponseBody{
		Token: token,
	})
}
