package routers

import (
	"crawl_html_from_dc/actions"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.Default()
}

func Load() *gin.Engine {
	r.POST("dc-send", actions.DcSend)
	r.POST("dc-receive", actions.DcReceive)

	return r
}
