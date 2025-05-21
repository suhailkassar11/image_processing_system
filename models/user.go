package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null" validate:"required"`
	Email    string `json:"email" gorm:"not null" validate:"required"`
	Password string `json:"password" gorm:"not null" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"not null" validate:"required"`
	Password string `json:"password" gorm:"not null" validate:"required"`
}
