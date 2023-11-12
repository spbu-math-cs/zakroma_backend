package handlers

import (
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

	dish, err := stores.GetDietWithId(id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, dish)
}
