package Router

import (
	"advt/app/api/controller"
	"github.com/gin-gonic/gin"
)

// Router http 路由
func (r RouteProvider) Router() *gin.Engine {
	var indexController = new(controller.IndexController)
	r.http.GET("/test", indexController.Index)
	return r.http
}
