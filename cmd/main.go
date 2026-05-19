package main

import (
	"data_hoarder_go/config"
	"data_hoarder_go/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	config.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	port := os.Getenv("PORT")

	fmt.Println(
		"Server running on port",
		port,
	)

	router.Run(
		"0.0.0.0:" + port,
	)
}