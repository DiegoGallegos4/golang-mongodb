package config

import (
	"context"
	"fitup/endpoints"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Router *mux.Router
	Db     *mongo.Database
}

func (app *App) Initialize(ctx context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://db:27017"))
	if err != nil {
		log.Fatalf("Couldn't connect to mongodb: %v", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Mongo client couldn't connect with background context: %v", err)
	}

	app.Db = client.Database("fitup")
	app.Router = mux.NewRouter().StrictSlash(false)
	app.setRouters()
}

func (app *App) setRouters() {
	app.Router.HandleFunc("/", hello)
	app.Get("/gyms", app.GetAllGyms)
}

// Handlers
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func (app *App) GetAllGyms(w http.ResponseWriter, r *http.Request) {
	endpoints.GetAllGymsHandler(app.Db, w, r)
}

// HTTP Method Wrappers
func (app *App) Get(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	api := app.Router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc(path, handler).Methods("GET")
}

// Middleware
func loggingHandler(h http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, h)
}

// Run the app on it's router
func (app *App) Run(host string) {
	port := os.Getenv("PORT")
	if port == "" {
		port = host
	}

	middleware := alice.New(loggingHandler, handlers.CompressHandler)
	fmt.Print("Running server in port: " + port)
	if err := http.ListenAndServe(":"+port, middleware.Then(app.Router)); err != nil {
		fmt.Print(err)
	}
}
