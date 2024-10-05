package handler

import "time"

type CreateEventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
}

type CreateEventResponse struct {
	ID  int64  `json:"id"`
	UID string `json:"uid"`
}
