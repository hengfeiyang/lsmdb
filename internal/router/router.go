package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hengfeiyang/lsmdb/internal/service"
)

func Route(app *gin.Engine) {
	app.POST("/api/v1/set", service.Set)
	app.GET("/api/v1/get/:key", service.Get)
	app.GET("/api/v1/list", service.List)
}
