package Controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/go-chi/render"

	// "github.com/alistairfink/Personal-Website-V4-Backend/src/models"
)

func PortfolioRoutes(controller *Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", controller.GetAllPortfolio)
	router.Get("/{id}", controller.GetPortfolio)
	router.Put("/{id}", controller.EditPortfolio)
	router.Post("/", controller.AddPortfolio)
	router.Delete("/{id}", controller.DeletePortfolio)
	return router
}

func (controller *Controller) GetAllPortfolio (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) GetPortfolio (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) EditPortfolio (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) AddPortfolio (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) DeletePortfolio (w http.ResponseWriter, r *http.Request) {

}
