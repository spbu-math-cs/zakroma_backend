package handlers

import (
	"fmt"
	"net/http"
	"zakroma_backend/stores"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetGroupCartList godoc
//
// @Tags cart
// @Accept json
// @Produce json
// @Success 200 {array} schemas.DishProduct
// @Security Bearer
// @Router /api/groups/cart [get]
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

// AddGroupCartProduct godoc
//
// @Tags cart
// @Accept json
// @Produce json
// @Param data body handlers.AddGroupCartProduct.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/cart/add [post]
func AddGroupCartProduct(c *gin.Context) {
	type RequestBody struct {
		ProductId int     `json:"product-id" example:"3"`
		Amount    float32 `json:"amount" example:"5"`
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

// RemoveGroupCartProduct godoc
//
// @Tags cart
// @Accept json
// @Produce json
// @Param data body handlers.RemoveGroupCartProduct.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/cart/remove [post]
func RemoveGroupCartProduct(c *gin.Context) {
	type RequestBody struct {
		ProductId int `json:"product-id" example:"3"`
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

// ChangeGroupCartProduct godoc
//
// @Tags cart
// @Accept json
// @Produce json
// @Param data body handlers.ChangeGroupCartProduct.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/cart/change [patch]
func ChangeGroupCartProduct(c *gin.Context) {
	type RequestBody struct {
		ProductId int     `json:"product-id" example:"3"`
		Amount    float32 `json:"amount" example:"5"`
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
