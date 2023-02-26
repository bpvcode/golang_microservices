package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	email string `gorm:"unique"`
}
