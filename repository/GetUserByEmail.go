package repository

import (
	"tendasclub/database"
	"tendasclub/models"
)

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	// Executa a consulta para buscar o usuário pelo email
	err := database.DB.QueryRow("SELECT id, name, email, password, number, role FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Number, &user.Role)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}