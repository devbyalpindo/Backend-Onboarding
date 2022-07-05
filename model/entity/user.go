package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id             string `gorm:"primaryKey;not_null;type:varchar(110)"`
	PersonalNumber string `gorm:"column:personal_number;type:varchar(200)"`
	Password       string `gorm:"type:varchar(100)"`
	Email          string `gorm:"type:varchar(100)"`
	Name           string `gorm:"type:varchar(100)"`
	RoleID         string `gorm:"column:role_id;type:varchar(110)"`
	Role           Role   `gorm:"foreignKey:RoleID;references:Id"`
	Active         bool   `gorm:"type:boolean;default:true"`
}
