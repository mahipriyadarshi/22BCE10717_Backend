package main

import (
	"log"
	"net/http"
	"github.com/joho/godotenv"


	"file-sharing-app/config"
	"file-sharing-app/handlers"
	"file-sharing-app/middleware"
)

func main() {
	// Initialize configurations
	_ = godotenv.Load()
	config.InitDB()
	config.InitRedis()
	config.InitAWS()

	// Register routes
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.Handle("/upload", middleware.JWTMiddleware(http.HandlerFunc(handlers.UploadHandler)))
	http.Handle("/files", middleware.JWTMiddleware(http.HandlerFunc(handlers.ListFilesHandler)))
	http.Handle("/share/", middleware.JWTMiddleware(http.HandlerFunc(handlers.ShareHandler)))

	log.Println("Server running on :5050")
	log.Fatal(http.ListenAndServe(":5050", nil))
}
