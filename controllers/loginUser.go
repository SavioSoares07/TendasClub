package controllers

import (
	"fmt"
	"tendasclub/models"
	"tendasclub/services"

	"golang.org/x/crypto/bcrypt"
)


func LoginUser(c models.Credentials) (string, error) {
	UserExists, err := services.UserExists(c.Email)
	if err != nil {
		return "", err
	}
	if !UserExists {
		return "Usário não existe", nil
	}

	user, err := services.GetUserByEmail(c.Email)
	if err != nil {		
		return "", err
	}

	storeHash := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(storeHash), []byte(c.Password))
	
	if err != nil {
		return "Senha incorreta", nil
	}


	tokenString, err := services.CreateToken(user.Email)
	if err != nil {
		return "Erro ao criar o token", err
	}
	fmt.Print(tokenString)
	return "Usuário logado com sucesso", nil
}
