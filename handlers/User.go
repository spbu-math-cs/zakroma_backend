package handlers

import (
	"fmt"
	"net/http"
	"zakroma_backend/stores"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetName godoc
//
// @Tags user
// @Produce json
// @Success 200 {object} handlers.GetName.ResponceBody
// @Router /api/user/name [get]
func GetName(c *gin.Context) {
	type ResponceBody struct {
		FirstName  string `json:"name" example:"Ivan"`
		SecondName string `json:"surname" example:"Ivanov"`
	}

	session := sessions.Default(c)
	hash := session.Get("hash")

	name, surname, err := stores.GetUserInits(fmt.Sprint(hash))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, ResponceBody{
		FirstName:  name,
		SecondName: surname,
	})
}
