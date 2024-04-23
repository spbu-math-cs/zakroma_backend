package handlers

import (
	"fmt"
	"net/http"
	"zakroma_backend/stores"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetGroupStoreList godoc
//
// @Tags store
// @Accept json
// @Produce json
// @Success 200 {array} schemas.DishProduct
// @Security Bearer
// @Router /api/groups/store [get]
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

// AddGroupStoreProduct godoc
//
// @Tags store
// @Accept json
// @Produce json
// @Param data body handlers.AddGroupStoreProduct.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/store/add [post]
func AddGroupStoreProduct(c *gin.Context) {
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
	if err := stores.AddGroupStoreProduct(fmt.Sprint(group),
		requestBody.ProductId, requestBody.Amount); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// RemoveGroupStoreProduct godoc
//
// @Tags store
// @Accept json
// @Produce json
// @Param data body handlers.RemoveGroupStoreProduct.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/store/remove [post]
func RemoveGroupStoreProduct(c *gin.Context) {
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
	if err := stores.RemoveGroupStoreProduct(fmt.Sprint(group),
		requestBody.ProductId); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// ChangeGroupStoreProduct godoc
//
// @Tags store
// @Accept json
// @Produce json
// @Param data body handlers.ChangeGroupStoreProduct.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/store/change [patch]
func ChangeGroupStoreProduct(c *gin.Context) {
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
	if err := stores.ChangeGroupStoreProduct(fmt.Sprint(group),
		requestBody.ProductId, requestBody.Amount); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
