package handlers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"zakroma_backend/stores"
)

func GetGroupStoreList(c *gin.Context) {
	session := sessions.Default(c)
	group := session.Get("group")

	cart, err := stores.GetGroupStoreList(fmt.Sprint(group))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cart)
}

func AddGroupStoreProduct(c *gin.Context) {
	type RequestBody struct {
		ProductId int     `json:"product-id"`
		Amount    float32 `json:"amount"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	session := sessions.Default(c)
	group := session.Get("group")
	if err := stores.AddGroupStoreProduct(fmt.Sprint(group),
		requestBody.ProductId, requestBody.Amount); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func RemoveGroupStoreProduct(c *gin.Context) {
	type RequestBody struct {
		ProductId int `json:"product-id"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	session := sessions.Default(c)
	group := session.Get("group")
	if err := stores.RemoveGroupStoreProduct(fmt.Sprint(group),
		requestBody.ProductId); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func ChangeGroupStoreProduct(c *gin.Context) {
	type RequestBody struct {
		ProductId int     `json:"product-id"`
		Amount    float32 `json:"amount"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	session := sessions.Default(c)
	group := session.Get("group")
	if err := stores.ChangeGroupStoreProduct(fmt.Sprint(group),
		requestBody.ProductId, requestBody.Amount); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
