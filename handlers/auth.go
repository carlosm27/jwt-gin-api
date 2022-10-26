package handlers

import (
	"github.com/carlosm27/jwtGinApi/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Username: input.Username, Password: input.Password}

	db, err := models.Setup()
	if err != nil {
		log.Println(err)
	}
	userRepo := models.RepoInterface(db)

	if _, err = userRepo.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}
