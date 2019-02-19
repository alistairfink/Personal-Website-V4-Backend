package Controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/go-chi/render"

	// "github.com/alistairfink/Personal-Website-V4-Backend/src/models"
)

func ExperienceRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetAllExperience)
	router.Get("/{id}", GetExperience)
	router.Put("/{id}", EditExperience)
	router.Post("/", AddExperience)
	router.Delete("/{id}", DeleteExperience)
	return router
}

func GetAllExperience (w http.ResponseWriter, r *http.Request) {

}

func GetExperience (w http.ResponseWriter, r *http.Request) {

}

func EditExperience (w http.ResponseWriter, r *http.Request) {

}

func AddExperience (w http.ResponseWriter, r *http.Request) {

}

func DeleteExperience (w http.ResponseWriter, r *http.Request) {

}