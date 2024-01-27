package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Company struct {
    ID   string `json:"id" bson:"_id,omitempty"`
    Name string `json:"name" bson:"name"`
}

func main() {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(context.TODO())

    // Create a router
    r := mux.NewRouter()
    r.HandleFunc("/company/{name}", getCompanyByName(client)).Methods("GET")

    // Start the server
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func getCompanyByName(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        name := vars["name"]

        var company Company
        collection := client.Database("Master").Collection("Companies")
        err := collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&company)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }

        // Return the result
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(company)
    }
}
