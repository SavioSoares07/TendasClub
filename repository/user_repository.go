package repository

import (
	"tendasclub/database"
	"tendasclub/models"
)

// Inserção de um novo usuário no banco de dados
func InsertUser(user models.User) error {
	_, err := database.DB.Exec("INSERT INTO users (name, email, password, number, role) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.Number, user.Role)
	return  err
}


// Verifica se um usuário com o email fornecido já existe no banco de dados
func UserExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	err := database.DB.QueryRow(query, email).Scan(&exists)
	return exists, err
}


// Busca um usuário pelo email no banco de dados
func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	// Executa a consulta para buscar o usuário pelo email
	err := database.DB.QueryRow("SELECT id, name, email, password, number, role FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Number, &user.Role)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

