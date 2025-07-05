package main

import (
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/reftch/go-react-template/configs"
	"github.com/reftch/go-react-template/internal/controllers"
)

func main() {
	startTime := time.Now() // Capture the start time

	// Serve files from the "static" directory
	fileServer := http.FileServer(http.Dir("./web/static"))
	// Strip the "/static/" prefix when looking for files
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	c := controllers.New()
	c.GET("/", c.HomeHandler)

	if configs.Envs.Environment == "development" {
		c.GET("/ws", c.WsHandler)
	}

	// Start server with some basic configurations
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// measure startup time
	r := new(big.Int)
	r.Binomial(1000, 10)
	elapsed := time.Since(startTime)
	log.Printf(" Server started on port 8080 in %s", elapsed)

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error starting server: %s\n", err)
	}
}
