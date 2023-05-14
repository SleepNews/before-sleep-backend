package repo

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content    string
	LikeNum    uint `gorm:"default:0"`
	DislikeNum uint `gorm:"default:0"`
	UserID     uint `json:"user_id"`
	User       User
	TopicID    uint `gorm:"foreignKey:TopicID"`
	ParentID   uint `gorm:"default:0"`
}
