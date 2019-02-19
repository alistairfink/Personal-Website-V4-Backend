package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/controllers"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/orm"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func Routes(db *mongo.Client) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/alistairfink", func(r chi.Router) {
		r.Mount("/test", Controllers.Routes())
	})

	return router
}

func main() {
	db := Orm.NewClient("mongodb://localhost:27017")
	router := Routes(db)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":41692", router))
}