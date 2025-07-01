package services

import (
	"tendasclub/models"
	"tendasclub/repository"
)

func CreateTimeRecord(timeRecord models.TimeRecord) (string, error) {
	err := repository.InsertTimeRecord(timeRecord)
	if err != nil {
		return "", err
	}

	return "Registro de tempo criado com sucesso!", nil
}