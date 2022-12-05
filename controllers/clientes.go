package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var p []models.Cliente
	db.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}
