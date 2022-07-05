package seeder

import (
	"fse-onboarding/model/entity"
	"testing"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSeedRole(t *testing.T) {

	dsn := "root:root@tcp(localhost:3306)/onboarding?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		t.Fatal("failed to connect database")
	}

	createID := uuid.New().String()
	role := []entity.Role{
		{
			Id:     createID,
			Title:  "signer",
			Active: true,
		},
	}

	errDB := DB.Create(&role).Error
	if errDB != nil {
		t.Fatal("error batch insert to mysql", errDB)
	}
}
