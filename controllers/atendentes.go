package controllers

import (
	"net/http"

	db "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

func GetAllAtendente(context *gin.Context) {
	var atendente []models.Atendente
	db.DB.Find(&atendente)
	context.JSON(200, atendente)
}

func GetAtendenteById(context *gin.Context) {
	var atendente models.Atendente

	id := context.Params.ByName("id")
	db.DB.First(&atendente, id)

	if atendente.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Not found": "Atendente não encontrado"})
		return
	}
	context.JSON(http.StatusOK, atendente)
}

func InsertAtendente(context *gin.Context) {
	var atendente models.Atendente

	if err := context.ShouldBindJSON(&atendente); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}
	db.DB.Create(&atendente)
	context.JSON(http.StatusOK, atendente)
}

func UpdateAtendente(context *gin.Context) {
	var atendente models.Atendente

	id := context.Params.ByName("id")
	db.DB.First(&atendente, id)

	if err := context.ShouldBindJSON(&atendente); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.DB.Save(&atendente)
	context.JSON(http.StatusOK, atendente)
}

func DeleteAtendente(context *gin.Context) {
	var atendente models.Atendente

	id := context.Params.ByName("id")
	db.DB.Delete(&atendente, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Atendente Deletado"})
}
