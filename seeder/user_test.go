package seeder

import (
	"fse-onboarding/helper"
	"fse-onboarding/model/entity"
	"testing"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSeedUserAdmin(t *testing.T) {

	dsn := "root:root@tcp(localhost:3306)/onboarding?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		t.Fatal("failed to connect database")
	}

	encryptPwd, err := helper.HashPassword("admin")

	if err != nil {
		t.Fatal("err encrypt password")
	}

	createID := uuid.New().String()
	user := entity.User{
		Id:             createID,
		PersonalNumber: "01",
		Password:       encryptPwd,
		Email:          "sharingvision@gmail.com",
		Name:           "sv",
		RoleID:         "2e49bfa7-7898-43ee-8627-a00745c0da5a",
	}

	errDB := DB.Create(&user).Error
	if errDB != nil {
		t.Fatal("error batch insert to mysql", err.Error())
	}
}

func TestSeedUserChecker(t *testing.T) {

	dsn := "root:root@tcp(localhost:3306)/onboarding?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		t.Fatal("failed to connect database")
	}

	encryptPwd, err := helper.HashPassword("checker")

	if err != nil {
		t.Fatal("err encrypt password")
	}

	createID := uuid.New().String()
	user := entity.User{
		Id:             createID,
		PersonalNumber: "02",
		Password:       encryptPwd,
		Email:          "sharingvision@gmail.com",
		Name:           "sv",
		RoleID:         "29a134e4-22f9-4464-961f-b3240b43b0f7",
	}

	errDB := DB.Create(&user).Error
	if errDB != nil {
		t.Fatal("error batch insert to mysql", err.Error())
	}
}
