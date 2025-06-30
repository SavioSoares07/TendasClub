package controllers

import (
	"fmt"
	"tendasclub/models"
	"tendasclub/repository"
)

func RegisterTimeRecord(email string, timeRecord models.TimeRecord) (string, error) {


	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("error retrieving user: %w", err)
	}

	record := models.TimeRecord{
		UserID: int64(user.ID),
		TimeStart: timeRecord.TimeStart,
		TimeEnd:   timeRecord.TimeEnd,
		Category: timeRecord.Category,
		Status:  timeRecord.Status,
		Duration: timeRecord.Duration,
		Notes: timeRecord.Notes,
		CreatedAt: timeRecord.CreatedAt,
	}

	fmt.Println("Registering time record for user:", record)


	return "teste", nil

}