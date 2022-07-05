package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id          string  `gorm:"primaryKey;not_null;type:varchar(110)"`
	Name        string  `gorm:"type:varchar(100)"`
	Description string  `gorm:"type:text"`
	Status      string  `gorm:"type:enum('inactive', 'approved', 'active');default:'inactive'"`
	MakerID     string  `gorm:"column:maker_id;type:varchar(110)"`
	Maker       User    `gorm:"foreignKey:MakerID;references:Id"`
	CheckerID   *string `gorm:"column:checked_id;type:varchar(110);default:null"`
	Checker     User    `gorm:"foreignKey:CheckerID;references:Id"`
	SignerID    *string `gorm:"column:signer_id;type:varchar(110):default:null"`
	Signer      User    `gorm:"foreignKey:SignerID;references:Id"`
}
