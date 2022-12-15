package controllers

import (
	"net/http"

	"crypto/md5"
	"encoding/hex"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

func GetAllLogin(context *gin.Context) {
	var login []models.Login
	var loginAux []models.Login

	database.Database.Find(&login)

	for _, loginInterator := range login {
		loginAux = append(loginAux, models.Login{Id: loginInterator.Id, Username: loginInterator.Username, Password: ""})
	}
	context.JSON(200, loginAux)
}

func GetByIdLogin(context *gin.Context) {
	var login models.Login
	var loginJson models.Login

	if err := context.ShouldBindJSON(&loginJson); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}
	database.Database.Where(models.Login{Username: loginJson.Username}).FirstOrInit(&login)

	hasher := md5.New()
	hasher.Write([]byte(loginJson.Password))
	loginJson.Password = hex.EncodeToString(hasher.Sum(nil))

	if login.Password != loginJson.Password {
		context.JSON(http.StatusNotFound, gin.H{
			"Not found": "Login/Senha incorreto"})
		return
	}
	context.JSON(http.StatusOK, login)
}

func InsertLogin(context *gin.Context) {
	var login models.Login

	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosLogin(&login); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
		return
	}
	hasher := md5.New()
	hasher.Write([]byte(login.Password))
	login.Password = hex.EncodeToString(hasher.Sum(nil))

	database.Database.Create(&login)
	context.JSON(http.StatusOK, login)
}

func UpdateLogin(context *gin.Context) {
	var login models.Login
	id := context.Params.ByName("id")
	database.Database.First(&login, id)

	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	hasher := md5.New()
	hasher.Write([]byte(login.Password))
	login.Password = hex.EncodeToString(hasher.Sum(nil))

	database.Database.Save(&login)
	context.JSON(http.StatusOK, login)
}

func DeleteLogin(context *gin.Context) {
	var login models.Login
	id := context.Params.ByName("id")
	database.Database.Delete(&login, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Login Deletado"})
}
