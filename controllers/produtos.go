package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

//Obtem todos os produtos vindos do banco de dados
func GetAllProdutos(context *gin.Context) {
	var produtos []models.Produtos
	database.Database.Find(&produtos)
	context.JSON(200, produtos)
}

//Obtem o json de um determinado produto, filtrado por username
func GetProdutoById(context *gin.Context) {
	var produto models.Produtos

	id := context.Params.ByName("id")
	database.Database.First(&produto, id)

	if produto.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Message": "Produto não encontrado"})
		return
	}
	context.JSON(http.StatusOK, produto)
}

//Cria um novo produto no banco de dados
func InsertProduto(context *gin.Context) {
	var produto models.Produtos

	if error := context.ShouldBindJSON(&produto); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message": error.Error()})
		return
	}

	if error := models.ValidaDadosProdutos(&produto); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message": error.Error()})
		return
	}

	database.Database.Create(&produto)
	context.JSON(http.StatusOK, produto)
}

//Atualiza as informações de um determinado produto no banco de dados
func UpdateProduto(context *gin.Context) {
	var produto models.Produtos

	id := context.Params.ByName("id")
	database.Database.First(&produto, id)

	if error := context.ShouldBindJSON(&produto); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": error.Error()})
		return
	}
	database.Database.Save(&produto)
	context.JSON(http.StatusOK, produto)
}

//Remove o produto indicado pelo id no banco de dados
func DeleteProduto(context *gin.Context) {
	var produto models.Produtos

	id := context.Params.ByName("id")
	database.Database.Delete(&produto, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Produto Deletado"})
}
