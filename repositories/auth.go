package repositories

// CreateAccessToken repository returns auth access Token
func CreateAccessToken(username string, password string) {

}

// CreateUser repository creates user
func CreateUser(username string, password string, fullname string) int64 {
	objID := InsertRow(
		"INSERT INTO Users(username, password, fullname) values(?,?,?)",
		username, password, fullname,
	)
	return objID
}
