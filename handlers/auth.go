package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tendasclub/controllers"
	"tendasclub/models"
	"tendasclub/validate"
)

//POST /api/login
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


//POST /api/register
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Lê os dados da requisição
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao ler os dados do usuário", http.StatusBadRequest)
		return
	}

	// Valida os dados
	err = validate.ValidadeUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Cria o usuário
	res, err := controllers.CreateUser(user)
	if err != nil {
		http.Error(w, "Erro ao criar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Usuário criado com sucesso: %v", res)
}
