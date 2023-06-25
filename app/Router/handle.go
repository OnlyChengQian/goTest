package Router

import (
	"github.com/gin-gonic/gin"
	"log"
)

type RouteInterface interface {
	AddServer()
	Router() *gin.Engine
}

type RouteProvider struct {
	http *gin.Engine
}

func NewRouterProvider(http interface{}) RouteInterface {
	if http != nil {
		v, ok := http.(*gin.Engine)
		if !ok {
			log.Fatalln("http 类型非gin")
		}
		//http
		return &RouteProvider{http: v}
	}
	//rpc
	return RouteProvider{http: nil}
}
