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

	// main records table
	createRecordsTable := `
	CREATE TABLE IF NOT EXISTS records (
		id INTEGER PRIMARY KEY,
		data TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(createRecordsTable)
	if err != nil {
		return nil, err
	}

	// immutable snapshots
	createVersionsTable := `
	CREATE TABLE IF NOT EXISTS record_versions (
		record_id INTEGER NOT NULL,
		version INTEGER NOT NULL,
		data TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(record_id, version)
	);
	`

	_, err = db.Exec(createVersionsTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}
