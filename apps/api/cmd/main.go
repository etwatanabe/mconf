package main

import (
    "log"
    "net/http"

    "github.com/urfave/negroni/v3"
    "github.com/gorilla/mux"

    "api/internal/config"
    "api/internal/handlers"
    "api/internal/services"
)

func main() {
    // Load configuration
    cfg := config.Load()

    // Services initialization
    openLibraryService := services.NewOpenLibraryService(cfg.OpenLibraryURL, cfg.RequestLimit)

    // Handlers initialization
    bookHandler := handlers.NewBookHandler(openLibraryService)

    // Routes configuration
    r := mux.NewRouter()

    // Health check route
    r.HandleFunc("/health", bookHandler.HealthCheck).Methods("GET")

    // API routes
    api := r.PathPrefix("/api/v1").Subrouter()
    api.HandleFunc("/search", bookHandler.SearchBooks).Methods("GET")

    n := negroni.Classic()
    
    n.UseHandler(r)

    log.Printf("Starting server on port %s", cfg.Port)
    log.Printf("Health check available at: http://localhost:%s/health", cfg.Port)
    log.Printf("Search endpoint: http://localhost:%s/api/v1/search?q=<query>", cfg.Port)

    log.Fatal(http.ListenAndServe(":"+cfg.Port, n))
}
