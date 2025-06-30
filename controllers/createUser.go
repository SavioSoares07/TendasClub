package controllers

import (
	"fmt"
	"tendasclub/database"
	"tendasclub/models"
	"tendasclub/repository"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (string,  error) {
	// Verifica se o usuário já existe
	exists, err := repository.UserExists(user.Email)
	if err != nil {
		return "", fmt.Errorf("erro ao verificar se o usuário já existe: %W", err)
	}

	if exists {
		return "", fmt.Errorf("usuário já cadastrado: %W", err)
	}

	// Criptografa a senha do usuário
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("erro ao criptografar a senha: %W", err)
	}

	// Insere o usuário no banco de dados
	_, err = database.DB.Exec("INSERT INTO users (name, email, password, number, role) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, hashedPassword, user.Number, user.Role)
	if err != nil {
		return "", fmt.Errorf("erro ao inserir o usuário no banco de dados: %w", err)
	}
	fmt.Printf("Usuário %s cadastrado com sucesso!", user.Name)

	messagem := "Usuário cadastrado com sucesso!"

	return messagem, nil
}
