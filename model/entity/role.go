package entity

type Role struct {
	Id     string `gorm:"primaryKey;not_null;type:varchar(110)"`
	Title  string `gorm:"type:enum('admin', 'maker', 'checker', 'signer', 'viewer');default:'viewer'"`
	Active bool   `gorm:"type:boolean;default:true"`
}
