package http

import (
	"advt/app/Router"
	"advt/internal/file"
	"github.com/gin-gonic/gin"
	"log"
)

func NewHttpProvider(configReader file.ConfigReader) HttpServerInterface {
	return &HttpServerProvider{configReader: configReader}
}

func (h *HttpServerProvider) AddJsonRpcServer(path string, server interface{}) {
}

func (h *HttpServerProvider) Run() {
	serverConfig := h.configReader.GetServerConfig()
	if serverConfig.Port != "" {
		if serverConfig.Debug == "true" {
			gin.SetMode(gin.DebugMode)
		}
		addr := ":" + serverConfig.Port
		http := gin.Default()
		custom := Router.NewRouterProvider(http)
		log.Print("http server listen " + addr)
		err := custom.Router().Run(addr)
		if err != nil {
			log.Fatal("站点开启异常" + err.Error())
		}
	}
}
