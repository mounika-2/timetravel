package service

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/rainbowmga/timetravel/entity"
)

type SQLiteRecordService struct {
	db *sql.DB
}

func NewSQLiteRecordService(db *sql.DB) SQLiteRecordService {
	return SQLiteRecordService{
		db: db,
	}
}

func (s *SQLiteRecordService) GetRecord(
	ctx context.Context,
	id int,
) (entity.Record, error) {

	query := `
	SELECT data
	FROM records
	WHERE id = ?
	`

	var rawJSON string

	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(&rawJSON)

	if err == sql.ErrNoRows {
		return entity.Record{}, ErrRecordDoesNotExist
	}

	if err != nil {
		return entity.Record{}, err
	}

	data := map[string]string{}

	err = json.Unmarshal([]byte(rawJSON), &data)
	if err != nil {
		return entity.Record{}, err
	}

	return entity.Record{
		ID:   id,
		Data: data,
	}, nil
}

func (s *SQLiteRecordService) CreateRecord(
	ctx context.Context,
	record entity.Record,
) error {

	rawJSON, err := json.Marshal(record.Data)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO records(id, data)
	VALUES(?, ?)
	`

	_, err = s.db.ExecContext(
		ctx,
		query,
		record.ID,
		string(rawJSON),
	)

	if err != nil {
		return err
	}

	return s.createVersionSnapshot(ctx, record)
}

func (s *SQLiteRecordService) UpdateRecord(
	ctx context.Context,
	id int,
	updates map[string]*string,
) (entity.Record, error) {

	record, err := s.GetRecord(ctx, id)
	if err != nil {
		return entity.Record{}, err
	}

	newRecord := record.Copy()

	for key, value := range updates {

		if value == nil {
			delete(newRecord.Data, key)
			continue
		}

		newRecord.Data[key] = *value
	}

	rawJSON, err := json.Marshal(newRecord.Data)
	if err != nil {
		return entity.Record{}, err
	}

	query := `
	UPDATE records
	SET data = ?
	WHERE id = ?
	`

	_, err = s.db.ExecContext(
		ctx,
		query,
		string(rawJSON),
		id,
	)

	if err != nil {
		return entity.Record{}, err
	}

	err = s.createVersionSnapshot(ctx, newRecord)
	if err != nil {
		return entity.Record{}, err
	}

	return newRecord, nil
}

func (s *SQLiteRecordService) createVersionSnapshot(
	ctx context.Context,
	record entity.Record,
) error {

	query := `
	SELECT COALESCE(MAX(version), 0) + 1
	FROM record_versions
	WHERE record_id = ?
	`

	var nextVersion int

	err := s.db.QueryRowContext(
		ctx,
		query,
		record.ID,
	).Scan(&nextVersion)

	if err != nil {
		return err
	}

	rawJSON, err := json.Marshal(record.Data)
	if err != nil {
		return err
	}

	insertQuery := `
	INSERT INTO record_versions(
		record_id,
		version,
		data,
		created_at
	)
	VALUES(?, ?, ?, CURRENT_TIMESTAMP)
	`

	_, err = s.db.ExecContext(
		ctx,
		insertQuery,
		record.ID,
		nextVersion,
		string(rawJSON),
	)

	return err
}

func (s *SQLiteRecordService) GetRecordVersion(
	ctx context.Context,
	id int,
	version int,
) (entity.Record, error) {

	query := `
	SELECT data
	FROM record_versions
	WHERE record_id = ?
	AND version = ?
	`

	var rawJSON string

	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
		version,
	).Scan(&rawJSON)

	if err == sql.ErrNoRows {
		return entity.Record{}, ErrRecordDoesNotExist
	}

	if err != nil {
		return entity.Record{}, err
	}

	data := map[string]string{}

	err = json.Unmarshal([]byte(rawJSON), &data)
	if err != nil {
		return entity.Record{}, err
	}

	return entity.Record{
		ID:   id,
		Data: data,
	}, nil
}

func (s *SQLiteRecordService) ListRecordVersions(
	ctx context.Context,
	id int,
) ([]entity.RecordVersion, error) {

	query := `
	SELECT version, data, created_at
	FROM record_versions
	WHERE record_id = ?
	ORDER BY created_at DESC
	`

	rows, err := s.db.QueryContext(
		ctx,
		query,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	versions := []entity.RecordVersion{}

	for rows.Next() {

		var version entity.RecordVersion
		var rawJSON string

		err := rows.Scan(
			&version.Version,
			&rawJSON,
			&version.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		data := map[string]string{}

		err = json.Unmarshal([]byte(rawJSON), &data)
		if err != nil {
			return nil, err
		}

		version.Data = data

		versions = append(versions, version)
	}

	return versions, nil
}
