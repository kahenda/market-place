package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kahenda/marketplace/internal/handler"
	"github.com/kahenda/marketplace/internal/middleware"
	"github.com/kahenda/marketplace/internal/repository"
	"github.com/kahenda/marketplace/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

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

	listingRepo := repository.NewListingRepository(db)
	listingSvc := service.NewListingService(listingRepo)
	listingHandler := handler.NewListingHandler(listingSvc)

	messageRepo := repository.NewMessageRepository(db)
	messageSvc := service.NewMessageService(messageRepo)
	messageHandler := handler.NewMessageHandler(messageSvc)

	reviewRepo := repository.NewReviewRepository(db)
	reviewSvc := service.NewReviewService(reviewRepo)
	reviewHandler := handler.NewReviewHandler(reviewSvc)

	imageRepo := repository.NewImageRepository(db)
	imageSvc := service.NewImageService(imageRepo)
	imageHandler := handler.NewImageHandler(imageSvc)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	r.MaxMultipartMemory = 8 << 20

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Kisumu Marketplace API is running"})
	})

	r.GET("/options", listingHandler.GetOptions)
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.POST("/listings", listingHandler.CreateListing)
		protected.GET("/listings", listingHandler.GetListings)
		protected.GET("/listing/:id", listingHandler.GetListingByID)

		protected.POST("/listings/:listing_id/images", imageHandler.UploadImage)
		protected.GET("/listings/:listing_id/images", imageHandler.GetImages)

		protected.POST("/messages", messageHandler.SendMessage)
		protected.GET("/messages", messageHandler.GetConversation)
		protected.GET("/messages/inbox", messageHandler.GetInbox)

		protected.POST("/reviews", reviewHandler.CreateReview)
		protected.GET("/reviews/:seller_id", reviewHandler.GetReviewsBySeller)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server starting on port", port)
	r.Run(":" + port)
}
