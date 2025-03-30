package main

import (
	"log"
	"net/http"
	"github.com/joho/godotenv"


	"app/config"
)

func main() {
	// Initialize configurations
	_ = godotenv.Load()
	config.InitDB()
	config.InitRedis()
	config.InitAWS()

	// Register routes
	http.HandleFunc("/register", config.RegisterHandler)
	http.HandleFunc("/login", config.LoginHandler)
	http.Handle("/upload", config.JWTMiddleware(http.HandlerFunc(handlers.UploadHandler)))
	http.Handle("/files", config.JWTMiddleware(http.HandlerFunc(handlers.ListFilesHandler)))
	http.Handle("/share/", config.JWTMiddleware(http.HandlerFunc(handlers.ShareHandler)))

	log.Println("Server running on :5050")
	log.Fatal(http.ListenAndServe(":5050", nil))
}
