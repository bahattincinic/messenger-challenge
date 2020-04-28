package repositories

import (
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/bahattincinic/messenger-challenge/models"
)

func tokenGenerate() string {
	b := make([]byte, 15)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// CreateAccessToken repository returns auth access Token
func CreateAccessToken(username string, password string) (accessToken string, err error) {
	row := fetchRows(
		"SELECT id, username, fullname from Users WHERE username = ? AND password = ?",
		username, password,
	)
	var user models.User

	if row.Next() {
		errRow := row.Scan(&user.ID, &user.Username, &user.FullName)
		CheckErr(errRow)

		accessToken = tokenGenerate()
		InsertRow(
			"INSERT INTO AccessTokens(access_token, user_id) VALUES (?,?)",
			accessToken, user.ID,
		)
	} else {
		err = errors.New("Invalid credentials")
	}
	return
}

// CreateUser repository creates user
func CreateUser(username string, password string, fullname string) int64 {
	objID := InsertRow(
		"INSERT INTO Users(username, password, fullname) values(?,?,?)",
		username, password, fullname,
	)
	return objID
}
