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
	//Public access
	app.Post("/login", controllers.Login)
	app.Post("/register", controllers.Register)
	app.Get("/profile/:id", controllers.ShowActiveProfile)
	app.Get("/stakeholder", controllers.GetActiveStakeholder)
	app.Get("/history", controllers.GetActiveHistory)
	app.Get("/article", controllers.GetActiveArticle)

	//Profile handler
	api := app.Group("/api")
	profile := api.Group("/profile")
	profile.Get("/", controllers.GetProfile)
	profile.Get("/:id", controllers.ShowProfile)
	profile.Put("/", controllers.UpdateProfile)
	//api.Use(middlewares.JWTMiddleware)

	//Stakeholder handler
	stakeholder := api.Group("/stakeholder")
	stakeholder.Get("/", controllers.GetStakeholder)
	stakeholder.Get("/:id", controllers.ShowStakeholder)
	stakeholder.Post("/", controllers.CreateStakeholder)
	stakeholder.Put("/", controllers.UpdateStakeholder)

	//History handler
	history := api.Group("/history")
	history.Get("/", controllers.GetHistory)
	history.Get("/:id", controllers.ShowHistory)
	history.Post("/", controllers.CreateHistory)
	history.Put("/", controllers.UpdateHistory)

	//Article handler
	article := api.Group("/article")
	article.Get("/", controllers.GetArticle)
	article.Get("/:id", controllers.ShowArticle)
	article.Post("/", controllers.CreateArticle)
	article.Put("/", controllers.UpdateArticle)

	app.Listen(":3000")
}
