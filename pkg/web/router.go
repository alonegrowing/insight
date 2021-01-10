package web

import (
	"io/ioutil"

	"github.com/alonegrowing/insight/pkg/web/handler"
	"github.com/alonegrowing/insight/pkg/web/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	route := gin.Default()
	route.Use(middleware.Logger())

	route.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	homeHandler := handler.NewHomePageHandler()

	route.GET("/api/topic", homeHandler.Get)
	route.GET("/api/topics", homeHandler.GetList)
	route.GET("/api/comments", homeHandler.GetCommentList)
	return route
}
