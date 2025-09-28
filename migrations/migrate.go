package main

import (
	"github.com/rudi/Trello-go/internal/models"
	"github.com/rudi/Trello-go/pkg/config"
	"github.com/rudi/Trello-go/pkg/database"
)

func init() {
	config.LoadEnvVars()
	database.ConnectDb()
}

func main() {
	database.DB.AutoMigrate(&models.Board{})
}
