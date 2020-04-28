package repositories

import (
	"github.com/bahattincinic/messenger-challenge/models"
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
