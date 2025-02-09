package database

import (
	"database/sql"
	"gofinn/config"
	"os"

	_ "github.com/ncruces/go-sqlite3/driver"

	_ "github.com/ncruces/go-sqlite3/embed"
)

var localDBInstance *sql.DB

func GetDatabaseInstance() *sql.DB {
	if localDBInstance == nil {
		dbFolder := config.Variables.DB_FOLDER
		os.MkdirAll(dbFolder, 0755)
		dbFileLocation := dbFolder + "/data.db"
		os.Create(dbFileLocation)

		db, err := sql.Open("sqlite3", "file:"+dbFileLocation+"?cache=shared&mode=rwc&_fk=1")
		if err != nil {
			panic("Failed to open database: " + dbFileLocation + "  |  " + err.Error())
		}
		localDBInstance = db
	}
	return localDBInstance
}
