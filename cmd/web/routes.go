package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/shunmasa/bookings/pkg/config"
	"github.com/shunmasa/bookings/pkg/handlers"
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

//image
   fileServer := http.FileServer(http.Dir("./static/"))
   mux.Handle("/static/*",http.StripPrefix("/static",fileServer))
   return mux
}