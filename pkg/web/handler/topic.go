package handler

import (
	"strconv"

	"github.com/alonegrowing/insight/pkg/basic/macro"

	"github.com/alonegrowing/insight/pkg/web/controller"
	"github.com/alonegrowing/insight/pkg/web/controller/impl"
	"github.com/gin-gonic/gin"
)

type PoemHandler struct {
	BaseHandler
	homeController controller.HomeController
}

func NewHomePageHandler() *PoemHandler {
	return &PoemHandler{
		homeController: impl.DefaultHomeController,
	}
}

func (d *PoemHandler) Get(r *gin.Context) {
	id, err := strconv.ParseInt(r.Request.FormValue("id"), 10, 64)
	if id <= 0 || err != nil {
		d.error(r, macro.ErrorParamIllegal)
	}
	d.success(r, macro.TopicData[0])
}

func (d *PoemHandler) GetList(r *gin.Context) {
	d.success(r, macro.TopicData)
}

func (d *PoemHandler) GetCommentList (r *gin.Context) {
	d.success(r, macro.Comments)
}
