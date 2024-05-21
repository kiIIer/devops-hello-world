package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialize Viper to read environment variables
	viper.AutomaticEnv()
	viper.SetDefault("WORLD", "world")

	world := viper.GetString("WORLD")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf("Hi, I am %s", world)
		fmt.Fprintln(w, response)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
