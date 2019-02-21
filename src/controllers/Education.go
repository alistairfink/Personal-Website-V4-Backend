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

func EducationRoutes(controller *Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", controller.GetAllEducation)
	router.Get("/{id}", controller.GetEducation)
	router.Put("/api_key/{key}", controller.EditEducation)
	router.Post("/api_key/{key}", controller.AddEducation)
	router.Delete("/{id}/api_key/{key}", controller.DeleteEducation)
	return router
}

func (controller *Controller) GetAllEducation (w http.ResponseWriter, r *http.Request) {
	result := Orm.GetAllEducation(controller.DB, controller.Config)
	if (result == nil) {
		http.Error(w, "Failed to Get Education", http.StatusBadRequest)
	}

	render.JSON(w, r, result)
}

func (controller *Controller) GetEducation (w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result := Orm.GetEducation(controller.DB, controller.Config, id)
	if (result == nil) {
		http.Error(w, "Failed to Get Portfolio Item", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) EditEducation (w http.ResponseWriter, r *http.Request) {
	apiKey := chi.URLParam(r, "key")
	if (apiKey == controller.Config.ApiKey) {
		b, readErr := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if (readErr != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		var education Models.Education
		err := json.Unmarshal(b, &education)
		if (err != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		result := Orm.EditEducation(controller.DB, controller.Config, &education)
		if (result == nil) {
			http.Error(w, "Error Updating Education Item", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
		return
	}

	http.Error(w, "Invalid Api Key", http.StatusBadRequest)
}

func (controller *Controller) AddEducation (w http.ResponseWriter, r *http.Request) {
	apiKey := chi.URLParam(r, "key")
	if (apiKey == controller.Config.ApiKey) {
		b, readErr := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if (readErr != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		var education Models.CreateEducation
		err := json.Unmarshal(b, &education)
		if (err != nil) {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		result := Orm.AddEducation(controller.DB, controller.Config, &education)
		if (result == nil) {
			http.Error(w, "Error Adding Education Item", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
		return
	}

	http.Error(w, "Invalid Api Key", http.StatusBadRequest)
}

func (controller *Controller) DeleteEducation (w http.ResponseWriter, r *http.Request) {
	apiKey := chi.URLParam(r, "key")
	if (apiKey == controller.Config.ApiKey) {
		id := chi.URLParam(r, "id")
		result := Orm.DeleteEducation(controller.DB, controller.Config, id)
		if (result) {
			http.Error(w, "", http.StatusOK)
		} else {
			http.Error(w, "Delete Education Item Failed", http.StatusBadRequest)
		}
		
		return
	}

	http.Error(w, "Invalid Api Key", http.StatusBadRequest)
}