package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

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

	return err
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

	return newRecord, nil
}
