package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

func GetAllProdutos(context *gin.Context) {
	var produtos []models.Produtos
	database.Database.Find(&produtos)
	context.JSON(200, produtos)
}

func GetProdutoById(context *gin.Context) {
	var produto models.Produtos

	id := context.Params.ByName("id")
	database.Database.First(&produto, id)

	if produto.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto não encontrado"})
		return
	}
	context.JSON(http.StatusOK, produto)
}

func InsertProduto(context *gin.Context) {
	var produto models.Produtos

	if err := context.ShouldBindJSON(&produto); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidaDadosProdutos(&produto); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}

	database.Database.Create(&produto)
	context.JSON(http.StatusOK, produto)
}

func UpdateProduto(context *gin.Context) {
	var produto models.Produtos

	id := context.Params.ByName("id")
	database.Database.First(&produto, id)

	if err := context.ShouldBindJSON(&produto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.Database.Save(&produto)
	context.JSON(http.StatusOK, produto)
}

func DeleteProduto(context *gin.Context) {
	var produto models.Produtos

	id := context.Params.ByName("id")
	database.Database.Delete(&produto, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Produto Deletado"})
}
