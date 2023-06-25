package app

import (
	"advt/app/queue/consume"
	"advt/internal/file"
	http2 "advt/internal/http"
)

func Run() {
	configReader := file.GetConfig()
	//先启动gin-http，再启动消费队列
	go func() {
		http := http2.NewHttpProvider(configReader)
		http.Run()
	}()
	//启动rabbit消费者监听
	consume.Run()
	//有监听消费者的时候，这个select不会执行
	select {}
}
