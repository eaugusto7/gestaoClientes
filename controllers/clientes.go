package controllers

import (
	"fmt"
	"net/http"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

// GetAllClientes godoc
// @Summary Todos Clientes
// @Description Obtem todos os clientes vindos do banco de dados
// @Tags Clientes
// @Produce json
// @Sucess 200 [object] model.Cliente
// @Router /api/v1/clientes [get]
func GetAllClientes(context *gin.Context) {
	var clientes []models.Cliente
	database.Database.Find(&clientes)
	context.JSON(200, clientes)
}

// GetClienteById godoc
// @Summary Busca Cliente por Id
// @Description Obtem  o json de um determinado cliente, filtrado por id
// @Tags Clientes
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Cliente
// @Failure 404 {object} httputil.HTTPError "Cliente não encontrado"
// @Router /api/v1/clientes/{id} [get]
func GetClienteById(context *gin.Context) {
	var cliente models.Cliente
	id := context.Params.ByName("id")
	database.Database.First(&cliente, id)

	if cliente.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Message": "Cliente não encontrado"})
		return
	}
	context.JSON(http.StatusOK, cliente)
}

// GetClienteByName godoc
// @Summary Busca Cliente por Nome
// @Description Obtem  o json de um determinado cliente, filtrado por nome
// @Tags Clientes
// @Produce json
// @Param   nome     path    string     true        "Nome"
// @Sucess 200 {object} model.Cliente
// @Failure 404 {object} httputil.HTTPError "Cliente não encontrado"
// @Router /api/v1/clientes/nome/{nome} [get]
func GetClienteByName(context *gin.Context) {
	var cliente models.Cliente
	nome := context.Params.ByName("nome")
	//database.Database.First(&cliente, nome)

	fmt.Println(nome)

	database.Database.Where("nome like ?", "%"+nome+"%").Find(&cliente)
	//database.Database.Where("nome = ?", nome).Find(&cliente)

	if cliente.Id == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Message": "Cliente não encontrado"})
		return
	}
	context.JSON(http.StatusOK, cliente)
}

// InsertCliente godoc
// @Summary Insere Cliente
// @Description Cria um novo cliente no banco de dados
// @Tags Clientes
// @Accept json
// @Produce json
// @Param   cliente     body    models.Cliente     true		"Json Cliente"
// @Sucess 200 {object} model.Cliente
// @Failure 404 {object} httputil.HTTPError "Erro: Cliente não cadastrado"
// @Router /api/v1/clientes [post]
func InsertCliente(context *gin.Context) {
	var cliente models.Cliente
	if error := context.ShouldBindJSON(&cliente); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message error: ": error.Error()})
		return
	}

	if err := models.ValidaDadosClientes(&cliente); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message error: ": err.Error()})
		return
	}
	database.Database.Create(&cliente)
	context.JSON(http.StatusOK, cliente)
}

// UpdateCliente godoc
// @Summary Atualiza Cliente
// @Description Atualiza as informações de um determinado cliente no banco de dados
// @Tags Clientes
// @Accept json
// @Produce json
// @Param   cliente     body    models.Cliente     true		"Json Cliente"
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object} model.Cliente
// @Failure 400 {object} httputil.HTTPError "Erro: Cliente não existe"
// @Router /api/v1/clientes/{id} [patch]
func UpdateCliente(context *gin.Context) {
	var cliente models.Cliente
	id := context.Params.ByName("id")
	database.Database.First(&cliente, id)

	if error := context.ShouldBindJSON(&cliente); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message error: ": error.Error()})
		return
	}

	if cliente.Id == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message error: ": "Cliente não existe"})
		return
	}

	database.Database.Save(&cliente)
	context.JSON(http.StatusOK, cliente)
}

// DeleteCliente godoc
// @Summary Deleta Cliente
// @Description Remove o cliente indicado pelo id no banco de dados
// @Tags Clientes
// @Produce json
// @Param   id     path    int     true        "Id"
// @Sucess 200 {object}
// @Failure 400 {object} httputil.HTTPError "Erro: Não encontrado"
// @Router /api/v1/clientes/{id} [delete]
func DeleteCliente(context *gin.Context) {
	var cliente models.Cliente
	id := context.Params.ByName("id")
	database.Database.Delete(&cliente, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Cliente Deletado"})
}
