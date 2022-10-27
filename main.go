package main

import (
	"github.com/carlosm27/jwtGinApi/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")

	r := SetupRouter()

	if r.Run(":"+port) != nil {
		log.Printf("Error running at port: %s", port)
	}
}

func DbInit() *gorm.DB {
	db, err := models.Setup()
	if err != nil {
		log.Println("Problem setting up database")
	}
	return db
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := DbInit()

	server := NewServer(db)

	group := r.Group("/api")

	group.POST("/register", server.Register)
	group.POST("/login", server.Login)

	protectedEndpoints := r.Group("/api/admin")
	protectedEndpoints.Use(JwtAuthMiddleware())
	protectedEndpoints.GET("/user", server.CurrentUser)

	return r

}
