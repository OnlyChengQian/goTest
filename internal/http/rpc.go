package http

import (
	"advt/app/Router"
	"advt/internal/file"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//可以直接用http请求

func NewRpcProvider(configReader file.ConfigReader) HttpServerInterface {
	return &RpcProvider{
		configReader: configReader,
	}
}

func (j *RpcProvider) AddJsonRpcServer(path string, server interface{}) {
	err := rpc.Register(server)
	if err != nil {
		log.Fatal(err)
	}
	isFunc := func(writer http.ResponseWriter, request *http.Request) {
		var conn = j.conn(writer, request)
		err = j.rpcServerRequest(conn)
		if err != nil {
			writer.WriteHeader(410)
		}
	}
	http.HandleFunc(path, isFunc)
}

// 这里可以使用http请求
func (j *RpcProvider) conn(w http.ResponseWriter, r *http.Request) io.ReadWriteCloser {
	return struct {
		io.Writer
		io.ReadCloser
	}{
		ReadCloser: r.Body,
		Writer:     w,
	}
}

func (j *RpcProvider) rpcServerRequest(conn io.ReadWriteCloser) error {
	return rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
}

func (j *RpcProvider) Handle(addr string) {
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err, "server listen err %s")
	}
}

func (j *RpcProvider) Run() {
	config := j.configReader.GetServerConfig()
	if config.RpcPort != "" {
		rpcServer := Router.NewRouterProvider(nil)
		rpcServer.AddServer()
		address := ":" + config.RpcPort
		j.Handle(address)
		log.Print("SERVER SUCCESS LISTEN " + address)
	}
}
