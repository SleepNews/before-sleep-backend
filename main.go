package main

import (
	"fmt"
	"github.com/finalfree/sleeping-news/api"
	"github.com/finalfree/sleeping-news/repo"
	"github.com/jessevdk/go-flags"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Options struct {
	GenerateTestData bool `short:"g" long:"generate-test-data" description:"generate test data"`
}

func setUpDB(options *Options) *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/sleeping_news?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&repo.User{}, &repo.Topic{}, &repo.Comment{}, &repo.CommentLike{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	if options.GenerateTestData {
		setUpTestData(db)
	}
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
	var options Options
	_, err := flags.ParseArgs(&options, os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
	server := &api.Server{
		DB: setUpDB(&options),
	}
	r := server.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":80")
}
