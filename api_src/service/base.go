package service

import (
	"database/sql"
	"encoding/json"
	"time"
)

type SqlNullTime struct {
	sql.NullTime
}

func (b SqlNullTime) MarshalJSON() ([]byte, error) {
	var timestamp time.Time
	if b.Valid {
		timestamp = b.Time
	}
	return json.Marshal(timestamp)
}

type Base struct {
	ID        uint         `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
	DeletedAt *SqlNullTime `gorm:"index" json:"deleted_at,omitempty"`
}
