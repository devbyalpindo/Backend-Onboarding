package dto

type UserRequest struct {
	PersonalNumber string `validate:"required" json:"personal_number"`
	Password       string `validate:"required" json:"password"`
	Email          string `validate:"required,email" json:"email"`
	Name           string `validate:"required" json:"name"`
	RoleID         string `json:"role_id"`
}
