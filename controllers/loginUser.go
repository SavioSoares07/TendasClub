package controllers

import (
	"fmt"
	"tendasclub/models"
	"tendasclub/services"

	"golang.org/x/crypto/bcrypt"
)


func LoginUser(c models.Credentials) (string, error) {

	// Verifica se o usuário existe
	// Se o usuário não existir, retornar um erro 404
	UserExists, err := services.UserExists(c.Email)
	if err != nil {
		return "", err
	}
	if !UserExists {
		return "Usário não existe", nil
	}

	//Guarda o usuário em uma variável
	user, err := services.GetUserByEmail(c.Email)
	if err != nil {		
		return "", err
	}

	//Guarda a senha do usuário em uma variável
	storeHash := user.Password
	
	//Compara a senha do usuário com a senha armazenada no banco de dados
	//Se a senha não for igual, retornar um erro 401
	err = bcrypt.CompareHashAndPassword([]byte(storeHash), []byte(c.Password))
	if err != nil {
		return "Senha incorreta", nil
	}

	//Se a senha for igual, criar o token de acesso
	//O token será usado para autenticar o usuário em requisições futuras
	tokenString, err := services.CreateToken(user.Email)
	if err != nil {
		return "Erro ao criar o token", err
	}
	fmt.Print(tokenString)
	return tokenString, nil
}
