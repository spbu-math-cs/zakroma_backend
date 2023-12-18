package handlers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"zakroma_backend/stores"
)

func GetAllUserGroups(c *gin.Context) {
	session := sessions.Default(c)
	hash := session.Get("hash")

	groups, err := stores.GetAllUserGroups(fmt.Sprint(hash))
	if err != nil {

	}

	c.JSON(http.StatusOK, groups)
}
