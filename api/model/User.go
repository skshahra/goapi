package model

import "github.com/golang-jwt/jwt/v4"

type UserData struct {
	ID 	  uint   `json:"id" gorm:"primary_key"`
	Firstname string `json:"firstname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Lastname  string `json:"lastname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Email     string `json:"email" gorm:"unique;not null;type:varchar(100)" validate:"required,min=3,max=100"`
	Password  string `json:"password" gorm:"type:varchar(100)" validate:"required,min=3,max=50"`
	GoogleAuth bool `json:"GoogleAuth" `
	Telephone string `json:"telephone" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Address string `json:"address" gorm:"type:varchar(100)" validate:"min=2,max=100"`
}



type UserLogin struct {
	Email     string `json:"email" gorm:"unique;not null;type:varchar(100)" validate:"required,min=3,max=100"`
	Password  string `json:"password" gorm:"type:varchar(100);not null" validate:"required,min=3,max=50"`
}


type UserSignup struct {
	Firstname string `json:"firstname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Lastname  string `json:"lastname" gorm:"type:varchar(100)" validate:"min=2,max=100"`
	Email     string `json:"email" gorm:"unique;not null;type:varchar(100)" validate:"required,min=3,max=100"`
	Password  string `json:"password" gorm:"type:varchar(100);not null" validate:"required,min=3,max=50"`
}



type Claims struct {
	UserName string `json:"UserName"`
	Email string `json:"email"`
	GoogleAuth bool `json:"GoogleAuth"`
	jwt.StandardClaims	
}