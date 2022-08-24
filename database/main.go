package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func connect() *sql.DB {
	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		println("Ocorreu um erro ao conectar no BD: ", err)
	}

	return db
}

func GetCount() int32 {
	var value int32
	db := connect()
	rows, err := db.Query("SELECT value FROM count limit 1")

	if err != nil {
		println("Ocorreu um erro: ", err)
	}

	db.Close()

	for rows.Next() {
		rows.Scan(&value)
	}

	return value
}

func SaveCount(value int32) {
	db := connect()
	delete, err := db.Prepare("DELETE FROM count")

	if err != nil {
		println("Ocorreu um erro: ", err)
	}

	delete.Exec()

	insert, err := db.Prepare("INSERT INTO count (value) values (?)")

	if err != nil {
		println("Ocorreu um erro: ", err)
	}

	insert.Exec(value)

	db.Close()
}

func CreateDatabaseIfNotExists() {
	_, err := os.Stat("database.db")

	if os.IsNotExist(err) {
		os.OpenFile("database.db", os.O_RDONLY|os.O_CREATE, 0666)
		db := connect()
		create, _ := db.Prepare("CREATE TABLE count (value INTEGER)")
		create.Exec()
		db.Close()
	}
}
