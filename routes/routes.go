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
	r.POST("/api/v1/clientes/insert", controllers.InsertClient)
	r.GET("/api/v1/clientes/getAll", controllers.GetAllClientes)
	r.GET("/api/v1/clientes/:id", controllers.GetByIdClientes)
	r.PATCH("/api/v1/clientes/:id", controllers.UpdateClient)
	r.DELETE("/api/v1/clientes/:id", controllers.DeleteClient)

	//CRUD - Servicos
	/*r.HandleFunc("/api/v1/servicos/insert", controllers.InsertServicos).Methods("Post")
	r.HandleFunc("/api/v1/servicos/getAll", controllers.GetAllServicos).Methods("Get")
	r.HandleFunc("/api/v1/servicos/{id}", controllers.GetByIdServicos).Methods("Get")
	r.HandleFunc("/api/v1/servicos/{id}", controllers.UpdateServicos).Methods("Put")
	r.HandleFunc("/api/v1/servicos/{id}", controllers.DeleteServicos).Methods("Delete")*/

	//CRUD - Atendimentos
	r.POST("/api/v1/atendimentos/insert", controllers.InsertAtendimentos)
	r.GET("/api/v1/atendimentos/getAll", controllers.GetAllAtendimentos)
	r.GET("/api/v1/atendimentos/:id", controllers.GetByIdAtendimentos)
	r.PATCH("/api/v1/atendimentos/:id", controllers.UpdateAtendimentos)
	r.DELETE("/api/v1/atendimentos/:id", controllers.DeleteAtendimento)

	//r.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
