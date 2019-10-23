package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"rivendell/config"
	"rivendell/app/handler"
)

// App has router and db instances
type App struct {
	Router *mux.Router
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Post("/create-cluster", a.handleRequest(handler.GreetHandler))
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// RequestHandlerFunction type definition
type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}