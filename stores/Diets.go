package stores

import "zakroma_backend/schemas"

func GetDietWithId(id int) (schemas.Diet, error) {
	db, err := CreateConnection()
	if err != nil {
		return schemas.Diet{}, err
	}

	var diet schemas.Diet
	err = db.
		QueryRow(`
			select
				diet_id,
				diet_name
			from
			    diet
			where
			    diet_id = $1`,
			id).
		Scan(&diet.Id,
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
			id)
	if err != nil {
		return schemas.Diet{}, err
	}

	defer dayDietsRows.Close()

	for dayDietsRows.Next() {
		var dayDietId int
		var index int
		if err = dayDietsRows.Scan(&dayDietId, &index); err != nil {
			return schemas.Diet{}, err
		}
		diet.DayDiets = append(diet.DayDiets, schemas.DayDiet{Id: dayDietId, Index: index})
	}

	return diet, nil
}

func GetDietWithHash(hash int) (schemas.Diet, error) {
	db, err := CreateConnection()
	if err != nil {
		return schemas.Diet{}, err
	}

	var diet schemas.Diet
	err = db.
		QueryRow(`
			select
			    diet_id,
				hash,
				diet_name
			from
			    diet
			where
			    hash = $1`,
			hash).
		Scan(&diet.Id,
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
	if err != nil {
		return schemas.Diet{}, err
	}

	defer dayDietsRows.Close()

	for dayDietsRows.Next() {
		var dayDietId int
		var index int
		if err = dayDietsRows.Scan(&dayDietId, &index); err != nil {
			return schemas.Diet{}, err
		}
		diet.DayDiets = append(diet.DayDiets, schemas.DayDiet{Id: dayDietId, Index: index})
	}

	return diet, nil
}

func CreateDiet(name string) (int, error) {
	db, err := CreateConnection()
	if err != nil {
		return -1, err
	}

	id := -1
	if err = db.QueryRow(`
		insert into
			diet(diet_name)
		values
			($1)
		returning diet_id
		`, name).Scan(&id); err != nil {
		return -1, err
	}

	for index := 0; index < 7; index++ {
		dietDayId := -1
		if err = db.QueryRow(`
			insert into
				diet_day(diet_day_name)
			values
				('')
			returning diet_day_id
			`).Scan(&dietDayId); err != nil {
			return -1, nil
		}

		if err = db.QueryRow(`
			insert into
				diet_day_diet(diet_id, diet_day_id, index)
			values
				($1, $2, $3)
			returning diet_day_id
			`, id, dietDayId, index).Scan(&dietDayId); err != nil {
			return -1, nil
		}
	}

	return id, nil
}
