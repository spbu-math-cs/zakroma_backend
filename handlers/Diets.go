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

	hash, err := stores.CreateDiet(requestBody.Name, fmt.Sprint(groupHash))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
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
