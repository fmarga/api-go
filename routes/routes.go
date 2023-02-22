package routes

import (
	"log"
	"net/http"

	"github.com/fmarga/api-go/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/all", controllers.All)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
