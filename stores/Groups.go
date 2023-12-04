package stores

func GetCurrentDietId(groupId int) (int, error) {
	db, err := CreateConnection()
	if err != nil {
		return -1, err
	}

	var currentDietId int
	if err = db.QueryRow(`
		select
		    current_diet_id
		from
		    groups
		where
		    group_id = $1`,
		groupId).Scan(
		&currentDietId); err != nil {
		return -1, err
	}

	return currentDietId, nil
}
