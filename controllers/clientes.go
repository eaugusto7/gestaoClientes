package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

//Obtem todos os clientes vindos do banco de dados
func GetAllClientes(context *gin.Context) {
	var clientes []models.Cliente
	database.Database.Find(&clientes)
	context.JSON(200, clientes)
}

//Obtem  o json de um determinado cliente, filtrado por id
func GetClienteById(context *gin.Context) {
	var cliente models.Cliente
	id := context.Params.ByName("id")
	database.Database.First(&cliente, id)

	if cliente.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Message": "Cliente não encontrado"})
		return
	}
	context.JSON(http.StatusOK, cliente)
}

//Cria um novo cliente no banco de dados
func InsertCliente(context *gin.Context) {
	var cliente models.Cliente
	if error := context.ShouldBindJSON(&cliente); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message error: ": error.Error()})
		return
	}

	if err := models.ValidaDadosClientes(&cliente); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message error: ": err.Error()})
		return
	}
	database.Database.Create(&cliente)
	context.JSON(http.StatusOK, cliente)
}

//Atualiza as informações de um determinado cliente no banco de dados
func UpdateCliente(context *gin.Context) {
	var cliente models.Cliente
	id := context.Params.ByName("id")
	database.Database.First(&cliente, id)

	if error := context.ShouldBindJSON(&cliente); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message error: ": error.Error()})
		return
	}

	database.Database.Save(&cliente)
	context.JSON(http.StatusOK, cliente)
}

//Remove o cliente indicado pelo id no banco de dados
func DeleteCliente(context *gin.Context) {
	var cliente models.Cliente
	id := context.Params.ByName("id")
	database.Database.Delete(&cliente, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Cliente Deletado"})
}
