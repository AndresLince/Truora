package main

import (
	//...
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func routers() *chi.Mux {
	router := chi.NewRouter()
	cors := Cors()
	router.Use(cors.Handler)

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Post("/createRSAKey", createRSAKey)
	router.Get("/getRSAKeys", getRSAKeys)
	router.Get("/getRSAKeys/{NombreLlave}", getRSAKey)
	router.Post("/encrypt", encryptPlainText)
	router.Post("/decrypt", decryptPlainText)

	return router
}

func Cors() *cors.Cors {

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	return cors
}
