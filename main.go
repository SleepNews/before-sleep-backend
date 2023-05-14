package main

import (
	"github.com/finalfree/before-sleep-backend/api"
	"github.com/finalfree/before-sleep-backend/repo"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func setupRouter(server *api.Server) *gin.Engine {
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

func setUpDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/before_sleep?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&repo.User{}, &repo.Topic{}, &repo.Comment{}, &repo.CommentLike{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	//setUpTestData(db)
	return db
}

func setUpTestData(db *gorm.DB) {
	user := repo.User{Name: "admin"}
	visitor := repo.User{Name: "visitor"}
	db.Create(&user)
	db.Create(&visitor)
	topic := repo.Topic{Title: "test topic", Content: "test content", UserID: user.ID}
	db.Create(&topic)
	db.Create(&repo.Comment{Content: "test comment", TopicID: topic.ID, UserID: visitor.ID})
}

func main() {
	server := &api.Server{
		DB: setUpDB(),
	}
	r := setupRouter(server)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":80")
}
