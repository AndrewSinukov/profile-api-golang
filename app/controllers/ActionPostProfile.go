package controllers

import (
	"../models"
	"../utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func InputProfile(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var project models.Profile
	err := decoder.Decode(&project)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	project.Codes = utils.GenerateId()
	err = models.InsertProfile(db, &project)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}

func UpdateProfile(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var profiles models.Profile
	id, _ := strconv.Atoi(vars["codes"])
	err := models.OneProfileGetting(db, id, &profiles)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&profiles)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	err = models.UpdateProfile(db, &profiles)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, profiles)
}

func DeletedProfile(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var profiles models.Profile
	id, _ := strconv.Atoi(vars["codes"])
	err := models.OneProfileGetting(db, id, &profiles)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = models.DeletedProfile(db, &profiles)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, profiles)
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	fmt.Println("status ", status)
	var res utils.ResponseData

	res.Status = status
	res.Meta = utils.ResponseMessage(status)
	res.Data = payload

	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, status int, message string) {
	var res utils.ResponseData
	resCode := utils.ResponseMessage(status)
	res.Status = status
	res.Meta = resCode
	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}