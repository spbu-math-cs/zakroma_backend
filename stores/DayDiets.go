package stores

import "zakroma_backend/schemas"

func GetDayDietWithId(id int) (schemas.DayDiet, error) {
	db, err := CreateConnection()
	if err != nil {
		return schemas.DayDiet{}, err
	}

	mealsRows, err := db.
		Query(`
			select
				diet_day_meals.meal_id,
				meals.meal_name,
				index
			from
			    diet_day_meals,
			    meals
			where
			    diet_day_id = $1 and
			    diet_day_meals.meal_id = meals.meal_id
			order by
			    index`,
			id)
	if err != nil {
		return schemas.DayDiet{}, err
	}

	defer mealsRows.Close()

	var dayDiet schemas.DayDiet
	dayDiet.Id = id
	for mealsRows.Next() {
		var meal schemas.Meal
		if err = mealsRows.Scan(&meal.Id, &meal.Name, &meal.Index); err != nil {
			return schemas.DayDiet{}, err
		}
		dayDiet.Meals = append(dayDiet.Meals, meal)
	}

	return dayDiet, nil
}
