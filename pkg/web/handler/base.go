package handler

import (
	"github.com/alonegrowing/insight/pkg/basic/macro"
	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	gin.HandlersChain
}

func (m *BaseHandler) success(r *gin.Context, data interface{}) {
	r.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}

func (m *BaseHandler) error(r *gin.Context, code int64) {
	r.JSON(200, gin.H{
		"code":    code,
		"message": macro.ErrMsg[code],
		"data":    []interface{}{},
	})
}
