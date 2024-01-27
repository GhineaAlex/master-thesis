// /internal/handlers/company.go

package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "internal/models"
    "pkg/db"
    "go.mongodb.org/mongo-driver/bson"
)

// CreateCompany handles the POST request to add a new company.
func CreateCompany(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var company models.Company
    if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := db.CompaniesCollection.InsertOne(context.TODO(), company)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(company)
}

// GetCompany handles the GET request to find a company by name.
func GetCompany(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var companyName = r.URL.Query().Get("name")

    var company models.Company
    if err := db.CompaniesCollection.FindOne(context.TODO(), bson.M{"name": companyName}).Decode(&company); err != nil {
        http.Error(w, "Company not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(company)
}
