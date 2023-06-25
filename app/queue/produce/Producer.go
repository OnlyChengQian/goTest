package produce

import (
	"advt/app/queue"
	facade2 "advt/internal/facade"
	"advt/internal/rabbitMQ"
	"errors"
	"github.com/streadway/amqp"
	"sync"
)

var facade facade2.InterfaceFacade

func init() {
	var once sync.Once
	once.Do(func() {
		facade = facade2.NewFacade()
	})
}

type Producer struct{}

func (Producer) Push(producer queue.ProducerCommonInterface, content string) error {
	provider, err := facade.GetRabbitMQProvider()
	if err != nil {
		return errors.New("201 " + err.Error())
	}
	err = provider.PushByExchange(rabbitMQ.PushParams{
		Content:      content,
		QueueName:    producer.GetProducerQueueName(),
		ExchangeName: producer.GetProducerExchangeName(),
		ExchangeType: amqp.ExchangeDirect,
		RoutingKey:   producer.GetProducerRoutingKey(),
		Durable:      true,
		Arguments:    producer.GetProducerOtherDeclareInfo().Arguments,
	})
	if err != nil {
		return errors.New("202 " + err.Error())
	}
	return nil
}
