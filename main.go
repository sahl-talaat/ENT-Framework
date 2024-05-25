package main

import (
	config "entdemo/Config"
	router "entdemo/Router"
	"log"
	"net/http"
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
	router.RegisterRouter(app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8090"
	}

	log.Printf("server listening on port %s", port)
	if err := http.ListenAndServe(":"+port, app); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
