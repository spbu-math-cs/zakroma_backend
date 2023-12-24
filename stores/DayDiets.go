package stores

import (
	"zakroma_backend/schemas"
)

func GetDayDietId(dietId int, index int) (int, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return -1, err
	}

	dayDietId := -1
	if err = db.QueryRow(`
		select
			diet_day_id
		from
			diet_day_diet
		where
			diet_id = $1 and
			index = $2`,
		dietId, index).Scan(
		&dayDietId); err != nil {
		return -1, err
	}

	return dayDietId, nil
}

func GetDayDietById(id int) (schemas.DayDiet, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return schemas.DayDiet{}, err
	}

	var dayDiet schemas.DayDiet
	dayDiet.Id = id
	if err = db.QueryRow(`
		select
			diet_day_name
		from
		    diet_day
		where
		    diet_day_id = $1`,
		id).Scan(
		&dayDiet.Name); err != nil {
		return schemas.DayDiet{}, err
	}

	mealsIdRows, err := db.Query(`
		select
			diet_day_meals.meal_id
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
	defer mealsIdRows.Close()

	for mealsIdRows.Next() {
		var mealId int
		if err = mealsIdRows.Scan(
			&mealId); err != nil {
			return schemas.DayDiet{}, err
		}
		meal, err := GetMealById(mealId)
		if err != nil {
			return schemas.DayDiet{}, err
		}
		dayDiet.Meals = append(dayDiet.Meals, meal)
	}

	return dayDiet, nil
}

func GetDayDietByIdWithoutDishes(id int) (schemas.DayDiet, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return schemas.DayDiet{}, err
	}

	var dayDiet schemas.DayDiet
	dayDiet.Id = id
	if err = db.QueryRow(`
		select
			diet_day_name
		from
		    diet_day
		where
		    diet_day_id = $1`,
		id).Scan(
		&dayDiet.Name); err != nil {
		return schemas.DayDiet{}, err
	}

	mealsIdRows, err := db.Query(`
		select
			diet_day_meals.meal_id
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
	defer mealsIdRows.Close()

	for mealsIdRows.Next() {
		var mealId int
		if err = mealsIdRows.Scan(
			&mealId); err != nil {
			return schemas.DayDiet{}, err
		}
		meal, err := GetMealByIdWithoutDishes(mealId)
		if err != nil {
			return schemas.DayDiet{}, err
		}
		dayDiet.Meals = append(dayDiet.Meals, meal)
	}

	return dayDiet, nil
}

func CreateDayDiet(dietId int, index int, name string) (int, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return -1, err
	}

	var dayDietId int
	if err = db.QueryRow(`
		insert into
			diet_day(diet_day_name)
		values 
			($1)
		returning
			diet_day_id`,
		name).Scan(
		&dayDietId); err != nil {
		return -1, err
	}

	if err = db.QueryRow(`
		insert into
			diet_day_diet(diet_id, diet_day_id, index)
		values
			($1, $2, $3)
		returning
			diet_day_id`,
		dietId, dayDietId, index).Scan(
		&dayDietId); err != nil {
		return -1, err
	}

	return dayDietId, nil
}
