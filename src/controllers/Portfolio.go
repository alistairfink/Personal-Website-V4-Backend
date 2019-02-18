package Controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/models"
)

func PortfolioRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetAllPortfolio)
	router.Get("/{id}", GetPortfolio)
	router.Put("/{id}", EditPortfolio)
	router.Post("/", AddPortfolio)
	router.Delete("/{id}", DeletePortfolio)
	return router
}

func GetAllPortfolio (w http.ResponseWriter, r *http.Request) {

}

func GetPortfolio (w http.ResponseWriter, r *http.Request) {

}

func EditPortfolio (w http.ResponseWriter, r *http.Request) {

}

func AddPortfolio (w http.ResponseWriter, r *http.Request) {

}

func DeletePortfolio (w http.ResponseWriter, r *http.Request) {

}

