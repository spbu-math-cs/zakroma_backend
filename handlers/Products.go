package handlers

import (
	"net/http"
	"strconv"
	"zakroma_backend/stores"

	"github.com/gin-gonic/gin"
)

// GetProductById godoc
//
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "id продукта"
// @Success 200 {object} schemas.Product
// @Router /api/products/{id} [get]
func GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	product, err := stores.GetProductById(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}
