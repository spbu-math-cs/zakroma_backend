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

func Login(c *gin.Context) {
	var user schemas.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := stores.ValidateUser(user.Username, user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	session := sessions.Default(c)
	session.Set("id", id)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save session",
		})
		return
	}

	token, err := generateToken(user.Username)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
