package main

import (
	config "entdemo/Config"
	router "entdemo/Router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

	defer config.Client.Close()

	app := gin.Default()
	app.GET("", func(c *gin.Context) {
		c.File("public/index.html")
	})
	router.RegisterRouter(app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8090"
	}

	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)

	}
}
