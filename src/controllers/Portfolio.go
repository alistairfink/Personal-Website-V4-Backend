package Controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/orm"
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
	result := Orm.GetAllPortfolio(controller.DB, controller.Config)
	if (result == nil) {
		errorResponse := make(map[string]string)
		errorResponse["message"] = "Failed to Get Portfolio Items"
		render.JSON(w, r, errorResponse)
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) GetPortfolio (w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result := Orm.GetPortfolio(controller.DB, controller.Config, id)
	if (result == nil) {
		errorResponse := make(map[string]string)
		errorResponse["message"] = "Failed to Get Portfolio Item"
		render.JSON(w, r, errorResponse)
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) EditPortfolio (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) AddPortfolio (w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) DeletePortfolio (w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result := Orm.DeletePortfolio(controller.DB, controller.Config, id)
	response := make(map[string]string)
	if (result) {
		response["message"] = "Deleted Portfolio Item Successfully"
	} else {
		response["message"] = "Delete Portfolio Item Failed"
	}
	render.JSON(w, r, response)
}
