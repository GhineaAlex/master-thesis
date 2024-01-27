// /internal/models/company.go

package models

// Company represents a company entity in the database.
type Company struct {
    Name            string `json:"name" bson:"name"`
    Email           string `json:"email" bson:"email"`
    DateOfFoundation string `json:"date_of_foundation" bson:"date_of_foundation"`
    TotalMoneyRaised string `json:"total_money_raised" bson:"total_money_raised"`
}