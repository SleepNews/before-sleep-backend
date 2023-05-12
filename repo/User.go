package repo

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"unique;size:16"`
}
