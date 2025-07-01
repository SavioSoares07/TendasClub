package services

import (
	"fmt"
	"tendasclub/models"
	"tendasclub/repository"
	"tendasclub/util/security"
)

func CreateUser(user models.User) (string, error){
	// Verifica se o usuário já existe
	
	exists, err := repository.UserExists(user.Email)
	if err != nil {
		return "", fmt.Errorf("erro ao verificar se o usuário já existe: %W", err)
	}

	if exists {
		return "", fmt.Errorf("usuário já cadastrado: %W", err)
	}

	//Criptografa a senha do usuário
	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return "", fmt.Errorf("erro ao criptografar a senha: %W", err)
	}

	

	user.Password = string(hashedPassword)

	err = repository.InsertUser(user)
	if err != nil {
		return "", fmt.Errorf("erro ao inserir o usuário no banco de dados: %W", err)
	}

	return "Usuário cadastrado com sucesso!", nil
}