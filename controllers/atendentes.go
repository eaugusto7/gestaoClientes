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

// GetAtendenteById godoc
// @Summary Busca Atendente por Id
// @Description Obtem  o json de um determinado atendente, filtrado por id
// @Tags Atendentes
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Atendente
// @Failure 404 {object} httputil.HTTPError "Atendente não encontrado"
// @Router /api/v1/atendentes/{id} [get]
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

// GetAtendenteByName godoc
// @Summary Busca Atendente por Nome
// @Description Obtem  o json de um determinado atendente, filtrado por nome
// @Tags Atendentes
// @Produce json
// @Param   nome     path    string     true        "Nome"
// @Sucess 200 {object} model.Atendente
// @Failure 404 {object} httputil.HTTPError "Atendente não encontrado"
// @Router /api/v1/atendentes/nome/{nome} [get]
func GetAtendenteByName(context *gin.Context) {
	var atendente models.Atendente
	nome := context.Params.ByName("nome")

	database.Database.Where("nome like ?", "%"+nome+"%").Find(&atendente)

	if atendente.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Message": "Atendente não encontrado"})
		return
	}
	context.JSON(http.StatusOK, atendente)
}

// InsertAtendente godoc
// @Summary Insere Atendente
// @Description Cria um novo atendente no banco de dados
// @Tags Atendentes
// @Accept json
// @Produce json
// @Param   atendente     body    models.Atendente     true		"Json Atendente"
// @Sucess 200 {object} model.Atendente
// @Failure 404 {object} httputil.HTTPError "Erro: Atendente não cadastrado"
// @Router /api/v1/atendentes [post]
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

// UpdateAtendente godoc
// @Summary Atualiza Atendente
// @Description Atualiza as informações de um determinado atendente no banco de dados
// @Tags Atendentes
// @Accept json
// @Produce json
// @Param   atendente     body    models.Atendente     true		"Json Atendente"
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Atendente
// @Failure 400 {object} httputil.HTTPError "Erro: Atendente não existe"
// @Router /api/v1/atendentes/{id} [patch]
func UpdateAtendente(context *gin.Context) {
	var atendente models.Atendente

	id := context.Params.ByName("id")
	database.Database.First(&atendente, id)

	if error := context.ShouldBindJSON(&atendente); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": error.Error()})
		return
	}

	if atendente.Id == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Não encontrado"})
		return
	}

	database.Database.Save(&atendente)
	context.JSON(http.StatusOK, atendente)
}

// DeleteAtendente godoc
// @Summary Deleta Atendente
// @Description Remove o atendente indicado pelo id no banco de dados
// @Tags Atendentes
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object}
// @Failure 400 {object} httputil.HTTPError "Erro: Não encontrado"
// @Router /api/v1/atendentes/{id} [delete]
func DeleteAtendente(context *gin.Context) {
	var atendente models.Atendente

	id := context.Params.ByName("id")
	database.Database.Delete(&atendente, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Atendente Deletado"})
}
