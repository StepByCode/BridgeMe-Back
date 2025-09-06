package main

import (
	"log"
	"os"

	"github.com/dokkiichan/BridgeMe-Back/internal/infrastructure/datastore"
	"github.com/dokkiichan/BridgeMe-Back/internal/interfaces/controllers"
	"github.com/dokkiichan/BridgeMe-Back/internal/interfaces/generated"
	"github.com/dokkiichan/BridgeMe-Back/internal/interfaces/repository"
	"github.com/dokkiichan/BridgeMe-Back/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	profileRepository := repository.NewProfileRepository(db)
	profileUseCase := usecase.NewProfileUseCase(profileRepository)
	profileController := controllers.NewProfileController(profileUseCase)

	e := echo.New()

	// CORSミドルウェアの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3003", "https://bridgeme.dokkiitech.dev"}, // 許可するオリジン
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

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
