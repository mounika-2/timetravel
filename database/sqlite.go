package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	createRecordsTable := `
	CREATE TABLE IF NOT EXISTS records (
		id INTEGER PRIMARY KEY,
		created_at DATETIME NOT NULL
	);
	`

	_, err = db.Exec(createRecordsTable)
	if err != nil {
		return nil, err
	}

	createVersionsTable := `
	CREATE TABLE IF NOT EXISTS record_versions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,

		record_id INTEGER NOT NULL,

		version INTEGER NOT NULL,

		data TEXT NOT NULL,

		created_at DATETIME NOT NULL,

		FOREIGN KEY(record_id) REFERENCES records(id)
	);
	`

	_, err = db.Exec(createVersionsTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}
