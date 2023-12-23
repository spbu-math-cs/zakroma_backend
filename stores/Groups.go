package stores

import (
	"zakroma_backend/schemas"
	"zakroma_backend/utils"
)

func GetCurrentDietId(groupHash string) (int, error) {
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
		    group_hash = $1`,
		groupHash).Scan(
		&currentDietId); err != nil {
		return -1, err
	}

	return currentDietId, nil
}

func GetGroupIdByHash(groupHash string) (int, error) {
	db, err := CreateConnection()
	if err != nil {
		return -1, err
	}

	var id int
	if err = db.QueryRow(`
		select
		    group_id
		from
		    groups
		where
		    group_hash = $1`,
		groupHash).Scan(
		&id); err != nil {
		return -1, err
	}

	return id, nil
}

func AddGroupDiet(groupHash string, dietId int) error {
	db, err := CreateConnection()
	if err != nil {
		return err
	}

	groupId, err := GetGroupIdByHash(groupHash)
	if err != nil {
		return err
	}

	if err = db.QueryRow(`
		insert into
			groups_diets(group_id, diet_id)
		values
			($1, $2)
		returning
			group_id`,
		groupId,
		dietId).Scan(
		&groupId); err != nil {
		return err
	}

	return nil
}

func CreateGroup(name string) (string, error) {
	db, err := CreateConnection()
	if err != nil {
		return "", err
	}

	hash, err := utils.GenerateRandomHash(64)
	if err != nil {
		return "", err
	}

	var id int
	if err = db.QueryRow(`
		insert into
			groups(group_name, group_hash)
		values
			($1, $2)
		returning
			group_id`,
		name,
		hash).Scan(
		&id); err != nil {
		return "", err
	}

	return hash, nil
}

func CreatePersonalGroup(hash string) error {
	db, err := CreateConnection()
	if err != nil {
		return err
	}

	var id int
	if err = db.QueryRow(`
		insert into
			groups(group_name, group_hash)
		values
			($1, $2)
		returning
			group_id`,
		"Личная группа",
		hash).Scan(
		&id); err != nil {
		return err
	}

	return nil
}

func GetAllUserGroups(userHash string) ([]schemas.Group, error) {
	userId, err := GetUserIdByHash(userHash)
	if err != nil {
		return []schemas.Group{}, err
	}

	db, err := CreateConnection()
	if err != nil {
		return []schemas.Group{}, err
	}

	rows, err := db.Query(`
		select
		    users_groups.group_id,
		    groups.group_hash,
		    groups.group_name
		from
		    users_groups,
		    groups
		where
		    users_groups.user_id = $1 and
		    users_groups.group_id = groups.group_id`,
		userId)
	if err != nil {
		return []schemas.Group{}, err
	}
	defer rows.Close()

	var groups []schemas.Group
	for rows.Next() {
		var group schemas.Group
		if err = rows.Scan(
			&group.Id,
			&group.Hash,
			&group.Name); err != nil {
			return []schemas.Group{}, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}
