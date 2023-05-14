package api

import (
	"fmt"
	"github.com/finalfree/before-sleep-backend/repo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (server *Server) GetTopicComments(c *gin.Context) {
	idStr := c.Params.ByName("topic_id")
	// parse id to uint
	id, _ := strconv.Atoi(idStr)
	var comments []repo.Comment
	server.DB.Where("topic_id = ?", id).Find(&comments)
	c.JSON(200, comments)
}

func (server *Server) PostComment(c *gin.Context) {
	var comment repo.Comment
	c.BindJSON(&comment)
	idStr := c.Params.ByName("topic_id")
	// parse id to uint
	id, _ := strconv.Atoi(idStr)
	comment.TopicID = uint(id)
	result := server.DB.Select("Content", "UserID", "TopicID", "ParentID").Create(&comment)
	if result.Error != nil {
		c.JSON(500, result.Error)
		return
	}
	c.JSON(200, comment)
}

func (server *Server) LikeComment(c *gin.Context) {
	server.LikeOrDislikeComment(c, true)
}

func (server *Server) DislikeComment(c *gin.Context) {
	server.LikeOrDislikeComment(c, false)
}

func (server *Server) LikeOrDislikeComment(c *gin.Context, like bool) {
	idStr := c.Query("user_id")
	// parse id to uint
	userID, _ := strconv.Atoi(idStr)
	idStr = c.Params.ByName("comment_id")
	// parse id to uint
	commentID, _ := strconv.Atoi(idStr)

	var commentLike repo.CommentLike
	result := server.DB.Where("user_id = ? AND comment_id = ?", userID, commentID).First(&commentLike)

	updateValues := map[string]interface{}{}

	if c.Request.Method == "DELETE" {
		if result.RowsAffected == 0 || like != commentLike.Like {
			fmt.Println("SQL: ", result.Statement.SQL.String())
			c.JSON(200, "already canceled "+getOperation(like))
			return
		}
		result = server.DB.Delete(&commentLike)
		if result.Error != nil {
			fmt.Println("SQL: ", result.Statement.SQL.String())
			c.JSON(500, result.Error)
			return
		}
		updateValues[getOperation(like)] = gorm.Expr(getOperation(like)+" - ?", 1)
	} else {
		if result.RowsAffected == 0 {
			commentLike = repo.CommentLike{UserID: uint(userID), CommentID: uint(commentID), Like: like}
			result = server.DB.Create(&commentLike)
			if result.Error != nil {
				fmt.Println("SQL: ", result.Statement.SQL.String())
				c.JSON(500, result.Error)
				return
			}
			updateValues[getOperation(like)] = gorm.Expr(getOperation(like)+" + ?", 1)
		} else {
			if like == commentLike.Like {
				fmt.Println("SQL: ", result.Statement.SQL.String())
				c.JSON(200, "already "+getOperation(like))
				return
			}
			commentLike.Like = like
			server.DB.Save(&commentLike)
			updateValues[getOperation(like)] = gorm.Expr(getOperation(like)+" + ?", 1)
			updateValues[getOperation(!like)] = gorm.Expr(getOperation(!like)+" - ?", 1)
		}
	}

	result = server.DB.Model(&repo.Comment{}).Where("id = ?", commentID).Updates(updateValues)
	if result.Error != nil {
		fmt.Println("SQL: ", result.Statement.SQL.String())
		c.JSON(500, result.Error)
		return
	}
	c.String(200, getOperation(like)+" success")
}

func getOperation(like bool) string {
	if like {
		return "like_num"
	}
	return "dislike_num"
}
