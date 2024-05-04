package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name      string
	UserName  string
	UserEmail string
}
