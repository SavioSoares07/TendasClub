package controllers

import (
	"fmt"
	"tendasclub/models"
	"tendasclub/repository"
	"tendasclub/services"
	"tendasclub/util/security"
)

//Controller da criação de conta
func CreateUser(user models.User) (string, error) {
    return services.CreateUser(user)
}

//Controller de fazer o login
func LoginUser(c models.Credentials) (string, error) {

	// Verifica se o usuário existe
	// Se o usuário não existir, retornar um erro 404
	UserExists, err := repository.UserExists(c.Email)
	if err != nil {
		return "", err
	}
	if !UserExists {
		return "Usário não existe", nil
	}

	//Guarda o usuário em uma variável
	user, err := repository.GetUserByEmail(c.Email)
	if err != nil {		
		return "", err
	}

	//Guarda a senha do usuário em uma variável
	storeHash := user.Password
		
	//Compara a senha do usuário com a senha armazenada no banco de dados
	//Se a senha não for igual, retornar um erro 401

	err = security.ComparePassword(storeHash, c.Password)
	if err != nil {
		return "Senha incorreta", nil
	}
	
	//Se a senha for igual, criar o token de acesso
	//O token será usado para autenticar o usuário em requisições futuras
	tokenString, err := services.CreateToken(user.Email)
	if err != nil {
		return "Erro ao criar o token", err
	}
	return tokenString, nil
}

//Autlizar senha do usuário

func UpdatePassword(email string, passwordChange models.PasswordChange) (string, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	err = security.ComparePassword(user.Password, passwordChange.OldPassword)
	if err != nil {
		return "", fmt.Errorf("senha incorreta")
	}

	hashedPassword, err := security.HashPassword(passwordChange.NewPassword)
	if err != nil {
		return "", fmt.Errorf("erro ao criptografar nova senha: %w", err)
	}

	user.Password = hashedPassword

	message, err := repository.UpdatePasswordUser(user)
	if err != nil {
		return "", fmt.Errorf("erro ao atualizar senha: %w", err)
	}

	fmt.Println(message)
	return "Senha alterada com sucesso", nil
}
