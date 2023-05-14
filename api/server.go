package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	DB *gorm.DB
}

func (server *Server) SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/topics/:topic_id/comments", server.GetTopicComments)
	r.POST("/topics/:topic_id/comments", server.PostComment)
	r.Match([]string{"POST", "DELETE"}, "/comments/:comment_id/like", server.LikeComment)
	r.Match([]string{"POST", "DELETE"}, "/comments/:comment_id/dislike", server.DislikeComment)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// return User ip
	r.GET("/ip", func(c *gin.Context) {
		c.String(http.StatusOK, c.Request.RemoteAddr)
	})

	return r
}
