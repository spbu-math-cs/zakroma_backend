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
