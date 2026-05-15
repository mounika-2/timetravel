package entity

import "time"

type RecordVersion struct {
	RecordID  int               `json:"record_id"`
	Version   int               `json:"version"`
	Data      map[string]string `json:"data"`
	CreatedAt time.Time         `json:"created_at"`
}
