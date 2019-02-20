package Controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/orm"
)

func ExperienceRoutes(controller *Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", controller.GetAllExperience)
	router.Get("/{id}", controller.GetExperience)
	router.Put("/{id}", controller.EditExperience)
	router.Post("/", controller.AddExperience)
	router.Delete("/{id}", controller.DeleteExperience)
	return router
}

func (controller *Controller) GetAllExperience (w http.ResponseWriter, r *http.Request) {
	result := Orm.GetAllExperience(controller.DB, controller.Config)
	render.JSON(w, r, result)
}

func (controller *Controller) GetExperience (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) EditExperience (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) AddExperience (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) DeleteExperience (w http.ResponseWriter, r *http.Request) {

}