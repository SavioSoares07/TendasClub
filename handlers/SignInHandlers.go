package handlers

import (
	"encoding/json"
	"net/http"
	"tendasclub/controllers"
	"tendasclub/models"
	"tendasclub/services"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {

	//Salvar dados do usuário
	var Credentials models.Credentials

	//Ler os dados do usuário
	//O usuário irá enviar os dados no corpo da requisição
	err := json.NewDecoder(r.Body).Decode(&Credentials)

	if err != nil {
		http.Error(w, "Erro ao ler as credenciais de acesso", http.StatusBadRequest)
		return
	}

	//Verificar se o usuário existe
	//Se o usuário não existir, retornar um erro 404
	exist, err := services.UserExists(Credentials.Email)
	if( err != nil) {
		http.Error(w, "Erro ao verificar se o usuário existe: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !exist {
		http.Error(w, "Usuário não existe", http.StatusNotFound)
		return
	}	
	
	// Define o header como JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//Chamada do controller para fazer o login do usuário
	//O controller irá chamar o model para verificar as credenciais
	token, err := controllers.LoginUser(Credentials)
	if err != nil {
		http.Error(w, "Erro ao fazer login: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login realizado com sucesso",
		"token":   token,
	})

}