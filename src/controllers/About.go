package Controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/models"
)

func AboutRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetAbout)
	router.Put("/", EditAbout)
	return router
}

func GetAbout (w http.ResponseWriter, r *http.Request) {

}

func EditAbout (w http.ResponseWriter, r *http.Request) {

}