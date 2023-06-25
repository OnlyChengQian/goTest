package controller

import (
	"advt/app/queue/produce"
	"fmt"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (IndexController) Index(ctx *gin.Context) {
	producer := &produce.Producer{}
	content := "6666666666"
	err := producer.Push(new(produce.ComputeShippingProducer), content)
	if err != nil {
		ctx.JSON(202, gin.H{
			"method":  "Get",
			"message": "hello world",
			"error":   err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"method":  "Get",
			"message": "success",
		})
	}
	//tableServer := &mongo.Mongo{}
	//data, err := tableServer.Get(new(mongo.Site), bson.D{{"site_code", "DE"}})
	//context.WithValue(ctx, "account", data)
	//Account := ctx.Value("account")
	//fmt.Println(Account)
	a, _ := facade.GetRedisProvider().Get().Do("get", "aaa")

	fmt.Println(a)
}
