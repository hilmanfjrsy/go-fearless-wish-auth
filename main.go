package main

import (
	"fmt"
	"go-todo-app/config"
	"go-todo-app/controllers"
	"go-todo-app/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	config.ConnectDatabase()
	model.AutoMigrate()
	r := gin.Default()
	fmt.Println("Ipun")
	v1 := r.Group("/api/v1")
	v1.POST("/login", controllers.Login)
	v1.POST("/register", controllers.Register)
	r.Run()
}
