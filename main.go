package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rinrey24/ptmmlj-be/database"

	"github.com/rinrey24/ptmmlj-be/controllers"
)

func main() {
	database.ConnectDatabase()
	sqlconn, err := database.DB.DB()
	if err != nil {
		log.Fatal("Something wrong with connection")
	}
	defer sqlconn.Close()
	// r := mux.NewRouter()
	// r.HandleFunc("/login", controllers.Login).Methods("POST")
	// r.HandleFunc("/register", controllers.Register).Methods("POST")
	// r.HandleFunc("/logout", controllers.Logout).Methods("GET")
	// api := r.PathPrefix("/api").Subrouter()
	// api.HandleFunc("/profile", controllers.GetProfile).Methods("GET")
	// api.Use(middlewares.JWTMiddleware)
	//log.Fatal(http.ListenAndServe(":3000", r))

	app := fiber.New()
	app.Post("/login", controllers.Login)
	app.Post("/register", controllers.Register)
	api := app.Group("/api")

	profile := api.Group("/profile")
	profile.Get("/", controllers.GetProfile)
	profile.Get("/:id", controllers.ShowProfile)
	profile.Put("/", controllers.UpdateProfile)
	//api.Use(middlewares.JWTMiddleware)

	stakeholder := api.Group("/stakeholder")
	stakeholder.Get("/", controllers.GetStakeholder)
	stakeholder.Get("/:id", controllers.ShowStakeholder)
	stakeholder.Post("/", controllers.CreateStakeholder)
	stakeholder.Put("/", controllers.UpdateStakeholder)

	app.Listen(":3000")
}
