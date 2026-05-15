package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

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
	FROM record_versions
	WHERE record_id = ?
	ORDER BY version DESC
	LIMIT 1
	`

	var rawJSON string

	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(&rawJSON)

	if errors.Is(err, sql.ErrNoRows) {
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

func (s *SQLiteRecordService) GetRecordVersion(
	ctx context.Context,
	id int,
	version int,
) (entity.RecordVersion, error) {

	query := `
	SELECT data, created_at
	FROM record_versions
	WHERE record_id = ?
	AND version = ?
	`

	var rawJSON string
	var createdAt time.Time

	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
		version,
	).Scan(&rawJSON, &createdAt)

	if errors.Is(err, sql.ErrNoRows) {
		return entity.RecordVersion{}, ErrRecordDoesNotExist
	}

	if err != nil {
		return entity.RecordVersion{}, err
	}

	data := map[string]string{}

	err = json.Unmarshal([]byte(rawJSON), &data)
	if err != nil {
		return entity.RecordVersion{}, err
	}

	return entity.RecordVersion{
		RecordID:  id,
		Version:   version,
		Data:      data,
		CreatedAt: createdAt,
	}, nil
}

func (s *SQLiteRecordService) CreateRecord(
	ctx context.Context,
	record entity.Record,
) error {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	now := time.Now().UTC()

	_, err = tx.ExecContext(
		ctx,
		`
		INSERT INTO records(id, created_at)
		VALUES(?, ?)
		`,
		record.ID,
		now,
	)

	if err != nil {
		return err
	}

	rawJSON, err := json.Marshal(record.Data)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		`
		INSERT INTO record_versions(
			record_id,
			version,
			data,
			created_at
		)
		VALUES(?, ?, ?, ?)
		`,
		record.ID,
		1,
		string(rawJSON),
		now,
	)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *SQLiteRecordService) getLatestVersion(
	ctx context.Context,
	recordID int,
) (int, error) {

	query := `
	SELECT version
	FROM record_versions
	WHERE record_id = ?
	ORDER BY version DESC
	LIMIT 1
	`

	var version int

	err := s.db.QueryRowContext(
		ctx,
		query,
		recordID,
	).Scan(&version)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, ErrRecordDoesNotExist
	}

	if err != nil {
		return 0, err
	}

	return version, nil
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

	latestVersion, err := s.getLatestVersion(ctx, id)
	if err != nil {
		return entity.Record{}, err
	}

	rawJSON, err := json.Marshal(newRecord.Data)
	if err != nil {
		return entity.Record{}, err
	}

	_, err = s.db.ExecContext(
		ctx,
		`
		INSERT INTO record_versions(
			record_id,
			version,
			data,
			created_at
		)
		VALUES(?, ?, ?, ?)
		`,
		id,
		latestVersion+1,
		string(rawJSON),
		time.Now().UTC(),
	)

	if err != nil {
		return entity.Record{}, err
	}

	return newRecord, nil
}

func (s *SQLiteRecordService) ListRecordVersions(
	ctx context.Context,
	id int,
) ([]entity.RecordVersion, error) {

	query := `
	SELECT version, data, created_at
	FROM record_versions
	WHERE record_id = ?
	ORDER BY version ASC
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

		version.RecordID = id
		version.Data = data

		versions = append(versions, version)
	}

	return versions, nil
}
