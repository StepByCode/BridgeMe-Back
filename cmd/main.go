package main

import (
	"log"
	"os"

	"github.com/dokkiichan/BridgeMe-Back/internal/infrastructure/datastore"
	"github.com/dokkiichan/BridgeMe-Back/internal/interfaces/controllers"
	"github.com/dokkiichan/BridgeMe-Back/internal/interfaces/generated"
	"github.com/dokkiichan/BridgeMe-Back/internal/interfaces/repository"
	"github.com/dokkiichan/BridgeMe-Back/internal/interfaces/middleware"
	"github.com/dokkiichan/BridgeMe-Back/internal/usecase"
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

	profileRepository := repository.NewProfileRepository(db)
	profileUseCase := usecase.NewProfileUseCase(profileRepository)
	profileController := controllers.NewProfileController(profileUseCase)

	e := echo.New()

	authMiddleware, err := middleware.NewAuthMiddleware()
	if err != nil {
		log.Fatal("Auth0ミドルウェアの初期化に失敗しました:", err)
	}

	// プロファイル関連のルートにミドルウェアを適用
	profileGroup := e.Group("/profiles")
	profileGroup.Use(middleware.Auth0EchoMiddleware(authMiddleware))

	generated.RegisterHandlers(profileGroup, profileController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
