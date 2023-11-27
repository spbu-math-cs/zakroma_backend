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

func CreateMeal(dietId int, dayDietIndex int, name string, dishes []int) (int, error) {
	db, err := CreateConnection()
	if err != nil {
		return -1, err
	}

	dayDietId := -1
	err = db.
		QueryRow(`
			select
				diet_day_id
			from
			    diet_day_diet
			where
			    diet_id = $1 and
			    index = $2`,
			dietId, dayDietIndex).
		Scan(&dayDietId)
	if err != nil {
		return -1, err
	}

	if dayDietId == -1 {
		dayDietId, err = CreateDayDiet(dietId, dayDietIndex)
		if err != nil {
			return -1, err
		}
	}

	mealIndex := 0
	rows, err := db.Query(`
		select
		    max(index)
		from
		    diet_day_meals
		where
		    diet_day_id = $1
		group by
		    diet_day_id
		`, dayDietId)

	for rows.Next() {
		err = rows.Scan(&mealIndex)
		if err != nil {
			return -1, err
		}
	}

	mealId := 0
	if err = db.QueryRow(`
		insert into
			meals(meal_name)
		values 
			($1)
		returning
			meals.meal_id
		`, name).Scan(&mealId); err != nil {
		return -1, err
	}

	for i := range dishes {
		dishId := dishes[i]
		rMealId := 0
		if err = db.QueryRow(`
			insert into
				meals_dishes(meal_id, dish_id, portions)
			values
				($1, $2, $3)
			returning
				meal_id
			`, mealId, dishId, 1).Scan(&rMealId); err != nil || rMealId != mealId {
			return -1, err
		}
	}

	return mealId, nil
}
