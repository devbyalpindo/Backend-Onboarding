package dto

type UserLogin struct {
	PersonalNumber string `validate:"required" json:"personal_number"`
	Password       string `validate:"required,min=8" json:"password"`
}
