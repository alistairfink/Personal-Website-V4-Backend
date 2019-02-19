package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/controllers"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/orm"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/config"
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
	config, err := Config.GetConfig()
	if (err != nil) {
		log.Fatal("Config Invalid")
	}

	db := Orm.NewClient(config.Mongo.Url + ":" + config.Mongo.Port)
	router := Routes(db)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":" + config.Port, router))
}