package handlers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"zakroma_backend/stores"
)

func GetGroupCartList(c *gin.Context) {
	session := sessions.Default(c)
	group := session.Get("group")

	cart, err := stores.GetGroupCartList(fmt.Sprint(group))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cart)
}

func AddGroupCartProduct(c *gin.Context) {
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
	if err := stores.AddGroupCartProduct(fmt.Sprint(group),
		requestBody.ProductId, requestBody.Amount); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func RemoveGroupCartProduct(c *gin.Context) {
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
	if err := stores.RemoveGroupCartProduct(fmt.Sprint(group),
		requestBody.ProductId); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func ChangeGroupCartProduct(c *gin.Context) {
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
	if err := stores.ChangeGroupCartProduct(fmt.Sprint(group),
		requestBody.ProductId, requestBody.Amount); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
