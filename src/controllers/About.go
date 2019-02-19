package Controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/go-chi/render"

	// "github.com/alistairfink/Personal-Website-V4-Backend/src/models"
)

func AboutRoutes(controller *Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", controller.GetAbout)
	router.Put("/", controller.EditAbout)
	return router
}

func (controller *Controller) GetAbout (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) EditAbout (w http.ResponseWriter, r *http.Request) {
	
}