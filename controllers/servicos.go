package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

// GetAllServicos godoc
// @Summary Todos Serviços
// @Description Obtem todos os serviços vindos do banco de dados
// @Tags Serviços
// @Produce json
// @Sucess 200 [object] model.Servico
// @Router /api/v1/servicos [get]
func GetAllServicos(context *gin.Context) {
	var servico []models.Servico
	database.Database.Find(&servico)
	context.JSON(200, servico)
}

// GetServicosById godoc
// @Summary Busca Serviço por Id
// @Description Obtem  o json de um determinado serviço, filtrado por id
// @Tags Serviços
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Servico
// @Failure 404 {object} httputil.HTTPError "Serviço não encontrado"
// @Router /api/v1/servicos/{id} [get]
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

// InsertServico godoc
// @Summary Insere Serviço
// @Description Cria um novo serviço no banco de dados
// @Tags Serviços
// @Accept json
// @Produce json
// @Param   servico     body    models.Servico     true		"Json Serviço"
// @Sucess 200 {object} model.Servico
// @Failure 404 {object} httputil.HTTPError "Erro: Serviço não cadastrado"
// @Router /api/v1/servicos [post]
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

// UpdateServico godoc
// @Summary Atualiza Serviço
// @Description Atualiza as informações de um determinado serviço no banco de dados
// @Tags Serviços
// @Accept json
// @Produce json
// @Param   produto     body    models.Servico     true		"Json Serviços"
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Servico
// @Failure 400 {object} httputil.HTTPError "Erro: Serviço não existe"
// @Router /api/v1/servicos/{id} [patch]
func UpdateServico(context *gin.Context) {
	var servico models.Servico

	id := context.Params.ByName("id")
	database.Database.First(&servico, id)

	if error := context.ShouldBindJSON(&servico); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": error.Error()})
		return
	}

	if servico.Id == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Não encontrado"})
		return
	}

	database.Database.Save(&servico)
	context.JSON(http.StatusOK, servico)
}

// DeleteServico godoc
// @Summary Deleta Serviço
// @Description Remove o serviço indicado pelo id no banco de dados
// @Tags Serviços
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object}
// @Failure 400 {object} httputil.HTTPError "Erro: Não encontrado"
// @Router /api/v1/servicos/{id} [delete]
func DeleteServico(context *gin.Context) {
	var servico models.Servico
	id := context.Params.ByName("id")
	database.Database.Delete(&servico, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Serviço Deletado"})
}
