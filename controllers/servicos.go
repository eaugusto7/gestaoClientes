package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

//Obtem todos os serviços vindos do banco de dados
func GetAllServicos(context *gin.Context) {
	var servico []models.Servico
	database.Database.Find(&servico)
	context.JSON(200, servico)
}

//Obtem o json de um determinado serviço, filtrado por id
func GetServicosById(context *gin.Context) {
	var servico models.Servico
	id := context.Params.ByName("id")
	database.Database.First(&servico, id)

	if servico.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Message": "Servico não encontrado"})
		return
	}

	context.JSON(http.StatusOK, servico)
}

//Cria um novo serviço no banco de dados
func InsertServico(context *gin.Context) {
	var servico models.Servico
	if error := context.ShouldBindJSON(&servico); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message": error.Error()})
		return
	}

	if error := models.ValidaDadosServico(&servico); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message": error.Error()})
		return
	}
	database.Database.Create(&servico)
	context.JSON(http.StatusOK, servico)
}

//Atualiza as informações de um determinado serviço no banco de dados
func UpdateServico(context *gin.Context) {
	var servico models.Servico

	id := context.Params.ByName("id")
	database.Database.First(&servico, id)

	if error := context.ShouldBindJSON(&servico); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": error.Error()})
		return
	}
	database.Database.Save(&servico)
	context.JSON(http.StatusOK, servico)
}

//Remove o serviço indicado pelo id no banco de dados
func DeleteServico(context *gin.Context) {
	var servico models.Servico
	id := context.Params.ByName("id")
	database.Database.Delete(&servico, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Serviço Deletado"})
}
