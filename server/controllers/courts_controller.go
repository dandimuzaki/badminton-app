package controllers

import (
	"net/http"

	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/models"
	"github.com/gin-gonic/gin"
)

func CreateCourt(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Price    float64 `json:"price"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	court := models.Court{
		Name: body.Name, 
		Location: body.Location, 
		Price: body.Price,
	}
	if err := initializers.DB.Create(&court).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Court already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Court created successfully", "data": court})
}

func GetCourts(c *gin.Context) {
	var courts []models.Court
	initializers.DB.Find(&courts)
	c.JSON(http.StatusOK, gin.H{"data": courts})
}

func GetCourtByID(c *gin.Context) {
	id := c.Param("id")
	var court models.Court
	if err := initializers.DB.First(&court, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Court not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": court})
}

func UpdateCourt(c *gin.Context) {
	id := c.Param("id")

	// Bind JSON body
	var body struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Price		 float64 `json:"price"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find existing court
	var court models.Court
	if err := initializers.DB.First(&court, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Court not found"})
		return
	}

	// Update fields
	court.Name = body.Name
	court.Location = body.Location
	court.Price = body.Price

	if err := initializers.DB.Save(&court).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Court updated successfully",
		"data":    court,
	})
}

func DeleteCourt(c *gin.Context) {
	id := c.Param("id")

	// Check if court exists first
	var court models.Court
	if err := initializers.DB.First(&court, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Court not found"})
		return
	}

	// Delete the record
	if err := initializers.DB.Delete(&court).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Court deleted successfully",
	})
}
