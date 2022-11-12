package handlers

import (
    "github.com/carlosm27/jwtGinApi/models"
    "github.com/carlosm27/jwtGinApi/utils"
	"github.com/gin-gonic/gin"
	"net/http"

)


type NewGrocery struct {
    Name     string `json:"name" binding:"required"`
    Quantity int    `json:"quantity" binding:"required"`
}



func (s *Server) GetGroceries(c *gin.Context) {

    user, err := utils.CurrentUser(c)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user.Groceries})

}


func (s *Server) PostGrocery(c *gin.Context) {

    var grocery models.Grocery

    if err := c.ShouldBindJSON(&grocery); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := utils.CurrentUser(c)
    if err != nil {

        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    grocery.UserId = user.ID

    if err := s.db.Create(&grocery).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, grocery)
}