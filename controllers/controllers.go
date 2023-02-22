package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fmarga/api-go/database"
	"github.com/fmarga/api-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home page")
}

func All(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func Personality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality []models.Personality

	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func NewPersonality(w http.ResponseWriter, r *http.Request) {
	var new models.Personality
	json.NewDecoder(r.Body).Decode(&new)
	database.DB.Create(&new)
	json.NewEncoder(w).Encode(new)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality []models.Personality

	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality []models.Personality

	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(personality)

}
