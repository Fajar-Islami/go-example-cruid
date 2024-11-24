package entity

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Email    string
		Name     string
		Password string
		Books    []Book
	}

	FilterUser struct {
		Limit, Offset int
		Title         string
	}
)
