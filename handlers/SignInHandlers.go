package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tendasclub/controllers"
	"tendasclub/models"
	"tendasclub/services"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {


	var Credentials models.Credentials
	err := json.NewDecoder(r.Body).Decode(&Credentials)

	if err != nil {
		http.Error(w, "Erro ao ler as credenciais de acesso", http.StatusBadRequest)
		return
	}

	exist, err := services.UserExists(Credentials.Email)
	if( err != nil) {
		http.Error(w, "Erro ao verificar se o usuário existe: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !exist {
		http.Error(w, "Usuário não existe", http.StatusNotFound)
		return
	}	

	res, err := controllers.LoginUser(Credentials)
	if err != nil {
		http.Error(w, "Erro ao fazer login: "+err.Error(), http.StatusInternalServerError)
		return
	}


	fmt.Fprintf(w, "Login Realizado com sucesso: %v", res)
}