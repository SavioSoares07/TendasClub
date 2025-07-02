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

// Obter todos os horarios
func GetAllTimeRecords() ([]models.TimeRecord, error){
	var allTimes []models.TimeRecord 
	
	rows, err := database.DB.Query("SELECT * FROM time_records")
	if err != nil {
		return nil, err
	}

	defer rows.Close() 
	for rows.Next(){
		var tr models.TimeRecord
		err := rows.Scan(
			&tr.ID,
			&tr.UserID,
			&tr.TimeStart,
			&tr.TimeEnd,
			&tr.Category,
			&tr.Status,
			&tr.Duration,
			&tr.Notes,
			&tr.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		allTimes = append(allTimes, tr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return allTimes, nil

}

//Obter dados pelo id
