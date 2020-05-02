package repositories

import (
	"errors"

	"github.com/bahattincinic/messenger-challenge/domain/models"
)

// Users is list of users
type Users []models.User

// FetchUsers Repository returns user list
func FetchUsers() Users {
	var users Users
	rows := fetchRows("SELECT id, username, fullname from Users")

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.FullName)
		CheckErr(err)
		users = append(users, user)
	}

	return users
}

// FetchUserByID repository returns user
func FetchUserByID(userID int64) (user models.User, err error) {
	row := fetchRows(
		"SELECT id, username, fullname FROM Users WHERE id = ?",
		userID,
	)

	if row.Next() {
		err = row.Scan(&user.ID, &user.Username, &user.FullName)
	} else {
		err = errors.New("Invalid credentials")
	}

	return
}

// FetchUserByUsername repository returns user
func FetchUserByUsername(username string) (user models.User, err error) {
	row := fetchRows(
		"SELECT id, username, fullname FROM Users WHERE username = ?",
		username,
	)

	if row.Next() {
		err = row.Scan(&user.ID, &user.Username, &user.FullName)
	} else {
		err = errors.New("Invalid credentials")
	}

	return
}
