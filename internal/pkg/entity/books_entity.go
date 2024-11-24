package entity

import "gorm.io/gorm"

type (
	Book struct {
		gorm.Model
		UserID      uint
		User        User
		Title       string
		Description string
		Author      string
	}

	FilterBooks struct {
		Limit, Offset int
		Title         string
	}
)
