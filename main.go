package main

import (
	"chaos/backend/config"
	"chaos/backend/controller"
	"chaos/backend/database"
	"chaos/backend/job"
	"chaos/backend/middleware"
	"chaos/backend/seed"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/carlescere/scheduler"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	// init every necessary service
	log.Println("Initializing Config")
	config.Init("config.toml")
	log.Println("Initializing Database")
	database.Init()
	log.Println("Auto migration")
	database.AutoMigration()

	// by running go run main seed
	// it will initialize the database
	argLength := len(os.Args[1:])
	if argLength != 0 {
		if os.Args[1] == "seed" {
			log.Println("Seeding Database")
			seed.SeedDatabase()
			log.Println("Finished Seeding Database")
			return
		}
	}

	// starting program
	go startScheduler()
	startServer()
}

func startServer() {
	router := getServer()

	port := config.Data.App.Port
	log.Printf("connect to http://localhost:%s/ for Backend Service", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func startScheduler() {
	log.Println("Starting Scheduler")
	scheduler.Every(1).Minutes().NotImmediately().Run(job.FetchBTCUSDPair)
	runtime.Goexit()
}

func getServer() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{config.Data.App.AllowedOrigin},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Restful API handler
	// auth controller
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", controller.Login)
	})

	// apikey controller
	router.Route("/apikey", func(r chi.Router) {
		r.With(middleware.AuthMiddleware()).Route("/", func(r chi.Router) {
			r.Post("/", controller.GenerateAPIkey)
		})
	})

	// price controller
	router.Route("/price", func(r chi.Router) {
		r.With(middleware.APIMiddleware()).Route("/", func(r chi.Router) {
			r.Get("/last", controller.GetPriceByLatest)
			r.Get("/bydate/{date}", controller.GetPriceByTime)
			r.Get("/range/{start}/{end}", controller.GetAverageByRange)
		})
	})

	return router
}
