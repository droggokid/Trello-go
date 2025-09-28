package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudi/Trello-go/internal/handlers"
	"github.com/rudi/Trello-go/pkg/config"
	"github.com/rudi/Trello-go/pkg/database"
)

func init() {
	config.LoadEnvVars()
	database.ConnectDb()
}

func main() {
	r := gin.Default()

	r.GET("/board", handlers.BoardIndex)
	r.GET("/board/:id", handlers.BoardShow)
	r.POST("/board", handlers.BoardCreate)
	r.PUT("/board/:id", handlers.BoardUpdate)
	r.DELETE("/board/:id", handlers.BoardDelete)

	r.Run()
}
