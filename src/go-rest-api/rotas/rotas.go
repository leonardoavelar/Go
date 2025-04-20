package rotas

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leonardoavelar/Go/src/go-rest-api/controles"
	"github.com/leonardoavelar/Go/src/go-rest-api/middleware"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controles.Home)
	r.HandleFunc("/usuario", controles.GetUsuarios).Methods("Get")
	r.HandleFunc("/usuario/{id}", controles.GetUsuario).Methods("Get")
	r.HandleFunc("/usuario", controles.PostUsuario).Methods("Post")
	r.HandleFunc("/usuario/{id}", controles.DeleteUsuario).Methods("Delete")
	r.HandleFunc("/usuario/{id}", controles.PutUsuario).Methods("PUT")

	log.Fatal((http.ListenAndServe(":8000", r)))
}
