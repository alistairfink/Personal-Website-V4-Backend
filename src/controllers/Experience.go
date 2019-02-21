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

func ExperienceRoutes(controller *Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", controller.GetAllExperience)
	router.Get("/{id}", controller.GetExperience)
	router.Put("/api_key/{key}", controller.EditExperience)
	router.Post("/api_key/{key}", controller.AddExperience)
	router.Delete("/{id}/api_key/{key}", controller.DeleteExperience)
	return router
}

func (controller *Controller) GetAllExperience (w http.ResponseWriter, r *http.Request) {
	result := Orm.GetAllExperience(controller.DB, controller.Config)
	if (result == nil) {
		http.Error(w, "Failed to Get Experience Items", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) GetExperience (w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result := Orm.GetExperience(controller.DB, controller.Config, id)
	if (result == nil) {
		http.Error(w, "Failed to Get Experience Item", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) EditExperience (w http.ResponseWriter, r *http.Request) {
	apiKey := chi.URLParam(r, "key")
	if (apiKey == controller.Config.ApiKey) {
 		b, readErr := ioutil.ReadAll(r.Body)
 		defer r.Body.Close()
 		if (readErr != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
 		}

 		var experienceItem Models.Experience
 		err := json.Unmarshal(b, &experienceItem)
 		if (err != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		result := Orm.EditExperience(controller.DB, controller.Config, &experienceItem)
		if (result == nil) {
			http.Error(w, "Error Updating Experience Item", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
		return
	}

	http.Error(w, "Invalid Api Key", http.StatusBadRequest)
}

func (controller *Controller) AddExperience (w http.ResponseWriter, r *http.Request) {
	apiKey := chi.URLParam(r, "key")
	if (apiKey == controller.Config.ApiKey) {
		b, readErr := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if (readErr != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		var experienceItem Models.CreateExperience
		err := json.Unmarshal(b, &experienceItem)
		if (err != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		result := Orm.AddExperience(controller.DB, controller.Config, &experienceItem)
		if (result == nil) {
			http.Error(w, "Error Adding Experience Item", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
		return
	}

	http.Error(w, "Invalid Api Key", http.StatusBadRequest)
}

func (controller *Controller) DeleteExperience (w http.ResponseWriter, r *http.Request) {
	apiKey := chi.URLParam(r, "key")
	if (apiKey == controller.Config.ApiKey) {
		id := chi.URLParam(r, "id")
		result := Orm.DeleteExperience(controller.DB, controller.Config, id)
		if (result) {
			http.Error(w, "", http.StatusOK)
		} else {
			http.Error(w, "Delete Experience Item Failed", http.StatusBadRequest)
		}

		return
	}

	http.Error(w, "Invalid Api Key", http.StatusBadRequest)
}