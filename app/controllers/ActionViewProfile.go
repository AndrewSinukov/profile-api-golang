package controllers

import (
	"../models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func ListProfiles(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	profiles := []models.Profile{}
	err := models.GetAllProfile(db, &profiles)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, profiles)
}

func OneProfile(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var profiles models.Profile
	id, _ := strconv.Atoi(vars["codes"])
	err := models.OneProfileGetting(db, id, &profiles)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		respondJSON(w, http.StatusOK, profiles)
		return
	}
}