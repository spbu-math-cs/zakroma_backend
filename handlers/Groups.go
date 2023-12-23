package handlers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"zakroma_backend/stores"
)

func CreateGroup(c *gin.Context) {
	type RequestBody struct {
		Name string `json:"name"`
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

func AddGroupUser(c *gin.Context) {
	type RequestBody struct {
		UserHash string `json:"user-hash"`
		Role     string `json:"role"`
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

func ChangeRole(c *gin.Context) {
	type RequestBody struct {
		UserHash string `json:"user-hash"`
		Role     string `json:"role"`
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

func AddGroupDiet(c *gin.Context) {
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

func ChangeCurrentGroup(c *gin.Context) {
	type RequestBody struct {
		GroupHash string `json:"group-hash"`
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
