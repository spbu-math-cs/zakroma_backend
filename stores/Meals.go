package stores

import "zakroma_backend/schemas"

func GetMealWithId(id int) (schemas.Meal, error) {
	db, err := CreateConnection()
	if err != nil {
		return schemas.Meal{}, err
	}

	var meal schemas.Meal
	err = db.
		QueryRow(`
			select
			    meal_id,
			    meal_name
			from
			    meals
			where
			    meal_id = $1`,
			id).
		Scan(&meal.Id,
			&meal.Name)

	if err != nil {
		return schemas.Meal{}, err
	}

	dishesRows, err := db.
		Query(`
			select
			    meals_dishes.dish_id,
			    meals_dishes.portions
			from
			    meals_dishes
			where
			    meals_dishes.meal_id = $1`,
			id)

	if err != nil {
		return schemas.Meal{}, err
	}

	defer dishesRows.Close()

	for dishesRows.Next() {
		var dishId int
		var portions float32
		err = dishesRows.Scan(&dishId, &portions)
		if err != nil {
			return schemas.Meal{}, err
		}
		var mealDish schemas.MealDish
		mealDish.Portions = portions
		mealDish.Dish, err = GetDishShortWithId(dishId)
		if err != nil {
			return schemas.Meal{}, err
		}
		meal.Dishes = append(meal.Dishes, mealDish)
	}

	return meal, nil
}
