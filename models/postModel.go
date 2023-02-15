package models

import "gorm.io/gorm"

type Post struct {
	// gorm.Model defines a boilerplate struct including ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Title   string
	Content string
}
