package main

import (
	"fmt"
	"log"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	orderAdapters "api/adapters/order"
	userAdapters "api/adapters/user"
	connect_db "api/configs/database"
	Entities "api/entities"
	orderUseCase "api/usecases/order"
	userUseCase "api/usecases/user"
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
	if err := db.AutoMigrate(&Entities.Order{}, &Entities.User{}); err != nil {
		fmt.Println(err)
		return
	}

	orderRepo := orderAdapters.NewGormOrderRepository(db)
	orderService := orderUseCase.NewOrderService(orderRepo)
	orderHandler := orderAdapters.NewHttpOrderHandler(orderService)

	router.POST("/orders", orderHandler.Create)
	router.GET("/orders", orderHandler.GetAll)

	userRepo := userAdapters.NewGormUserRepository(db)
	userService := userUseCase.NewUserService(userRepo)
	userHandler := userAdapters.NewHttpUserHandler(userService)

	router.POST("/users", userHandler.Create)
	router.GET("/users", userHandler.GetAll)
	router.GET("/users/:id", userHandler.GetByID)
	router.PUT("/users/:id", userHandler.Update)
	router.DELETE("/users/:id", userHandler.Delete)

	// Start HTTP server
	PORT := os.Getenv("PORT")
	port := fmt.Sprintf(":%v", PORT)
	fmt.Println("Server Running on Port", port)
	http.ListenAndServe(port, router)
}
