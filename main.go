package main

import (
	"log"
	config "myapp/Config"
	routes "myapp/Router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = config.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	app := gin.Default()
	app.GET("", func(c *gin.Context) {
		c.File("public/index.html")
	})
	routes.RegisterRouter(app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8090"
	}

	err = app.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)

	}
}
