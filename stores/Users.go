package stores

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"zakroma_backend/schemas"
)

func ValidateUser(username string, password string) (int, error) {
	db, err := CreateConnection()
	if err != nil {
		return -1, err
	}

	var user schemas.User
	err = db.QueryRow(`
		select
			user_id,
			user_name,
			password_hash
		from
			users
		where
			user_name = $1`,
		username).Scan(
		&user.Id,
		&user.Username,
		&user.Password)

	if err != nil {
		return -1, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return -1, fmt.Errorf("wrong password")
	}

	return user.Id, nil
}
