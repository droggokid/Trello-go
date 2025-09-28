package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rudi/Trello-go/internal/models"
	"github.com/rudi/Trello-go/pkg/database"
)

func BoardCreate(c *gin.Context) {
	var body struct {
		Name string
	}

	c.Bind(&body)

	board := models.Board{Name: body.Name}

	result := database.DB.Create(&board)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"board": board,
	})
}

func BoardIndex(c *gin.Context) {
	var boards []models.Board
	database.DB.Find(&boards)

	c.JSON(200, gin.H{
		"board": boards,
	})
}

func BoardShow(c *gin.Context) {
	id := c.Param("id")

	var board models.Board
	database.DB.First(&board, id)

	c.JSON(200, gin.H{
		"board": board,
	})
}

func BoardUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name string
	}

	c.Bind(&body)

	var board models.Board
	database.DB.First(&board, id)

	database.DB.Model(&board).Updates(models.Board{
		Name: body.Name,
	})

	c.JSON(200, gin.H{
		"board": board,
	})
}

func BoardDelete(c *gin.Context) {
	id := c.Param("id")

	database.DB.Delete(&models.Board{}, id)

	c.Status(200)
}
