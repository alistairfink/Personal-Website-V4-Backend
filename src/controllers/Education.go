package Controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/go-chi/render"

	// "github.com/alistairfink/Personal-Website-V4-Backend/src/models"
)

func EducationRoutes(controller *Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", controller.GetAllEducation)
	router.Get("/{id}", controller.GetEducation)
	router.Put("/{id}", controller.EditEducation)
	router.Post("/", controller.AddEducation)
	router.Delete("/{id}", controller.DeleteEducation)
	return router
}

func (controller *Controller) GetAllEducation (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) GetEducation (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) EditEducation (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) AddEducation (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) DeleteEducation (w http.ResponseWriter, r *http.Request) {

}