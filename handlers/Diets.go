package handlers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"zakroma_backend/stores"
)

func GetDietByHash(c *gin.Context) {
	hash := c.Params.ByName("hash")
	if len(hash) == 0 {
		c.String(http.StatusBadRequest, "something bad with field 'hash'")
		return
	}

	diet, err := stores.GetDietByHash(hash)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, diet)
}

func CreateDiet(c *gin.Context) {
	type RequestBody struct {
		Name string `json:"name"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	session := sessions.Default(c)
	groupHash := session.Get("group")
	user := session.Get("hash")

	hash, err := stores.CreateDiet(requestBody.Name, fmt.Sprint(groupHash))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := stores.ChangeCurrentDiet(fmt.Sprint(user), fmt.Sprint(groupHash), hash); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"hash": hash})
}

func GetCurrentDiet(c *gin.Context) {
	session := sessions.Default(c)
	groupHash := session.Get("group")

	diet, err := stores.GetCurrentDiet(fmt.Sprint(groupHash))
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, diet)
}

func ChangeDietName(c *gin.Context) {
	type RequestBody struct {
		DietHash string `json:"diet-hash"`
		Name     string `json:"name"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	if err := stores.ChangeDietName(requestBody.DietHash, requestBody.Name); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func GetGroupDiets(c *gin.Context) {
	session := sessions.Default(c)
	group := session.Get("group")

	diets, err := stores.GetGroupDiets(fmt.Sprint(group))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, diets)
}

func ChangeCurrentDiet(c *gin.Context) {
	type RequestBody struct {
		DietHash string `json:"diet-hash"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	session := sessions.Default(c)
	hash := session.Get("hash")
	group := session.Get("group")

	if err := stores.ChangeCurrentDiet(fmt.Sprint(hash), fmt.Sprint(group), requestBody.DietHash); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
