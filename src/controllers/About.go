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

func AboutRoutes(controller *Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", controller.GetAbout)
	router.Put("/api_key/{key}", controller.EditAbout)
	return router
}

func (controller *Controller) GetAbout (w http.ResponseWriter, r *http.Request) {
	result := Orm.GetAbout(controller.DB, controller.Config)
	if (result == nil) {
		http.Error(w, "Failed to Get About", http.StatusBadRequest)
	}

	render.JSON(w, r, result)
}

func (controller *Controller) EditAbout (w http.ResponseWriter, r *http.Request) {
	apiKey := chi.URLParam(r, "key")
	if (apiKey == controller.Config.ApiKey) {
		b, readErr := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if (readErr != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		var aboutEdit Models.About
		err := json.Unmarshal(b, &aboutEdit)
		if (err != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		result := Orm.EditAbout(controller.DB, controller.Config, &aboutEdit)
		if (result == nil) {
			http.Error(w, "Error Updating About", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
		return
	}

	http.Error(w, "Invalid Api Key", http.StatusBadRequest)
}