package repository

import (
	"tendasclub/database"
	"tendasclub/models"
)

// InserirTimeRecord insere um novo registro de tempo no banco de dados.
func InsertTimeRecord(timeRecord models.TimeRecord) error {
	_, err := database.DB.Exec("INSERT INTO time_records (user_id, time_start, time_end, category, status, duration, notes, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		timeRecord.UserID, timeRecord.TimeStart, timeRecord.TimeEnd, timeRecord.Category, timeRecord.Status, timeRecord.Duration, timeRecord.Notes, timeRecord.CreatedAt)
	return err
}

//Verificar se horario está disponivel
func CheckTimeRecord(timeRecord models.TimeRecord) (bool, error){
	var recorded bool

	query := "select exists(select 1 from time_records where time_start = ?)"
	err := database.DB.QueryRow(query, timeRecord.TimeStart	).Scan(&recorded)

	return recorded, err
}	