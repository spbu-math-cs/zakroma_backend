package stores

import (
	"fmt"
	"zakroma_backend/schemas"
	"zakroma_backend/utils"
)

func GetCurrentDietId(groupHash string) (int, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
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
	if err == nil {
		defer db.Close()
	}
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
	if err == nil {
		defer db.Close()
	}
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

func AddGroupDietByHash(userHash string, groupHash string, dietHash string) error {
	userRole, err := GetUserRole(userHash, groupHash)
	if err != nil {
		return err
	}

	if userRole != "Admin" {
		return fmt.Errorf("no permission")
	}

	dietId, err := GetDietIdByHash(dietHash)
	if err != nil {
		return err
	}

	return AddGroupDiet(groupHash, dietId)
}

func CreateGroup(name string, user string) (string, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
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
			groups(group_name, group_hash, current_diet_id)
		values
			($1, $2)
		returning
			group_id`,
		name,
		hash).Scan(
		&id); err != nil {
		return "", err
	}

	userId, err := GetUserIdByHash(user)
	if err != nil {
		return "", err
	}

	if err := db.QueryRow(`
		insert into
			users_groups(user_id, group_id, role)
		values
			($1, $2, $3)
		returning
			user_id`,
		userId,
		id,
		"Admin").Scan(
		&userId); err != nil {
		return "", err
	}

	return hash, nil
}

func CreatePersonalGroup(hash string) error {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
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

	userId, err := GetUserIdByHash(hash)
	if err != nil {
		return err
	}

	if err := db.QueryRow(`
		insert into
			users_groups(user_id, group_id, role)
		values
			($1, $2, $3)
		returning
			user_id`,
		userId,
		id,
		"Admin").Scan(
		&userId); err != nil {
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
	if err == nil {
		defer db.Close()
	}
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

func GetUserRole(userHash string, groupHash string) (string, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return "", err
	}

	userId, err := GetUserIdByHash(userHash)
	if err != nil {
		return "", err
	}

	groupId, err := GetGroupIdByHash(groupHash)
	if err != nil {
		return "", err
	}

	var role string
	if err = db.QueryRow(`
		select
		    role
		from
		    users_groups
		where
		    user_id = $1 and
		    group_id = $2`,
		userId,
		groupId).Scan(
		&role); err != nil {
		return "", err
	}

	return role, nil
}

func CheckUserGroup(userHash string, groupHash string) error {
	_, err := GetUserRole(userHash, groupHash)
	if err != nil {
		return err
	}

	return nil
}

func AddGroupUser(userHash string, groupHash string, newUserHash string, role string) error {
	userRole, err := GetUserRole(userHash, groupHash)
	if err != nil {
		return err
	}

	if userRole != "Admin" {
		return fmt.Errorf("no permission")
	}

	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return err
	}

	userId, err := GetUserIdByHash(newUserHash)
	if err != nil {
		return err
	}

	groupId, err := GetGroupIdByHash(groupHash)
	if err != nil {
		return err
	}

	if err := db.QueryRow(`
		insert into
			users_groups(user_id, group_id, role)
		values
			($1, $2, $3)
		returning
			user_id`,
		userId,
		groupId,
		role).Scan(
		&userId); err != nil {
		return err
	}

	return nil
}

func ChangeRole(userHash string, groupHash string, newUserHash string, role string) error {
	userRole, err := GetUserRole(userHash, groupHash)
	if err != nil {
		return err
	}

	if userRole != "Admin" {
		return fmt.Errorf("no permission")
	}

	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return err
	}

	userId, err := GetUserIdByHash(newUserHash)
	if err != nil {
		return err
	}

	groupId, err := GetGroupIdByHash(groupHash)
	if err != nil {
		return err
	}

	if err := db.QueryRow(`
		update
		    users_groups
		set
			role = $3
		where
		    user_id = $1 and
		    group_id = $2
		returning
			user_id`,
		userId,
		groupId,
		role).Scan(
		&userId); err != nil {
		return err
	}

	return nil
}

func MoveDietToCart(user string, group string, diet string, days []int) error {
	userRole, err := GetUserRole(user, group)
	if err != nil {
		return err
	}

	if userRole != "Admin" {
		return fmt.Errorf("no permission")
	}

	dietProducts, err := GetDietProducts(diet, days)
	if err != nil {
		return err
	}

	cart, err := GetGroupCartList(group)
	if err != nil {
		return err
	}

	i := 0
	j := 0
	for i < len(dietProducts) && j < len(cart) {
		if dietProducts[i].ProductId == cart[j].ProductId {
			if err := ChangeGroupCartProduct(group, dietProducts[i].ProductId,
				dietProducts[i].Amount+cart[j].Amount); err != nil {
				return err
			}
			i += 1
			j += 1
		} else if dietProducts[i].ProductId < cart[j].ProductId {
			if err := AddGroupCartProduct(group, dietProducts[i].ProductId,
				dietProducts[i].Amount); err != nil {
				return err
			}
			i += 1
		} else {
			j += 1
		}
	}
	for i < len(dietProducts) {
		if err := AddGroupCartProduct(group, dietProducts[i].ProductId,
			dietProducts[i].Amount); err != nil {
			return err
		}
		i += 1
	}

	return nil
}

func MoveCartToStore(user string, group string) error {
	userRole, err := GetUserRole(user, group)
	if err != nil {
		return err
	}

	if userRole != "Admin" {
		return fmt.Errorf("no permission")
	}

	cart, err := GetGroupCartList(group)
	if err != nil {
		return err
	}

	store, err := GetGroupStoreList(group)
	if err != nil {
		return err
	}

	i := 0
	j := 0
	for i < len(cart) && j < len(store) {
		if cart[i].ProductId == store[j].ProductId {
			if err := ChangeGroupStoreProduct(group, cart[i].ProductId, cart[i].Amount+store[j].Amount); err != nil {
				return err
			}
			if err := RemoveGroupCartProduct(group, cart[i].ProductId); err != nil {
				return err
			}
			i += 1
			j += 1
		} else if cart[i].ProductId < store[j].ProductId {
			if err := AddGroupStoreProduct(group, cart[i].ProductId, cart[i].Amount); err != nil {
				return err
			}
			if err := RemoveGroupCartProduct(group, cart[i].ProductId); err != nil {
				return err
			}
			i += 1
		} else {
			j += 1
		}
	}
	for i < len(cart) {
		if err := AddGroupStoreProduct(group, cart[i].ProductId, cart[i].Amount); err != nil {
			return err
		}
		if err := RemoveGroupCartProduct(group, cart[i].ProductId); err != nil {
			return err
		}
		i += 1
	}

	return nil
}
