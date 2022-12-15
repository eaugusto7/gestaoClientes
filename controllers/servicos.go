package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

func GetAllServicos(context *gin.Context) {
	var servico []models.Servico
	database.Database.Find(&servico)
	context.JSON(200, servico)
}

func GetByIdServicos(context *gin.Context) {
	var servico models.Servico
	id := context.Params.ByName("id")
	database.Database.First(&servico, id)

	if servico.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Not found": "Servico não encontrado"})
		return
	}

	context.JSON(http.StatusOK, servico)
}

func InsertServicos(context *gin.Context) {
	var servico models.Servico
	if err := context.ShouldBindJSON(&servico); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidaDadosServico(&servico); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}
	database.Database.Create(&servico)
	context.JSON(http.StatusOK, servico)
}

func UpdateServicos(context *gin.Context) {
	var servico models.Servico

	id := context.Params.ByName("id")
	database.Database.First(&servico, id)

	if err := context.ShouldBindJSON(&servico); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.Database.Save(&servico)
	context.JSON(http.StatusOK, servico)
}

func DeleteServicos(context *gin.Context) {
	var servico models.Servico
	id := context.Params.ByName("id")
	database.Database.Delete(&servico, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Serviço Deletado"})
}
