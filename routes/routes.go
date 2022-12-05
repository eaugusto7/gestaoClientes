package routes

import (
	"log"
	"net/http"

	"github.com/eaugusto7/gestaoClientes/controllers"
	"github.com/eaugusto7/gestaoClientes/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)

	//Create
	r.HandleFunc("/api/v1/clientes/insert", controllers.InsertClient).Methods("Post")

	//Read
	r.HandleFunc("/api/v1/clientes/getAll", controllers.GetAll).Methods("Get")
	r.HandleFunc("/api/v1/clientes/{id}", controllers.GetById).Methods("Get")

	//Update
	r.HandleFunc("/api/v1/clientes/{id}", controllers.UpdateClient).Methods("Put")

	//Delete
	r.HandleFunc("/api/v1/clientes/{id}", controllers.DeleteClient).Methods("Delete")

	r.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
