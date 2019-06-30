package main

import (
	"log"
	"net/http"

	"github.com/adiachenko/golang-demo/cmd/http/routes"
	"github.com/adiachenko/golang-demo/config"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	router := chi.NewRouter()

	router.Use(cors.New(config.CORS).Handler)

	// router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.StripSlashes)
	router.Use(chiMiddleware.Recoverer)

	log.Fatal(http.ListenAndServe(":8000", routes.Register(router)))
}
