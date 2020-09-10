package router

import (
	"SenseHoney/app/api"
	"github.com/gin-gonic/gin"
)

type Service struct {
	api.Service
}

func InitRouter(s *Service) *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(s.SuccessJSON("SenseHoney"))

	})
	r.POST("/api/report", func(c *gin.Context) {
		c.JSON(s.ReportHandler(c))
	})

	r.POST("/api/log", func(c *gin.Context) {
		c.JSON(s.LogsHandler(c))
	})

	r.GET("/ws", func(c *gin.Context) {
		s.WsHandler(c)
	})

	r.GET("/dataInit", func(c *gin.Context) {
		c.String(200, string(s.DataInfo()))
	})
	return r

}
