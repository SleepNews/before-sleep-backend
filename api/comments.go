package api

import (
	"github.com/finalfree/before-sleep-backend/repo"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *Server) CommentDetailHandler(c *gin.Context) {
	idStr := c.Params.ByName("id")
	// parse id to uint
	id, _ := strconv.Atoi(idStr)
	var result repo.Comment
	s.DB.Preload("User").First(&result, id)
	c.JSON(200, result)
}
