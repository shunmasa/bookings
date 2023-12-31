package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/shunmasa/bookings/internal/config"
	"github.com/shunmasa/bookings/internal/handlers"
)


func routes(app *config.AppConfig) http.Handler{
//    mux :=pat.New()
   
//    mux.Get("/",http.HandlerFunc(handlers.Repo.Home))
//    mux.Get("/about",http.HandlerFunc(handlers.Repo.About))
//    return mux

   mux:= chi.NewRouter()

   mux.Use(middleware.Recoverer)
   mux.Use(NoSurf)
   mux.Use(SessionLoad)
//    mux.Use(writeToConsole)
   //middlewhere

   mux.Get("/",handlers.Repo.Home)
   mux.Get("/about",handlers.Repo.About)
   mux.Get("/generals-quarters",handlers.Repo.Generals)
   mux.Get("/majors-suite",handlers.Repo.Majors)
   mux.Get("/search-availability",handlers.Repo.Availability)
   mux.Post("/search-availability",handlers.Repo.PostAvailability)
   mux.Post("/search-availability-json",handlers.Repo.AvailabilityJSON)
   mux.Get("/contact",handlers.Repo.Contact)
   mux.Get("/make-reservation",handlers.Repo.Reservation)
//image
   fileServer := http.FileServer(http.Dir("./static/"))
   mux.Handle("/static/*",http.StripPrefix("/static",fileServer))
   return mux
}