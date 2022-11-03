package main

import (
	"net/http"
	"web/api"
	"web/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDB()
	r := gin.Default()

	// test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Routees
	r.GET("/picture", api.GetAllPictures)
	r.GET("/picture/:id", api.GetPicture)
	r.POST("/picture", api.CreatePicture)
	r.PATCH("/picture/:id", api.UpdatePicture)
	r.DELETE("/picture/:id", api.DeletePicture)

	r.Run() // listen and serve on 0.0.0.0:8080
}
