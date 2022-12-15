package controllers

import (
	"fmt"
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllClientes(context *gin.Context) {
	var clientes []models.Cliente
	database.Database.Find(&clientes)
	context.JSON(200, clientes)
}

func GetByIdClientes(context *gin.Context) {
	var cliente models.Cliente
	id := context.Params.ByName("id")
	database.Database.First(&cliente, id)

	if cliente.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Not found": "Cliente não encontrado"})
		return
	}

	context.JSON(http.StatusOK, cliente)
}

func InsertClient(context *gin.Context) {
	var cliente models.Cliente
	if err := context.ShouldBindJSON(&cliente); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidaDadosClientes(&cliente); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}

	database.Database.Create(&cliente)
	context.JSON(http.StatusOK, cliente)
}

func UpdateClient(context *gin.Context) {
	var cliente models.Cliente
	id := context.Params.ByName("id")
	database.Database.First(&cliente, id)

	if err := context.ShouldBindJSON(&cliente); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.Database.Save(&cliente)
	context.JSON(http.StatusOK, cliente)
}

func DeleteClient(context *gin.Context) {
	var cliente models.Cliente
	id := context.Params.ByName("id")
	database.Database.Delete(&cliente, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Cliente Deletado"})
}
