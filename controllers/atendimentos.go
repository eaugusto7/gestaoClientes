package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

// GetAllAtendimentos godoc
// @Summary Todos Atendimentos
// @Description Obtem todos os atendimentos vindos do banco de dados
// @Tags Atendimentos
// @Produce json
// @Sucess 200 [object] model.Atendimento
// @Router /api/v1/atendimentos [get]
func GetAllAtendimentos(context *gin.Context) {
	var atendimento []models.Atendimento
	database.Database.Find(&atendimento)
	context.JSON(200, atendimento)
}

// GetAtendimentoById godoc
// @Summary Busca Atendimento por Id
// @Description Obtem  o json de um determinado atendimento, filtrado por id
// @Tags Atendimentos
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Atendimento
// @Failure 404 {object} httputil.HTTPError "Atendimento não encontrado"
// @Router /api/v1/atendimentos/{id} [get]
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

// InsertAtendimento godoc
// @Summary Insere Atendimento
// @Description Cria um novo atendimento no banco de dados
// @Tags Atendimentos
// @Accept json
// @Produce json
// @Param   atendimento     body    models.Atendimento     true		"Json Atendimento"
// @Sucess 200 {object} model.Atendimento
// @Failure 404 {object} httputil.HTTPError "Erro: Atendimento não cadastrado"
// @Router /api/v1/atendimentos [post]
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

// UpdateAtendimento godoc
// @Summary Atualiza Atendimento
// @Description Atualiza as informações de um determinado atendimento no banco de dados
// @Tags Atendimentos
// @Accept json
// @Produce json
// @Param   atendimento     body    models.Atendimento     true		"Json Atendimento"
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Atendimento
// @Failure 400 {object} httputil.HTTPError "Erro: Atendimento não existe"
// @Router /api/v1/atendimentos/{id} [patch]
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

// DeleteAtendimento godoc
// @Summary Deleta Atendimento
// @Description Remove o atendente indicado pelo id no banco de dados
// @Tags Atendimentos
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object}
// @Failure 400 {object} httputil.HTTPError "Erro: Não encontrado"
// @Router /api/v1/atendimentos/{id} [delete]
func DeleteAtendimento(context *gin.Context) {
	var atendimento models.Atendimento
	id := context.Params.ByName("id")
	database.Database.Delete(&atendimento, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Atendimento Deletado"})
}
