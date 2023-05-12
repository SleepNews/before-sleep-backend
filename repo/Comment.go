package repo

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	Like    uint
	Dislike uint
	UserID  uint
	User    User
	TopicID uint
	Topic   Topic
}
