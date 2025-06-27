package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tendasclub/controllers"
	"tendasclub/models"
	"tendasclub/validate"
)

func SignUpHandler(w http.ResponseWriter, r* http.Request) {


	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao ler os dados do usuário", http.StatusBadRequest)
		return
	}

	err = validate.Validade(user)
	if(err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := controllers.CreateUser(user)
	if err != nil {
		http.Error(w, "Erro ao criar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Usuário criado com sucesso: %v", res)
	w.WriteHeader(http.StatusCreated)
}