package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tendasclub/database"
	"tendasclub/models"
	"tendasclub/validate"

	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(w http.ResponseWriter, r* http.Request) {

	database.ConnectionDB()

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro ao criptografar a senha", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO users (name, email, password, number, role) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, hashedPassword, user.Number, user.Role)
	if err != nil {
		http.Error(w, "Erro ao inserir o usuário no banco de dados", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Usuário %s cadastrado com sucesso!", user.Name)
}