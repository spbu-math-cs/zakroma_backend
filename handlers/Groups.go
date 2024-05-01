package handlers

import (
	"fmt"
	"net/http"
	"zakroma_backend/stores"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CreateGroup godoc
//
// @Tags groups
// @Accept json
// @Produce json
// @Param data body handlers.CreateGroup.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/create [post]
func CreateGroup(c *gin.Context) {
	type RequestBody struct {
		Name string `json:"name" example:"группа"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	session := sessions.Default(c)
	hash := session.Get("hash")

	groupHash, err := stores.CreateGroup(requestBody.Name, fmt.Sprint(hash))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	session.Set("group", groupHash)
	if err = session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// AddGroupUser godoc
//
// @Tags groups
// @Accept json
// @Produce json
// @Param data body handlers.AddGroupUser.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/user/add [post]
func AddGroupUser(c *gin.Context) {
	type RequestBody struct {
		UserHash string `json:"user-hash" example:"0183ae4f4083f289ff1c58c320b663764c2f129b0cd01d9e78a34bb33dba248f"`
		Role     string `json:"role" example:"Admin"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	session := sessions.Default(c)
	hash := session.Get("hash")
	group := session.Get("group")

	if err := stores.AddGroupUser(fmt.Sprint(hash), fmt.Sprint(group), requestBody.UserHash, requestBody.Role); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// ChangeRole godoc
//
// @Tags groups
// @Accept json
// @Produce json
// @Param data body handlers.ChangeRole.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/role [patch]
func ChangeRole(c *gin.Context) {
	type RequestBody struct {
		UserHash string `json:"user-hash" example:"0183ae4f4083f289ff1c58c320b663764c2f129b0cd01d9e78a34bb33dba248f"`
		Role     string `json:"role" example:"User"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	session := sessions.Default(c)
	hash := session.Get("hash")
	group := session.Get("group")

	if err := stores.ChangeRole(fmt.Sprint(hash), fmt.Sprint(group), requestBody.UserHash, requestBody.Role); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// AddGroupDiet godoc
//
// @Tags groups
// @Accept json
// @Produce json
// @Param data body handlers.AddGroupDiet.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/diet/add [post]
func AddGroupDiet(c *gin.Context) {
	type RequestBody struct {
		DietHash string `json:"diet-hash" example:"a337b13aaa3d71ffb24707d9f73d3f5ad6bcb7388da5a35618965aa0dbd18aab"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	session := sessions.Default(c)
	hash := session.Get("hash")
	group := session.Get("group")

	if err := stores.AddGroupDietByHash(fmt.Sprint(hash), fmt.Sprint(group), requestBody.DietHash); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := stores.ChangeCurrentDiet(fmt.Sprint(hash), fmt.Sprint(group), requestBody.DietHash); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// ChangeCurrentGroup godoc
//
// @Tags groups
// @Accept json
// @Produce json
// @Param data body handlers.ChangeCurrentGroup.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/change [patch]
func ChangeCurrentGroup(c *gin.Context) {
	type RequestBody struct {
		GroupHash string `json:"group-hash" example:"bb664cca44adce86a19b1c5f3f6b42b37356eebfd4e937773d0ac93a9c56c6ca"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	session := sessions.Default(c)
	hash := session.Get("hash")

	if err := stores.CheckUserGroup(fmt.Sprint(hash), requestBody.GroupHash); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	session.Set("group", requestBody.GroupHash)
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// GetAllUserGroups godoc
//
// @Tags groups
// @Accept json
// @Produce json
// @Success 200 {array} schemas.Group
// @Security Bearer
// @Router /api/groups/list [get]
func GetAllUserGroups(c *gin.Context) {
	session := sessions.Default(c)
	hash := session.Get("hash")

	groups, err := stores.GetAllUserGroups(fmt.Sprint(hash))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, groups)
}

// MoveDietToCart godoc
//
// @Tags groups
// @Accept json
// @Produce json
// @Param data body handlers.MoveDietToCart.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/groups/move/diet/cart [patch]
func MoveDietToCart(c *gin.Context) {
	type RequestBody struct {
		DietHash string `json:"diet-hash" example:"a337b13aaa3d71ffb24707d9f73d3f5ad6bcb7388da5a35618965aa0dbd18aab"`
		Days     []int  `json:"days"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	session := sessions.Default(c)
	user := session.Get("hash")
	group := session.Get("group")

	if err := stores.MoveDietToCart(fmt.Sprint(user), fmt.Sprint(group),
		requestBody.DietHash, requestBody.Days); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// MoveCartToStore godoc
//
// @Tags groups
// @Accept json
// @Produce json
// @Success 200
// @Security Bearer
// @Router /api/groups/move/cart/store [patch]
func MoveCartToStore(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("hash")
	group := session.Get("group")

	if err := stores.MoveCartToStore(fmt.Sprint(user), fmt.Sprint(group)); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// GetGroupMembers godoc
//
// @Tags groups
// @Accept json
// @Produce json
// @Success 200 {array} schemas.User
// @Security Bearer
// @Router /api/groups/members [get]
func GetGroupMembers(c *gin.Context) {
	session := sessions.Default(c)
	groupHash := session.Get("group")

	if groupHash == nil {
		c.String(http.StatusBadRequest, "no current group set")
		return
	}

	members, err := stores.GetGroupMembers(fmt.Sprint(groupHash))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, members)
}
