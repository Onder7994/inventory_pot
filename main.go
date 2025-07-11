package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"top_ten/internal/handlers"
	"top_ten/internal/middleware"
	"top_ten/internal/utils"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	utils.InitDB()
	defer utils.CloseDB()

	middleware.InitLogger()
	defer middleware.CloseLogger()

	app_port := os.Getenv("APP_PORT")
	if app_port == "" {
		fmt.Println("APP_PORT variable not set, using default port 8080")
		app_port = "8080"
	}
	addr := fmt.Sprintf(":%s", app_port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/admin", handlers.AdminHandler)
	mux.HandleFunc("/debug", handlers.DebugHandler)

	loggedMux := middleware.LoggingMiddleware(mux)

	fmt.Printf("Honeypot top_ten server running at %s port\n", app_port)
	log.Fatal(http.ListenAndServe(addr, loggedMux))
}
