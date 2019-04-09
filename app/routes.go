package app

import (
	"../config"
	"../migrate"
	"./controllers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = migrate.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Post("/profile", a.InputProfile)
	a.Get("/profile", a.ListProfiles)
	a.Get("/profile/{codes:[1-9]+}", a.OneProfile)
	a.Put("/profile/{codes:[1-9]+}", a.UpdateProfile)
	a.Delete("/profile/{codes:[1-9]+}", a.DeletedProfile)
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("Get")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("Put")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("Delete")
}

func (a *App) InputProfile(w http.ResponseWriter, r *http.Request) {
	controllers.InputProfile(a.DB, w, r)
}

func (a *App) ListProfiles(w http.ResponseWriter, r *http.Request) {
	controllers.ListProfiles(a.DB, w, r)
}

func (a *App) OneProfile(w http.ResponseWriter, r *http.Request) {
	controllers.OneProfile(a.DB, w, r)
}

func (a *App) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateProfile(a.DB, w, r)
}

func (a *App) DeletedProfile(w http.ResponseWriter, r *http.Request) {
	controllers.DeletedProfile(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}