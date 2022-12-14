package routes

import (
	"log"
	"net/http"

	"github.com/eaugusto7/gestaoClientes/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

func HandleRequest() {
	r := gin.Default()
	r.Static("/assets", "./assets")

	//CRUD Clientes
	r.POST("/api/v1/clientes", controllers.InsertClient)
	r.GET("/api/v1/clientes", controllers.GetAllClientes)
	r.GET("/api/v1/clientes/:id", controllers.GetByIdClientes)
	r.PATCH("/api/v1/clientes/:id", controllers.UpdateClient)
	r.DELETE("/api/v1/clientes/:id", controllers.DeleteClient)

	//CRUD - Servicos
	r.POST("/api/v1/servicos", controllers.InsertServicos)
	r.GET("/api/v1/servicos", controllers.GetAllServicos)
	r.GET("/api/v1/servicos/:id", controllers.GetByIdServicos)
	r.PATCH("/api/v1/servicos/:id", controllers.UpdateServicos)
	r.DELETE("/api/v1/servicos/:id", controllers.DeleteServicos)

	//CRUD - Atendimentos
	r.POST("/api/v1/atendimentos", controllers.InsertAtendimentos)
	r.GET("/api/v1/atendimentos", controllers.GetAllAtendimentos)
	r.GET("/api/v1/atendimentos/:id", controllers.GetByIdAtendimentos)
	r.PATCH("/api/v1/atendimentos/:id", controllers.UpdateAtendimentos)
	r.DELETE("/api/v1/atendimentos/:id", controllers.DeleteAtendimento)

	//CRUD - Atendente
	r.POST("/api/v1/atendentes", controllers.InsertAtendente)
	r.GET("/api/v1/atendentes", controllers.GetAllAtendente)
	r.GET("/api/v1/atendentes/:id", controllers.GetAtendenteById)
	r.PATCH("/api/v1/atendentes/:id", controllers.UpdateAtendente)
	r.DELETE("/api/v1/atendentes/:id", controllers.DeleteAtendente)

	//r.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
