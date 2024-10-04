package domain

import (
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	UID         string    `json:"uid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
