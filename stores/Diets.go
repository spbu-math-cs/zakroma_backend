package stores

import (
	"fmt"
	"zakroma_backend/schemas"
	"zakroma_backend/utils"
)

func GetDietHashById(id int) (string, error) {
	db, err := CreateConnection()
	if err != nil {
		return "", err
	}

	var dietHash string
	err = db.QueryRow(`
		select
			diet_hash
		from
			diet
		where
			diet_id = $1`,
		id).Scan(
		&dietHash)
	if err != nil {
		return "", err
	}

	return dietHash, nil
}

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
			diet_hash = $1`,
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
			diet_hash,
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

func CreateDiet(name string, groupHash string) (string, error) {
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
			diet(diet_name, diet_hash)
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
			return "", err
		}
	}

	if err = AddGroupDiet(groupHash, id); err != nil {
		return "", err
	}

	return hash, nil
}

func GetCurrentDiet(groupHash string) (schemas.Diet, error) {
	dietId, err := GetCurrentDietId(groupHash)
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
		    diet_hash = $1
		returning
			diet_hash`,
		dietHash,
		name).Scan(
		&dietHash); err != nil {
		return err
	}

	return nil
}

func GetGroupDiets(groupHash string) ([]string, error) {
	db, err := CreateConnection()
	if err != nil {
		return []string{}, err
	}

	groupId, err := GetGroupIdByHash(groupHash)
	if err != nil {
		return []string{}, err
	}

	dietsRows, err := db.Query(`
		select
		    diet_id
		from
		    groups_diets
		where
		    group_id = $1`,
		groupId)
	if err != nil {
		return []string{}, err
	}
	defer dietsRows.Close()

	var dietsIds []int
	for dietsRows.Next() {
		var dietId int
		if err = dietsRows.Scan(&dietId); err != nil {
			return []string{}, err
		}
		dietsIds = append(dietsIds, dietId)
	}

	var dietsHashes []string
	for i := range dietsIds {
		dietHash, err := GetDietHashById(dietsIds[i])
		if err != nil {
			return []string{}, err
		}
		dietsHashes = append(dietsHashes, dietHash)
	}

	return dietsHashes, nil
}

func ChangeCurrentDiet(userHash string, groupHash string, dietHash string) error {
	userRole, err := GetUserRole(userHash, groupHash)
	if err != nil {
		return err
	}

	if userRole != "Admin" {
		return fmt.Errorf("no permission")
	}

	db, err := CreateConnection()
	if err != nil {
		return err
	}

	dietId, err := GetDietIdByHash(dietHash)
	if err != nil {
		return err
	}

	if err = db.QueryRow(`
		update
			groups
		set
		    current_diet_id = $1
		where
		    group_hash = $2
		returning
			current_diet_id`,
		dietId,
		groupHash).Scan(
		&dietId); err != nil {
		return err
	}

	return nil
}
