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

func CreateDayDiet(dietId int, index int, name string) (int, error) {
	db, err := CreateConnection()
	if err != nil {
		return -1, err
	}

	dayDietId := -1

	rows, err := db.Query(`
		select
		    diet_day_id
		from
		    diet_day_diet
		where
		    diet_id = $1 and index = $2
		limit 1
		`, dietId, index)

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&dayDietId); err != nil {
			return -1, err
		}
	}

	if dayDietId == -1 {
		if err = db.QueryRow(`
			insert into
				diet_day(diet_day_name)
			values 
				($1)
			returning diet_day_id
			`, name).Scan(&dayDietId); err != nil {
			return -1, err
		}
		if err = db.QueryRow(`
			insert into
				diet_day_diet(diet_id, diet_day_id, index)
			values
				($1, $2, $3)
			returning diet_day_id
			`, dietId, dayDietId, index).Scan(&dayDietId); err != nil {
			return -1, err
		}
	}

	return dayDietId, nil
}
