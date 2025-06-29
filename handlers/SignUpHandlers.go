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
	
	//Criação de uma variável do tipo User
	//Essa variável irá guardar os dados do usuário que serão enviados na requisição
	var user models.User

	//Guarda os dados do usuário que foram enviados na requisição
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao ler os dados do usuário", http.StatusBadRequest)
		return
	}
	//Validação dos dados do usuário
	err = validate.Validade(user)
	if(err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Chamada do controller para criar o usuário
	//O controller irá chamar o model para salvar no banco de dados
	res, err := controllers.CreateUser(user)
	if err != nil {
		http.Error(w, "Erro ao criar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Usuário criado com sucesso: %v", res)
	w.WriteHeader(http.StatusCreated)
}