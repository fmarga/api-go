package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fmarga/api-go/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home page")
}

func All(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
