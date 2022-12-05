package controllers

import (
	"encoding/json"
	"net/http"

	db "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gorilla/mux"
)

func GetAllAtendimentos(w http.ResponseWriter, r *http.Request) {
	var a []models.Atendimento
	db.DB.Find(&a)
	json.NewEncoder(w).Encode(a)
}

func GetByIdAtendimentos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var atendimento models.Atendimento
	db.DB.First(&atendimento, id)
	json.NewEncoder(w).Encode(atendimento)
}

func InsertAtendimentos(w http.ResponseWriter, r *http.Request) {
	var newAtendimento models.Atendimento
	json.NewDecoder(r.Body).Decode(&newAtendimento)
	db.DB.Create(&newAtendimento)
	json.NewEncoder(w).Encode(newAtendimento)
}

func UpdateAtendimentos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var atendimento models.Atendimento
	db.DB.First(&atendimento, id)
	json.NewDecoder(r.Body).Decode(&atendimento)
	db.DB.Save(&atendimento)
	json.NewEncoder(w).Encode(atendimento)
}

func DeleteAtendimento(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var atendimento models.Atendimento
	db.DB.Delete(&atendimento, id)
	json.NewEncoder(w).Encode(atendimento)
}
