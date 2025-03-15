package database

import (
	"database/sql"
	"go-chi-sqlite-jwt-starter/config"
	"log"
	"os"

	_ "github.com/ncruces/go-sqlite3/driver"

	_ "github.com/ncruces/go-sqlite3/embed"
)

var localDBInstance *sql.DB

func GetDatabaseInstance() *sql.DB {
	return localDBInstance
}

func Initialize() {
	log.Printf("Initializing database...")
	defer log.Printf("Database initialized.")

	dbFolder := config.Variables.DB_FOLDER
	os.MkdirAll(dbFolder, 0755)
	dbFileLocation := dbFolder + "/data.db"

	if _, err := os.Stat(dbFileLocation); os.IsNotExist(err) {
		file, err := os.Create(dbFileLocation)
		if err != nil {
			panic("Failed to create database file: " + err.Error())
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", "file:"+dbFileLocation+"?cache=shared&mode=rwc&_fk=1")
	if err != nil {
		panic("Failed to open database: " + dbFileLocation + "  |  " + err.Error())
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	pingErr := db.Ping()
	if pingErr != nil {
		panic("Failed to ping database: " + pingErr.Error())
	}

	localDBInstance = db

	InitializeTables()
	InitializeTriggers()
}

func InitializeTables() {
	for _, table := range Tables {
		_, err := localDBInstance.Exec(table)
		if err != nil {
			panic("Failed to create table: " + err.Error())
		}
	}
}

func InitializeTriggers() {
	for _, trigger := range Triggers {
		_, err := localDBInstance.Exec(trigger)
		if err != nil {
			panic("Failed to create trigger: " + err.Error())
		}
	}
}
