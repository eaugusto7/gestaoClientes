package controllers

import (
	"encoding/json"
	"net/http"

	db "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gorilla/mux"
)

func GetAllServicos(w http.ResponseWriter, r *http.Request) {
	var s []models.Servico
	db.DB.Find(&s)
	json.NewEncoder(w).Encode(s)
}

func GetByIdServicos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var servico models.Servico
	db.DB.First(&servico, id)
	json.NewEncoder(w).Encode(servico)
}

func InsertServicos(w http.ResponseWriter, r *http.Request) {
	var newServico models.Servico
	json.NewDecoder(r.Body).Decode(&newServico)
	db.DB.Create(&newServico)
	json.NewEncoder(w).Encode(newServico)
}

func UpdateServicos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var servico models.Servico
	db.DB.First(&servico, id)
	json.NewDecoder(r.Body).Decode(&servico)
	db.DB.Save(&servico)
	json.NewEncoder(w).Encode(servico)
}

func DeleteServicos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var servico models.Servico
	db.DB.Delete(&servico, id)
	json.NewEncoder(w).Encode(servico)
}
