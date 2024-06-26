package stores

import (
	"fmt"
	"zakroma_backend/schemas"
	"zakroma_backend/utils"

	"golang.org/x/crypto/bcrypt"
)

func Login(user schemas.User) (string, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return "", err
	}

	var hashedPassword string
	if err = db.QueryRow(`
		select
			user_hash,
			password_hash
		from
			users
		where
			user_email = $1`,
		user.Email).Scan(
		&user.Hash,
		&hashedPassword); err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		return "", fmt.Errorf("wrong password")
	}

	return user.Hash, nil
}

func Register(user schemas.User) (string, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return "", err
	}

	usersRows, err := db.Query(`
		select
		    user_id
		from
			users
		where
			user_email = $1`,
		user.Email)
	if err != nil {
		return "", nil
	}
	defer usersRows.Close()

	user.Id = -1
	for usersRows.Next() {
		if err = usersRows.Scan(
			&user.Id); err != nil {
			return "", nil
		}
	}

	if user.Id != -1 {
		return "", fmt.Errorf("User with email = '%s' already exists", user.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user.Hash, err = utils.GenerateRandomHash(64)
	if err != nil {
		return "", err
	}

	if err = db.QueryRow(`
		insert into
			users(password_hash, user_name, user_surname, user_email, birth_date, user_hash)
		values
			($1, $2, $3, $4, CAST($5 as DATE), $6)
		returning
			user_id`,
		hashedPassword,
		user.Name,
		user.Surname,
		user.Email,
		user.BirthDate,
		user.Hash).Scan(
		&user.Id); err != nil {
		return "", err
	}

	if err = CreatePersonalGroup(user.Hash); err != nil {
		return "", err
	}

	return user.Hash, nil
}

func GetUserIdByHash(hash string) (int, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return -1, err
	}

	id := -1
	if err = db.QueryRow(`
		select
		    user_id
		from
		    users
		where
		    user_hash = $1`,
		hash).Scan(
		&id); err != nil {
		return -1, err
	}

	return id, nil
}

func GetUserInits(hash string) (string, string, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return "", "", err
	}

	var name, surname string
	if err = db.QueryRow(`
		select
			user_name, user_surname
		from
			users
		where
			user_hash = $1`,
		hash).Scan(
		&name, &surname); err != nil {
		return "", "", err
	}

	return name, surname, nil
}
