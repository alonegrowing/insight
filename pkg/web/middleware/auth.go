package middleware

import (
	"github.com/alonegrowing/insight/pkg/basic/macro"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(r *gin.Context) {
		id := r.Query("id")
		if id == "2" {
			r.JSON(200, gin.H{
				"code":    macro.ErrorAuthFailed,
				"message": macro.ErrMsg[macro.ErrorAuthFailed],
				"data":    []int{},
			})
			r.Abort()
		}
		r.Next()
	}
}
