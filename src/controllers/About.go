package Controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/orm"
)

func AboutRoutes(controller *Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", controller.GetAbout)
	router.Put("/", controller.EditAbout)
	return router
}

func (controller *Controller) GetAbout (w http.ResponseWriter, r *http.Request) {
	result := Orm.GetAbout(controller.DB, controller.Config)
	render.JSON(w, r, result)
}

func (controller *Controller) EditAbout (w http.ResponseWriter, r *http.Request) {
	
}