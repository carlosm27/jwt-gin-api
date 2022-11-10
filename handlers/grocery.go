package handlers

import (
    "github.com/carlosm27/jwtGinApi/models"
	"github.com/gin-gonic/gin"
	"net/http"

)


type NewGrocery struct {
    Name     string `json:"name" binding:"required"`
    Quantity int    `json:"quantity" binding:"required"`
}



func (s *Server) GetGroceries(c *gin.Context) {

    var groceries []models.Grocery


    if err := s.db.Find(&groceries).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, groceries)

}


func (s *Server) PostGrocery(c *gin.Context) {

    var grocery NewGrocery

    if err := c.ShouldBindJSON(&grocery); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newGrocery := models.Grocery{Name: grocery.Name, Quantity: grocery.Quantity}


    if err := s.db.Create(&newGrocery).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, newGrocery)
}