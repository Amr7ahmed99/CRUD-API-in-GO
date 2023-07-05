package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model // for createdAt and updatedAt
	Title      string
	Body       string
}
