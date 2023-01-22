package daos

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string
	Description string
	Author      string
}
