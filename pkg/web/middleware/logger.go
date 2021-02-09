package middleware

import (
	"strings"
	"time"

	"github.com/alonegrowing/insight/pkg/treasure/log"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()
		handlerNamePathParts := strings.Split(c.HandlerName(), "/")
		handleName := handlerNamePathParts[len(handlerNamePathParts)-1]
		log.Infof("%20v | %3d | %13v | %15s | %-7s %s\n%s", handleName, statusCode, latency, clientIP, method, path, comment)
	}
}
