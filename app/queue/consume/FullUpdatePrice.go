package consume

import (
	"advt/app/queue"
	"advt/internal/rabbitMQ"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

type FullUpdatePrice struct {
}

func (r *FullUpdatePrice) GetRoutingKey() string {
	return queue.TestQueueRoutingKey
}

func (r *FullUpdatePrice) GetQueueName() string {
	return queue.TestQueueName
}
func (r *FullUpdatePrice) GetExchangeName() string {
	return queue.DefaultExchangeName
}

func (r *FullUpdatePrice) GetOtherDeclareInfo() *rabbitMQ.ExchangeBuild {
	return &rabbitMQ.ExchangeBuild{
		ExchangeName: r.GetExchangeName(),
		ExchangeType: amqp.ExchangeDirect,
		Durable:      true,
		Arguments: map[string]interface{}{
			"x-max-priority": 20,
		},
	}
}

func (r *FullUpdatePrice) Process(data string) error {
	facade.GetRedisProvider().Get().Do("set", "Text", data)
	fmt.Println(data)
	return errors.New("FullUpdatePrice测试错误" + r.GetQueueName())
}
func (r *FullUpdatePrice) FailsAfter(data string, err error) {

}
