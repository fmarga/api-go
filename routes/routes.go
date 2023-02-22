package routes

import (
	"log"
	"net/http"

	"github.com/fmarga/api-go/controllers"
	"github.com/fmarga/api-go/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMid)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/all", controllers.All).Methods("Get")
	r.HandleFunc("/api/all/{id}", controllers.Personality).Methods("Get")
	r.HandleFunc("/api/all", controllers.NewPersonality).Methods("Post")
	r.HandleFunc("/api/all/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/all/{id}", controllers.UpdatePersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
