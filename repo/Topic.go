package repo

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint
	User    User
}
