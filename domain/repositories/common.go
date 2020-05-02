package repositories

import (
	"database/sql"
	"os"
)

// CheckErr raises error
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getConnection() *sql.DB {
	wd, err := os.Getwd()
	CheckErr(err)

	var databasePath string = wd + "/messenger.db"

	db, err := sql.Open("sqlite3", databasePath)
	CheckErr(err)

	return db
}

// InsertRow insert row to database
func InsertRow(query string, args ...interface{}) int64 {
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	CheckErr(err)

	res, err := stmt.Exec(args...)
	CheckErr(err)

	id, err := res.LastInsertId()
	CheckErr(err)

	return id
}

func fetchRows(query string, args ...interface{}) *sql.Rows {
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	CheckErr(err)

	rows, err := stmt.Query(args...)
	CheckErr(err)

	return rows
}
