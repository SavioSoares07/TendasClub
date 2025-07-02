package controllers

import (
	"fmt"
	"tendasclub/models"
	"tendasclub/repository"
	"tendasclub/services"
	"time"
)

//Controlador Criação de Registro

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
		CreatedAt: time.Now().Local().Format("2006-01-02 15:04:05"),
	}

	RecordedTime, err := repository.CheckTimeRecord(record)
	if err != nil{
		return "", fmt.Errorf("erro ao verificar se o horario está marcado: %w", err)
	}


	if RecordedTime{
		return "", fmt.Errorf("horario já marcado : %W", err)
	}

	return services.CreateTimeRecord(record)

}

//Obter todos os tempos
func GetAllTime() ([]models.TimeRecord, error){
	return repository.GetAllTimeRecords()
}

//Obter dados pelo email
func GetTimeByEmail(email string) ([]models.TimeRecord, error){
	return repository.GetAllTimeRecordsByEmail(email)
}