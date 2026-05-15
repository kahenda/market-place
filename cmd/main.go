package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kahenda/marketplace/internal/handler"
	"github.com/kahenda/marketplace/internal/repository"
	"github.com/kahenda/marketplace/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}
	log.Println("Connected to database successfully!")

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Marketplace API is running"})
	})

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server starting on port", port)
	r.Run(":" + port)
}
