package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

//Obtem todos os atendimentos vindos do banco de dados
func GetAllAtendimentos(context *gin.Context) {
	var atendimento []models.Atendimento
	database.Database.Find(&atendimento)
	context.JSON(200, atendimento)
}

//Obtem  o json de um determinado atendimento, filtrado por id
func GetAtendimentoById(context *gin.Context) {
	var atendimento models.Atendimento
	id := context.Params.ByName("id")
	database.Database.First(&atendimento, id)

	if atendimento.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Message": "Atendimento não encontrado"})
		return
	}

	context.JSON(http.StatusOK, atendimento)
}

//Cria um novo atendimento no banco de dados
func InsertAtendimento(context *gin.Context) {
	var atendimento models.Atendimento
	if error := context.ShouldBindJSON(&atendimento); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message": error.Error()})
		return
	}
	if error := models.ValidaDadosAtendimento(&atendimento); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message": error.Error()})
		return
	}

	database.Database.Create(&atendimento)
	context.JSON(http.StatusOK, atendimento)
}

//Atualiza as informações de um determinado atendimento no banco de dados
func UpdateAtendimento(context *gin.Context) {
	var atendimento models.Atendimento
	id := context.Params.ByName("id")
	database.Database.First(&atendimento, id)

	if error := context.ShouldBindJSON(&atendimento); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": error.Error()})
		return
	}

	database.Database.Save(&atendimento)
	context.JSON(http.StatusOK, atendimento)
}

//Remove o atendente indicado pelo id no banco de dados
func DeleteAtendimento(context *gin.Context) {
	var atendimento models.Atendimento
	id := context.Params.ByName("id")
	database.Database.Delete(&atendimento, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Atendimento Deletado"})
}
