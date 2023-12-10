package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"zakroma_backend/middleware"
	"zakroma_backend/schemas"
	"zakroma_backend/stores"
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

func Login(c *gin.Context) {
	var user schemas.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	user.Id, err = stores.Login(user)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	session := sessions.Default(c)
	session.Set("id", user.Id)
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := generateToken(user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Register(c *gin.Context) {
	var user schemas.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := stores.Register(user)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	session := sessions.Default(c)
	session.Set("id", id)
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := generateToken(user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
