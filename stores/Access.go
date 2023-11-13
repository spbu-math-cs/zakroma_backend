package stores

func CheckDietAccessWithId(dietId int, userId int) bool {
	db, err := CreateConnection()
	if err != nil {
		return false
	}

	rows, err := db.
		Query(`
			select
			    users_groups.user_id,
			    users_groups.group_id,
			    groups_diets.diet_id
			from
			    users_groups,
			    groups_diets
			where
			    users_groups.user_id = $1 and
			    users_groups.group_id = groups_diets.group_id and
			    groups_diets.diet_id = $2`,
			userId, dietId)
	if err != nil {
		return false
	}

	defer rows.Close()

	for rows.Next() {
		var groupId int
		err = rows.Scan(&userId, &groupId, &dietId)
		if err != nil {
			return false
		}
		return true
	}

	return false
}

func CheckDayDietAccessWithId(dayDietId int, userId int) bool {
	db, err := CreateConnection()
	if err != nil {
		return false
	}

	rows, err := db.
		Query(`
			select
			    users_groups.user_id,
			    users_groups.group_id,
			    groups_diets.diet_id,
			    diet_day_diet.diet_day_id
			from
			    users_groups,
			    groups_diets,
			    diet_day_diet
			where
			    users_groups.user_id = $1 and
			    users_groups.group_id = groups_diets.group_id and
			    groups_diets.diet_id = diet_day_diet.diet_id and
			    diet_day_diet.diet_day_id = $2`,
			userId, dayDietId)
	if err != nil {
		return false
	}

	defer rows.Close()

	for rows.Next() {
		var groupId int
		var dietId int
		err = rows.Scan(&userId, &groupId, &dietId, &dayDietId)
		if err != nil {
			return false
		}
		return true
	}

	return false
}
