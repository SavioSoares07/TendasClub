package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"tendasclub/services"
)

func RegisterTimeHandler(w http.ResponseWriter, r *http.Request){
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Token de autenticação não fornecido", http.StatusUnauthorized)
		return
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		http.Error(w, "Formato de token inválido", http.StatusUnauthorized)
		return
	}
	token := tokenParts[1]

	email, err := services.VerifyToken(token)
	if err != nil {
		http.Error(w, "Token inválido: "+err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprintln(w, email)

}