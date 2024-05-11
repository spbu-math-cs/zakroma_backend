package stores

import (
	"zakroma_backend/schemas"
	"zakroma_backend/utils"
)

func GetMealIdByHash(hash string) (int, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return -1, err
	}

	var mealId int
	err = db.QueryRow(`
		select
			meal_id
		from
			meals
		where
			meal_hash = $1`,
		hash).Scan(
		&mealId)
	if err != nil {
		return -1, err
	}

	return mealId, nil
}

func GetMealByHash(hash string) (schemas.Meal, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return schemas.Meal{}, err
	}

	var meal schemas.Meal
	var tagId int
	err = db.QueryRow(`
		select
			meals.meal_id,
			meals.meal_hash,
			case 
			    when meals.meal_name is null
			        then ''
				else
				    meals.meal_name
			end,
			case
			    when meals.tag_id is null
			        then -1
				else
					meals.tag_id
			end
		from
			meals
		where
			meals.meal_hash = $1`,
		hash).Scan(
		&meal.Id,
		&meal.Hash,
		&meal.Name,
		&tagId)
	if err != nil {
		return schemas.Meal{}, err
	}
	if len(meal.Name) == 0 {
		var tag string
		err = db.QueryRow(`
			select
				tag
			from
				tags_for_meals
			where
				tag_id = $1`,
			tagId).Scan(
			&tag)
		if err != nil {
			return schemas.Meal{}, err
		}
		meal.Name = tag
	}

	dishesRows, err := db.Query(`
		select
			dishes.dish_hash
		from
			meals_dishes join dishes on meals_dishes.dish_id = dishes.dish_id
		where
			meals_dishes.meal_id = $1`,
		meal.Id)
	if err != nil {
		return schemas.Meal{}, err
	}
	defer dishesRows.Close()

	for dishesRows.Next() {
		var dishHash string
		if err = dishesRows.Scan(
			&dishHash); err != nil {
			return schemas.Meal{}, err
		}

		meal.Dishes = append(meal.Dishes, dishHash)
		//TODO (вроде сделал)
	}

	meal.DishesAmount = len(meal.Dishes)

	return meal, nil
}

func GetMealById(id int) (schemas.Meal, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return schemas.Meal{}, err
	}

	var meal schemas.Meal
	var tagId int
	err = db.QueryRow(`
		select
			meals.meal_id,
			meals.meal_hash,
			case 
			    when meals.meal_name is null
			        then ''
				else
				    meals.meal_name
			end,
			case
			    when meals.tag_id is null
			        then -1
				else
					meals.tag_id
			end
		from
			meals
		where
			meals.meal_id = $1`,
		id).Scan(
		&meal.Id,
		&meal.Hash,
		&meal.Name,
		&tagId)
	if err != nil {
		return schemas.Meal{}, err
	}
	if len(meal.Name) == 0 {
		var tag string
		err = db.QueryRow(`
			select
				tag
			from
				tags_for_meals
			where
				tag_id = $1`,
			tagId).Scan(
			&tag)
		if err != nil {
			return schemas.Meal{}, err
		}
		meal.Name = tag
	}

	dishesRows, err := db.Query(`
		select
			dishes.dish_hash
		from
			meals_dishes join dishes on meals_dishes.dish_id = dishes.dish_id
		where
			meals_dishes.meal_id = $1`,
		meal.Id)
	if err != nil {
		return schemas.Meal{}, err
	}
	defer dishesRows.Close()

	for dishesRows.Next() {
		var dishHash string
		if err = dishesRows.Scan(
			&dishHash); err != nil {
			return schemas.Meal{}, err
		}

		meal.Dishes = append(meal.Dishes, dishHash)
		//TODO (вроде сделал)
	}

	meal.DishesAmount = len(meal.Dishes)

	return meal, nil
}

var MEALS_CACHE_SIZE = 100000
var MealsCache = map[int]schemas.Meal{}
var MealsCacheQueue = utils.CreateQueue(MEALS_CACHE_SIZE)

func CreateMeal(dietHash string, dayDietIndex int, name string) (string, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return "", err
	}

	dietId, err := GetDietIdByHash(dietHash)
	if err != nil {
		return "", err
	}

	dayDietId, err := GetDayDietId(dietId, dayDietIndex)
	if err != nil {
		return "", err
	}

	mealIndex := -1
	maxMealRow, err := db.Query(`
		select
		    max(index)
		from
		    diet_day_meals
		where
		    diet_day_id = $1
		group by
		    diet_day_id`,
		dayDietId)
	if err != nil {
		return "", err
	}
	defer maxMealRow.Close()

	for maxMealRow.Next() {
		err = maxMealRow.Scan(&mealIndex)
		if err != nil {
			return "", err
		}
	}
	mealIndex++

	tags, err := GetAllMealsTags()
	if err != nil {
		return "", err
	}

	tagId := -1
	for i := range tags {
		if tags[i].Name == name {
			tagId = tags[i].TagId
		}
	}

	hash, err := utils.GenerateRandomHash(64)
	if err != nil {
		return "", err
	}

	mealId := 0
	if tagId == -1 {
		if err = db.QueryRow(`
			insert into
				meals(meal_name, meal_hash)
			values 
				($1, $2)
			returning
				meal_id`,
			name, hash).Scan(
			&mealId); err != nil {
			return "", err
		}
	} else {
		if err = db.QueryRow(`
		insert into
			meals(tag_id, meal_hash)
		values 
			($1, $2)
		returning
			meal_id`,
			tagId, hash).Scan(
			&mealId); err != nil {
			return "", err
		}
	}

	if err = db.QueryRow(`
		insert into
			diet_day_meals(diet_day_id, meal_id, index)
		values
			($1, $2, $3)
		returning
			meal_id`,
		dayDietId, mealId, mealIndex).Scan(
		&dayDietId); err != nil {
		return "", err
	}

	return hash, nil
}

func GetAllMealsTags() ([]schemas.Tag, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return []schemas.Tag{}, err
	}

	tagsRows, err := db.Query(`
		select
		    tag_id,
		    tag
		from
		    tags_for_meals`)
	if err != nil {
		return []schemas.Tag{}, err
	}
	defer tagsRows.Close()

	var tags []schemas.Tag
	for tagsRows.Next() {
		var tag schemas.Tag
		if err = tagsRows.Scan(
			&tag.TagId,
			&tag.Name); err != nil {
			return []schemas.Tag{}, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func AddMealDish(mealHash string, dishHash string, portions int) error {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return err
	}

	mealId, err := GetMealIdByHash(mealHash)
	if err != nil {
		return err
	}

	dishId, err := GetDishIdByHash(dishHash)
	if err != nil {
		return err
	}

	if err = db.QueryRow(`
		insert into
			meals_dishes(meal_id, dish_id, portions)
		values
			($1, $2, $3)
		returning
			meal_id`,
		mealId, dishId, portions).Scan(
		&mealId); err != nil {
		return err
	}

	return nil
}
