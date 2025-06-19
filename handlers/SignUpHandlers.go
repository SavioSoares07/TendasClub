package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tendasclub/models"
)

func SignUpHandler(w http.ResponseWriter, r* http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Erro ao ler os dados do usuário", http.StatusBadRequest)
		return
	}
	fmt.Printf("Nome do usuário: %s, Email do usuário: %s)", user.Name, user.Email)
}