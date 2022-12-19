package controllers

import (
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// GetAllHorarios godoc
// @Summary Todos Quadros de Horários
// @Description Obtem todos os quadros de horários vindos do banco de dados
// @Tags Quadro de Horários
// @Produce json
// @Sucess 200 [object] model.Quadroshorarios
// @Router /api/v1/horarios [get]
func GetAllHorarios(context *gin.Context) {
	var Horario []models.Quadroshorarios
	database.Database.Find(&Horario)
	context.JSON(200, Horario)
}

// GetHorarioById godoc
// @Summary Busca Horario por Id
// @Description Obtem  o json de um determinado quadro de horários, filtrado por id
// @Tags Quadro de Horários
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Quadroshorarios
// @Failure 404 {object} httputil.HTTPError "Quadro de horários não encontrado"
// @Router /api/v1/horarios/{id} [get]
func GetHorarioById(context *gin.Context) {
	var horario models.Quadroshorarios
	id := context.Params.ByName("id")
	database.Database.First(&horario, id)

	if horario.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Message": "Quadro de horários não encontrado"})
		return
	}
	context.JSON(http.StatusOK, horario)
}

// InsertQuadroHorario godoc
// @Summary Insere Quadro de Horários
// @Description Cria um novo quadro de horários no banco de dados
// @Tags Quadro de Horários
// @Accept json
// @Produce json
// @Param   atendente     body    models.Quadroshorarios     true		"Json Quadroshorario"
// @Sucess 200 {object} model.Quadroshorarios
// @Failure 404 {object} httputil.HTTPError "Erro: Quadro de horários não cadastrado"
// @Router /api/v1/horarios [post]
func InsertQuadroHorario(context *gin.Context) {
	var horario models.Quadroshorarios

	if error := context.ShouldBindJSON(&horario); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message": error.Error()})
		return
	}
	database.Database.Create(&horario)
	context.JSON(http.StatusOK, horario)
}

// UpdateHorarios godoc
// @Summary Atualiza Quadro de Horário
// @Description Atualiza as informações de um determinado quadro de horários no banco de dados
// @Tags Quadro de Horários
// @Accept json
// @Produce json
// @Param   produto     body    models.Quadroshorarios     true		"Json Quadro de Horarios"
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Quadroshorarios
// @Failure 400 {object} httputil.HTTPError "Erro: Quadro de horários não existe"
// @Router /api/v1/horarios/{id} [patch]
func UpdateHorarios(context *gin.Context) {
	var horario models.Quadroshorarios

	id := context.Params.ByName("id")
	database.Database.First(&horario, id)

	if error := context.ShouldBindJSON(&horario); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": error.Error()})
		return
	}

	if horario.Id == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Não encontrado"})
		return
	}

	database.Database.Save(&horario)
	context.JSON(http.StatusOK, horario)
}

// DeleteHorario godoc
// @Summary Deleta Quadro de Horários
// @Description Remove o quadro de horário indicado pelo id no banco de dados
// @Tags Quadro de Horários
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object}
// @Failure 400 {object} httputil.HTTPError "Erro: Não encontrado"
// @Router /api/v1/horarios/{id} [delete]
func DeleteHorario(context *gin.Context) {
	var horario models.Quadroshorarios

	id := context.Params.ByName("id")
	database.Database.Delete(&horario, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Quadro de horário Deletado"})
}
