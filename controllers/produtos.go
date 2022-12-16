package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

// GetAllProdutos godoc
// @Summary Todos Produtos
// @Description Obtem todos os produtos vindos do banco de dados
// @Tags Produtos
// @Produce json
// @Sucess 200 [object] model.Produtos
// @Router /api/v1/produtos [get]
func GetAllProdutos(context *gin.Context) {
	var produtos []models.Produtos
	database.Database.Find(&produtos)
	context.JSON(200, produtos)
}

// GetProdutoById godoc
// @Summary Busca Produto por Id
// @Description Obtem  o json de um determinado produto, filtrado por id
// @Tags Produtos
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Produtos
// @Failure 404 {object} httputil.HTTPError "Produto não encontrado"
// @Router /api/v1/produtos/{id} [get]
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

// InsertProduto godoc
// @Summary Insere Produto
// @Description Cria um novo produto no banco de dados
// @Tags Produtos
// @Accept json
// @Produce json
// @Param   produto     body    models.Produtos     true		"Json Produto"
// @Sucess 200 {object} model.Produto
// @Failure 404 {object} httputil.HTTPError "Erro: Produto não cadastrado"
// @Router /api/v1/produtos [post]
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

// UpdateProduto godoc
// @Summary Atualiza Produto
// @Description Atualiza as informações de um determinado produto no banco de dados
// @Tags Produtos
// @Accept json
// @Produce json
// @Param   produto     body    models.Produtos     true		"Json Produto"
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Produto
// @Failure 400 {object} httputil.HTTPError "Erro: Produto não existe"
// @Router /api/v1/produtos/{id} [patch]
func UpdateProduto(context *gin.Context) {
	var produto models.Produtos

	id := context.Params.ByName("id")
	database.Database.First(&produto, id)

	if error := context.ShouldBindJSON(&produto); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": error.Error()})
		return
	}

	if produto.Id == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Não encontrado"})
		return
	}

	database.Database.Save(&produto)
	context.JSON(http.StatusOK, produto)
}

// DeleteProduto godoc
// @Summary Deleta Produto
// @Description Remove o produto indicado pelo id no banco de dados
// @Tags Produtos
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object}
// @Failure 400 {object} httputil.HTTPError "Erro: Não encontrado"
// @Router /api/v1/produtos/{id} [delete]
func DeleteProduto(context *gin.Context) {
	var produto models.Produtos

	id := context.Params.ByName("id")
	database.Database.Delete(&produto, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Produto Deletado"})
}
