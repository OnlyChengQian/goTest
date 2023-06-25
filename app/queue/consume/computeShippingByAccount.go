package consume

import (
	"advt/app/api/model/mongo"
	"advt/app/queue"
	"advt/internal/rabbitMQ"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
)

type ComputerShippingByAccount struct {
}

func (r *ComputerShippingByAccount) GetRoutingKey() string {
	return queue.ComputeShippingRoutingKey
}

func (r *ComputerShippingByAccount) GetQueueName() string {
	return queue.ComputeShippingQueueName
}
func (r *ComputerShippingByAccount) GetExchangeName() string {
	return queue.DefaultExchangeName
}

func (r *ComputerShippingByAccount) GetOtherDeclareInfo() *rabbitMQ.ExchangeBuild {
	return &rabbitMQ.ExchangeBuild{
		ExchangeName: r.GetExchangeName(),
		ExchangeType: amqp.ExchangeDirect,
		Durable:      true,
		Arguments: map[string]interface{}{
			"x-max-priority": 20,
		},
	}
}

// Generate a concatenated query mysql to obtain product information method

func (r *ComputerShippingByAccount) Process(data string) error {
	common := &mongo.Mongo{}
	account, err := common.First(new(mongo.Account), bson.M{"account": data})
	if err != nil {
		return errors.New("获取数据错误" + err.Error())
	}
	fmt.Println(account)
	return errors.New("ComputerShippingByAccount测试错误" + r.GetQueueName())
}

func (r *ComputerShippingByAccount) FailsAfter(data string, err error) {

}
