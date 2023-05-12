package api

import "gorm.io/gorm"

type Server struct {
	DB *gorm.DB
}
