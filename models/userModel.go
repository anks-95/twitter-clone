package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsVerified bool   `json:"is_verified"`
	Token      string `json:"-"`
}
