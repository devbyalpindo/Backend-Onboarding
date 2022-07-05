package dto

type User struct {
	Id             string `json:"id"`
	PersonalNumber string `json:"personal_number"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Role           Role   `json:"role"`
	Active         bool   `json:"active"`
}
