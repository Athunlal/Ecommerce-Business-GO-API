package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname   string `JSON:"name" validate:"required,min=3,max=12"`
	Lastname    string
	Email       string `JSON:"email"`
	Password    string `validate:"requierd, gte=6"`
	PhoneNumber int
	Otp         string
}

type Admin struct {
	gorm.Model
	Firstname   string
	Lastname    string
	Email       string
	Password    string
	PhoneNumber int
}
