package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tendasclub/controllers"
	"tendasclub/models"
)

//POST /registerTime
func RegisterTimeHandler(w http.ResponseWriter, r *http.Request){
	email := r.Header.Get("X-User-Email")
	if email == "" {
		http.Error(w, "Email não encontrado no cabeçalho da requisição", http.StatusUnauthorized)
		return
	}


	var TimeRecord models.TimeRecord

	err := json.NewDecoder(r.Body).Decode(&TimeRecord)
	if err != nil {
		http.Error(w, "Erro ao ler o registro de tempo: "+err.Error(), http.StatusBadRequest)
		return
	}

	res, err := controllers.RegisterTimeRecord(email, TimeRecord)
	if err != nil {
		http.Error(w, "Erro ao registrar o tempo: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(res)
	fmt.Fprintln(w, "Tempo registrado com sucesso para o usuário:", email)
}

//GET /GetAllTimeRecords

func GetAllTimeRecords(w http.ResponseWriter, r *http.Request){
	var allTimesRecords []models.TimeRecord

	allTimesRecords, err := controllers.GetAllTime()
	if err != nil {
		http.Error(w, "Erro ao ler todos os dados gravados" +err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(allTimesRecords)
}

// Get /GetTimeRecodrById
func GetTimeRecordById(w http.ResponseWriter, r *http.Request){

}