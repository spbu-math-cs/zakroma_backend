package stores

import (
	"zakroma_backend/schemas"
	"zakroma_backend/utils"
)

func GetDietIdByHash(hash string) (int, error) {
	db, err := CreateConnection()
	if err != nil {
		return -1, err
	}

	var dietId int
	err = db.QueryRow(`
		select
			diet_id
		from
			diet
		where
			hash = $1`,
		hash).Scan(
		&dietId)
	if err != nil {
		return -1, err
	}

	return dietId, nil
}

func GetDietByHash(hash string) (schemas.Diet, error) {
	db, err := CreateConnection()
	if err != nil {
		return schemas.Diet{}, err
	}

	var diet schemas.Diet
	err = db.QueryRow(`
		select
			diet_id,
			hash,
			diet_name
		from
			diet
		where
			hash = $1`,
		hash).Scan(
		&diet.Id,
		&diet.Hash,
		&diet.Name)
	if err != nil {
		return schemas.Diet{}, err
	}

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
		dayDiet, err := GetDayDietById(dayDietId)
		if err != nil {
			return schemas.Diet{}, err
		}

		dayDiet.Index = index
		dayDiet.MealsAmount = len(dayDiet.Meals)
		dayDiet.Meals = dayDiet.Meals[:min(3, len(dayDiet.Meals))]

		diet.DayDiets = append(diet.DayDiets, dayDiet)
	}

	return diet, nil
}

func GetDietById(id int) (schemas.Diet, error) {
	db, err := CreateConnection()
	if err != nil {
		return schemas.Diet{}, err
	}

	var diet schemas.Diet
	err = db.QueryRow(`
		select
			diet_id,
			hash,
			diet_name
		from
			diet
		where
			diet_id = $1`,
		id).Scan(
		&diet.Id,
		&diet.Hash,
		&diet.Name)
	if err != nil {
		return schemas.Diet{}, err
	}

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
		dayDiet, err := GetDayDietById(dayDietId)
		if err != nil {
			return schemas.Diet{}, err
		}

		dayDiet.Index = index
		dayDiet.MealsAmount = len(dayDiet.Meals)
		dayDiet.Meals = dayDiet.Meals[:min(3, len(dayDiet.Meals))]

		diet.DayDiets = append(diet.DayDiets, dayDiet)
	}

	return diet, nil
}

var DefaultDayDietName = [7]string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота", "Воскресенье"}

func CreateDiet(name string) (string, error) {
	db, err := CreateConnection()
	if err != nil {
		return "", err
	}

	hash, err := utils.GenerateRandomHash(64)
	if err != nil {
		return "", err
	}

	id := -1
	if err = db.QueryRow(`
		insert into
			diet(diet_name, hash)
		values
			($1, $2)
		returning
		    diet_id`,
		name, hash).Scan(
		&id); err != nil {
		return "", err
	}

	for index := 0; index < 7; index++ {
		_, err := CreateDayDiet(id, index, DefaultDayDietName[index])
		if err != nil {
			return "", nil
		}
	}

	return hash, nil
}

func GetCurrentDiet(groupId int) (schemas.Diet, error) {
	dietId, err := GetCurrentDietId(groupId)
	if err != nil {
		return schemas.Diet{}, err
	}

	diet, err := GetDietById(dietId)
	if err != nil {
		return schemas.Diet{}, err
	}

	// TODO: get current day?
	currentDay := 3

	diet.DayDiets = diet.DayDiets[currentDay : currentDay+1]
	diet.DayDiets[0].Meals = []schemas.Meal{}

	return diet, nil
}

func ChangeDietName(dietHash string, name string) error {
	db, err := CreateConnection()
	if err != nil {
		return err
	}

	if err := db.QueryRow(`
		update
			diet
		set
		    diet_name = $2
		where
		    hash = $1
		returning
			hash`,
		dietHash,
		name).Scan(
		&dietHash); err != nil {
		return err
	}

	return nil
}
