package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rafasuzuki/students-manager/configs"
	"github.com/rafasuzuki/students-manager/handlers"
)

func Initialize() {
	r := chi.NewRouter()

	r.Post("/students", handlers.Create)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
