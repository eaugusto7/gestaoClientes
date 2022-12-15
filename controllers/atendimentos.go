package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

func GetAllAtendimentos(context *gin.Context) {
	var atendimento []models.Atendimento
	database.Database.Find(&atendimento)
	context.JSON(200, atendimento)
}

func GetByIdAtendimentos(context *gin.Context) {
	var atendimento models.Atendimento
	id := context.Params.ByName("id")
	database.Database.First(&atendimento, id)

	if atendimento.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Not found": "Atendimento não encontrado"})
		return
	}

	context.JSON(http.StatusOK, atendimento)
}

func InsertAtendimentos(context *gin.Context) {
	var atendimento models.Atendimento
	if err := context.ShouldBindJSON(&atendimento); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosAtendimento(&atendimento); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}

	database.Database.Create(&atendimento)
	context.JSON(http.StatusOK, atendimento)
}

func UpdateAtendimentos(context *gin.Context) {
	var atendimento models.Atendimento
	id := context.Params.ByName("id")
	database.Database.First(&atendimento, id)

	if err := context.ShouldBindJSON(&atendimento); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.Database.Save(&atendimento)
	context.JSON(http.StatusOK, atendimento)
}

func DeleteAtendimento(context *gin.Context) {
	var atendimento models.Atendimento
	id := context.Params.ByName("id")
	database.Database.Delete(&atendimento, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Atendimento Deletado"})
}
