package stores

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"zakroma_backend/schemas"
)

func Login(user schemas.User) (int, error) {
	db, err := CreateConnection()
	if err != nil {
		return -1, err
	}

	var hashedPassword string
	if err = db.QueryRow(`
		select
			user_id,
			password_hash
		from
			users
		where
			user_email = $1`,
		user.Email).Scan(
		&user.Id,
		&hashedPassword); err != nil {
		return -1, nil
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		return -1, fmt.Errorf("wrong password")
	}

	return user.Id, nil
}

func Register(user schemas.User) (int, error) {
	db, err := CreateConnection()
	if err != nil {
		return -1, err
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
		return -1, nil
	}
	defer usersRows.Close()

	user.Id = -1
	for usersRows.Next() {
		if err = usersRows.Scan(
			&user.Id); err != nil {
			return -1, nil
		}
	}

	if user.Id != -1 {
		return -1, fmt.Errorf("User with email = '%s' already exists", user.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	if err = db.QueryRow(`
		insert into
			users(password_hash, user_name, user_surname, user_email, birth_date)
		values
			($1, $2, $3, $4, CAST($5 as DATE))
		returning
			user_id`,
		hashedPassword,
		user.Name,
		user.Surname,
		user.Email,
		user.BirthDate).Scan(
		&user.Id); err != nil {
		return -1, err
	}

	return user.Id, nil
}
