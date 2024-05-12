package handlers

import (
	"fmt"
	"net/http"
	"zakroma_backend/schemas"
	"zakroma_backend/stores"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetDietByHash(c *gin.Context) {
	hash := c.Params.ByName("hash")
	if len(hash) == 0 {
		c.String(http.StatusBadRequest, "something bad with field 'hash'")
		return
	}

	diet, err := stores.GetDietByHashWithoutDishes(hash)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, diet)
}

func GetDietByHashWithoutDishes(hash string) (schemas.Diet, error) {
	db, err := stores.CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return schemas.Diet{}, err
	}

	var diet schemas.Diet
	err = db.QueryRow(`
		select
			diet_id,
			diet_hash,
			diet_name
		from
			diet
		where
			diet_hash = $1`,
		hash).Scan(
		&diet.Id,
		&diet.Hash,
		&diet.Name)
	if err != nil {
		return schemas.Diet{}, err
	}

	fmt.Println(diet.Id, diet.Hash, diet.Name)

	dayDietsRows, err := db.
		Query(`
			select
			    diet_day_id,
			    index
			from
			    diet_day_diet
			where
			    diet_id = $1
			order by
			    index`,
			diet.Id)
	defer dayDietsRows.Close()
	if err != nil {
		return schemas.Diet{}, err
	}

	for dayDietsRows.Next() {
		var dayDietId int
		var index int
		if err = dayDietsRows.Scan(
			&dayDietId,
			&index); err != nil {
			return schemas.Diet{}, err
		}

		diet.DayDiets = append(diet.DayDiets, dayDietId)
	}

	return diet, nil
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
	user := session.Get("hash")

	hash, err := stores.CreateDiet(requestBody.Name, fmt.Sprint(groupHash))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := stores.ChangeCurrentDiet(fmt.Sprint(user), fmt.Sprint(groupHash), hash); err != nil {
		c.String(http.StatusNotFound, err.Error())
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

func GetGroupDiets(c *gin.Context) {
	session := sessions.Default(c)
	group := session.Get("group")

	diets, err := stores.GetGroupDiets(fmt.Sprint(group))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, diets)
}

func ChangeCurrentDiet(c *gin.Context) {
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

	if err := stores.ChangeCurrentDiet(fmt.Sprint(hash), fmt.Sprint(group), requestBody.DietHash); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func GetDietProducts(c *gin.Context) {
	type RequestBody struct {
		DietHash string `json:"diet-hash"`
		Days     []int  `json:"days"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	products, err := stores.GetDietProducts(requestBody.DietHash, requestBody.Days)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}
