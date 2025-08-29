package main

import (
	"log"
	"os"

	"github.com/dokkiichan/BridgeMe-Back/infrastructure/datastore"
	"github.com/dokkiichan/BridgeMe-Back/interfaces/controllers"
	"github.com/dokkiichan/BridgeMe-Back/interfaces/generated"
	"github.com/dokkiichan/BridgeMe-Back/interfaces/repositories"
	"github.com/dokkiichan/BridgeMe-Back/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}

	client, db, err := datastore.NewDB(mongoURI)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer datastore.CloseDB(client)

	profileRepository := repositories.NewProfileRepository(db)
	profileInteractor := usecase.NewProfileInteractor(profileRepository)
	profileController := controllers.NewProfileController(profileInteractor)

	e := echo.New()

	generated.RegisterHandlers(e, profileController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}