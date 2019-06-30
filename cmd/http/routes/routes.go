package routes

import (
	"github.com/adiachenko/golang-demo/cmd/http/actions"
	"github.com/go-chi/chi"
)

func Register(router *chi.Mux) *chi.Mux {
	router.Get("/", actions.Home)

	return router
}
