package repo

import "gorm.io/gorm"

type CommentLike struct {
	gorm.Model
	UserID    uint `json:"user_id" gorm:"foreignKey:UserID;index:user_id_comment_id"`
	CommentID uint `gorm:"foreignKey:CommentID;index:user_id_comment_id"`
	Like      bool `gorm:"not null"`
}
