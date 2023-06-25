package http

import (
	"advt/internal/file"
)

type HttpServerInterface interface {
	AddJsonRpcServer(path string, server interface{})
	Run()
}

type HttpServerProvider struct {
	configReader file.ConfigReader
}

type RpcProvider struct {
	configReader file.ConfigReader
}
