package main

import (
	config "entdemo/Config"
	middleware "entdemo/Middleware"
	router "entdemo/Router"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const defaultPort = "8090"

var err error

func main() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	// initiate Ent Client
	client, err := config.NewEntClient()
	if err != nil {
		log.Printf("err : %s", err.Error())
	}

	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}

	config.SetClient(client)

	r := mux.NewRouter()
	r.Use(middleware.Header)
	router.RegisterRouter(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server start on port " + port)
	log.Fatal(srv.ListenAndServe())
}
