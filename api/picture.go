package api

import (
	"net/http"
	"web/models"
	"web/repository"

	"github.com/gin-gonic/gin"
)

// Get all pictures
func GetAllPictures(c *gin.Context) {
	var pictures []models.Picture
	repository.DB.Find(&pictures)

	c.JSON(http.StatusOK, gin.H{"pictures": pictures})
}

// Get picture by ID
func GetPicture(c *gin.Context) {
	var picture models.Picture
	if err := repository.DB.First(&picture, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Picture doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"picture": picture})
}

// Create new records
func CreatePicture(c *gin.Context) {
	var newPicture models.CreatePicture

	if err := c.ShouldBindJSON(&newPicture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	picture := models.Picture{
		Name:     newPicture.Name,
		Author:   newPicture.Author,
		Year:     newPicture.Year,
		Material: newPicture.Material,
	}

	repository.DB.Create(&picture)

	c.JSON(http.StatusOK, gin.H{"picture": picture})
}

// Update picture by ID
func UpdatePicture(c *gin.Context) {
	var picture models.Picture

	if err := repository.DB.First(&picture, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Picture doesn't exist"})
		return
	}

	var receivedData models.UpdatePicture

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPicture := models.Picture{
		Name:     receivedData.Name,
		Author:   receivedData.Author,
		Year:     receivedData.Year,
		Material: receivedData.Material,
	}

	repository.DB.Model(&picture).Updates(updatedPicture)

	c.JSON(http.StatusOK, gin.H{"picture": picture})

}

// Delete picture record by ID
func DeletePicture(c *gin.Context) {
	var picture models.Picture
	if err := repository.DB.First(&picture, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Picture doesn't exist"})
		return
	}

	repository.DB.Delete(&picture)
	c.JSON(http.StatusNoContent, gin.H{"result": "deleted"})
}
