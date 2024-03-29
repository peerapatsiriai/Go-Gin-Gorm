package main

import (
	"fmt"
	"log"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	adapters "api/adapters"
	connect_db "api/configs/database"
	entities "api/entities"
	usecases "api/usecases"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	db := connect_db.ConnectDB()

	// Initialize Gin router
	router := gin.New()

	// Run database migrations
	if err := db.AutoMigrate(&entities.Order{}); err != nil {
		fmt.Println(err)
		return
	}

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := usecases.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	router.POST("/orders", orderHandler.Create)
	router.GET("/orders", orderHandler.GetAll)

	// Start HTTP server
	PORT := os.Getenv("PORT")
	port := fmt.Sprintf(":%v", PORT)
	fmt.Println("Server Running on Port", port)
	http.ListenAndServe(port, router)
}
