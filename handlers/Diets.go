package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zakroma_backend/stores"
)

func GetDietWithId(c *gin.Context) {
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

	access := stores.CheckDietAccessWithId(id, userId.(int))

	if !access {
		c.Status(http.StatusUnauthorized)
		return
	}

	diet, err := stores.GetDietWithId(id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, diet)
}

func GetDietWithHash(c *gin.Context) {
	hash, err := strconv.Atoi(c.Params.ByName("hash"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dish, err := stores.GetDietWithHash(hash)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, dish)
}

func CreateDiet(c *gin.Context) {
	type RequestBody struct {
		Name string `json:"name"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := stores.CreateDiet(requestBody.Name)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
