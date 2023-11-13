package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zakroma_backend/stores"
)

func GetMealWithId(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	session := sessions.Default(c)
	userId := session.Get("id")

	if userId == nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	access := stores.CheckMealAccessWithId(id, userId.(int))

	if !access {
		c.Status(http.StatusUnauthorized)
		return
	}

	meal, err := stores.GetMealWithId(id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, meal)
}
