package main

import (
    "log"
    "net/http"
    "/internal/handlers"
    "github.com/gorilla/mux"
    "/pkg/db"
)

func main() {
    // Initialize database connection (optional if you have a separate setup for this)
    if err := db.Init(); err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Create a new router
    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/api/add_company", handlers.CreateCompany).Methods("POST")
    r.HandleFunc("/api/company", handlers.GetCompany).Methods("GET")

    // Start the server
    log.Println("Starting server on port 8000")
    if err := http.ListenAndServe(":8000", r); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
