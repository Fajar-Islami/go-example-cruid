package daos

import "gorm.io/gorm"

type (
	Book struct {
		gorm.Model
		Title       string
		Description string
		Author      string
	}

	FilterBooks struct {
		Limit, Offset int
		Title         string
	}
)
