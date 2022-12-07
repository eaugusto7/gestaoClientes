package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAll(context *gin.Context) {
	var clientes []models.Cliente
	db.DB.Find(&clientes)
	context.JSON(200, clientes)
}

func GetById(context *gin.Context) {
	var cliente models.Cliente
	id := context.Params.ByName("id")
	db.DB.First(&cliente, id)

	if cliente.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Not found": "Cliente não encontrado"})
		return
	}

	context.JSON(http.StatusOK, cliente)
}

func InsertClient(w http.ResponseWriter, r *http.Request) {
	var newClient models.Cliente
	json.NewDecoder(r.Body).Decode(&newClient)
	db.DB.Create(&newClient)
	json.NewEncoder(w).Encode(newClient)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var cliente models.Cliente
	db.DB.First(&cliente, id)
	json.NewDecoder(r.Body).Decode(&cliente)
	db.DB.Save(&cliente)
	json.NewEncoder(w).Encode(cliente)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var cliente models.Cliente
	db.DB.Delete(&cliente, id)
	json.NewEncoder(w).Encode(cliente)
}
