package Controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/go-chi/render"

	// "github.com/alistairfink/Personal-Website-V4-Backend/src/models"
)

func EducationRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetAllEducation)
	router.Get("/{id}", GetEducation)
	router.Put("/{id}", EditEducation)
	router.Post("/", AddEducation)
	router.Delete("/{id}", DeleteEducation)
	return router
}

func GetAllEducation (w http.ResponseWriter, r *http.Request) {

}

func GetEducation (w http.ResponseWriter, r *http.Request) {

}

func EditEducation (w http.ResponseWriter, r *http.Request) {

}

func AddEducation (w http.ResponseWriter, r *http.Request) {

}

func DeleteEducation (w http.ResponseWriter, r *http.Request) {

}