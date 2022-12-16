package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// GetAllAtendentes godoc
// @Summary Todos Atendentes
// @Description Obtem todos os atendentes vindos do banco de dados
// @Tags Atendentes
// @Produce json
// @Sucess 200 [object] model.Atendente
// @Router /api/v1/atendentes [get]
func GetAllAtendentes(context *gin.Context) {
	var atendente []models.Atendente
	database.Database.Find(&atendente)
	context.JSON(200, atendente)
}

//Obtem  o json de um determinado atendente, filtrado por id
func GetAtendenteById(context *gin.Context) {
	var atendente models.Atendente
	id := context.Params.ByName("id")
	database.Database.First(&atendente, id)

	if atendente.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Message": "Atendente não encontrado"})
		return
	}
	context.JSON(http.StatusOK, atendente)
}

//Cria um novo atendente no banco de dados
func InsertAtendente(context *gin.Context) {
	var atendente models.Atendente

	if error := context.ShouldBindJSON(&atendente); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message": error.Error()})
		return
	}

	if error := models.ValidaDadosAtendente(&atendente); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message": error.Error()})
		return
	}

	database.Database.Create(&atendente)
	context.JSON(http.StatusOK, atendente)
}

//Atualiza as informações de um determinado atendente no banco de dados
func UpdateAtendente(context *gin.Context) {
	var atendente models.Atendente

	id := context.Params.ByName("id")
	database.Database.First(&atendente, id)

	if error := context.ShouldBindJSON(&atendente); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": error.Error()})
		return
	}

	database.Database.Save(&atendente)
	context.JSON(http.StatusOK, atendente)
}

//Remove o atendente indicado pelo id no banco de dados
func DeleteAtendente(context *gin.Context) {
	var atendente models.Atendente

	id := context.Params.ByName("id")
	database.Database.Delete(&atendente, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Atendente Deletado"})
}
