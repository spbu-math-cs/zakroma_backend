package stores

import (
	"zakroma_backend/schemas"
	"zakroma_backend/utils"
)

func GetMealIdByHash(hash string) (int, error) {
	db, err := CreateConnection()
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
	if err != nil {
		return schemas.Meal{}, err
	}

	var meal schemas.Meal
	err = db.QueryRow(`
		select
			meal_id,
			meal_hash,
			meal_name
		from
			meals
		where
			meal_hash = $1`,
		hash).Scan(
		&meal.Id,
		&meal.Hash,
		&meal.Name)
	if err != nil {
		return schemas.Meal{}, err
	}

	dishesRows, err := db.Query(`
		select
			meals_dishes.dish_id,
			meals_dishes.portions
		from
			meals_dishes
		where
			meals_dishes.meal_id = $1`,
		meal.Id)
	if err != nil {
		return schemas.Meal{}, err
	}
	defer dishesRows.Close()

	for dishesRows.Next() {
		var dishId int
		var portions float32
		if err = dishesRows.Scan(
			&dishId,
			&portions); err != nil {
			return schemas.Meal{}, err
		}

		var mealDish schemas.MealDish
		mealDish.Portions = portions
		mealDish.Dish, err = GetDishShortById(dishId)
		if err != nil {
			return schemas.Meal{}, err
		}

		meal.Dishes = append(meal.Dishes, mealDish)
	}

	meal.DishesAmount = len(meal.Dishes)

	return meal, nil
}

func GetMealById(id int) (schemas.Meal, error) {
	db, err := CreateConnection()
	if err != nil {
		return schemas.Meal{}, err
	}

	var meal schemas.Meal
	err = db.QueryRow(`
		select
			meal_id,
			meal_hash,
			meal_name
		from
			meals
		where
			meal_id = $1`,
		id).Scan(
		&meal.Id,
		&meal.Hash,
		&meal.Name)
	if err != nil {
		return schemas.Meal{}, err
	}

	dishesRows, err := db.Query(`
		select
			meals_dishes.dish_id,
			meals_dishes.portions
		from
			meals_dishes
		where
			meals_dishes.meal_id = $1`,
		meal.Id)
	if err != nil {
		return schemas.Meal{}, err
	}
	defer dishesRows.Close()

	for dishesRows.Next() {
		var dishId int
		var portions float32
		if err = dishesRows.Scan(
			&dishId,
			&portions); err != nil {
			return schemas.Meal{}, err
		}

		var mealDish schemas.MealDish
		mealDish.Portions = portions
		mealDish.Dish, err = GetDishShortById(dishId)
		if err != nil {
			return schemas.Meal{}, err
		}

		meal.Dishes = append(meal.Dishes, mealDish)
	}

	meal.DishesAmount = len(meal.Dishes)

	return meal, nil
}

func CreateMeal(dietHash string, dayDietIndex int, name string) (string, error) {
	db, err := CreateConnection()
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

	hash, err := utils.GenerateRandomHash(64)
	if err != nil {
		return "", err
	}

	mealId := 0
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

func AddMealDish(mealHash string, dishHash string, portions int) error {
	db, err := CreateConnection()
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
