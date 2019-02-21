package Controllers

import (
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/orm"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/models"
)

func PortfolioRoutes(controller *Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", controller.GetAllPortfolio)
	router.Get("/{id}", controller.GetPortfolio)
	router.Put("/api_key/{key}", controller.EditPortfolio)
	router.Post("/api_key/{key}", controller.AddPortfolio)
	router.Delete("/{id}/api_key/{key}", controller.DeletePortfolio)
	return router
}

func (controller *Controller) GetAllPortfolio (w http.ResponseWriter, r *http.Request) {
	result := Orm.GetAllPortfolio(controller.DB, controller.Config)
	if (result == nil) {
		http.Error(w, "Failed to Get Portfolio Items", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) GetPortfolio (w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result := Orm.GetPortfolio(controller.DB, controller.Config, id)
	if (result == nil) {
		http.Error(w, "Failed to Get Portfolio Item", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) EditPortfolio (w http.ResponseWriter, r *http.Request) {
	apiKey := chi.URLParam(r, "key")
	if (apiKey == controller.Config.ApiKey) {
		b, readErr := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if (readErr != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		var portfolioItem Models.Portfolio
		err := json.Unmarshal(b, &portfolioItem)
		if (err != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		result := Orm.EditPortfolio(controller.DB, controller.Config, &portfolioItem)
		if (result == nil) {
			http.Error(w, "Error Updating Portfolio Item", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
		return 
	}

	http.Error(w, "Invalid Api Key", http.StatusBadRequest)
}

func (controller *Controller) AddPortfolio (w http.ResponseWriter, r *http.Request) {
	apiKey := chi.URLParam(r, "key")
	if (apiKey == controller.Config.ApiKey) {
		b, readErr := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if (readErr != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		var portfolioItem Models.CreatePortfolio
		err := json.Unmarshal(b, &portfolioItem)
		if (err != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		result := Orm.AddPortfolio(controller.DB, controller.Config, &portfolioItem)
		if (result == nil) {
			http.Error(w, "Error Adding Portfolio Item", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
		return 
	}

	http.Error(w, "Invalid Api Key", http.StatusBadRequest)
}

func (controller *Controller) DeletePortfolio (w http.ResponseWriter, r *http.Request) {
	apiKey := chi.URLParam(r, "key")
	if (apiKey == controller.Config.ApiKey) {
		id := chi.URLParam(r, "id")
		result := Orm.DeletePortfolio(controller.DB, controller.Config, id)
		if (result) {
			http.Error(w, "", http.StatusOK)
		} else {
			http.Error(w, "Delete Portfolio Item Failed", http.StatusBadRequest)
		}
		return 
	}

	http.Error(w, "Invalid Api Key", http.StatusBadRequest)
}
