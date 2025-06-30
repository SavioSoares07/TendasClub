package models

import "tendasclub/enum"

type TimeRecord struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
	Category  enum.Category `json:"category"`
	Status   enum.Status   `json:"status"`
	Duration  int64  `json:"duration"` // Duration in minutes
	Notes     string `json:"notes"`
	CreatedAt string `json:"created_at"`
}