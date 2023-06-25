package produce

import (
	"advt/app/queue"
	"advt/internal/rabbitMQ"
	"github.com/streadway/amqp"
)

type ComputeShippingProducer struct{}

func (c *ComputeShippingProducer) GetProducerRoutingKey() string {
	return queue.ComputeShippingRoutingKey
}

func (c *ComputeShippingProducer) GetProducerQueueName() string {
	return queue.ComputeShippingQueueName
}

func (c *ComputeShippingProducer) GetProducerExchangeName() string {
	return queue.DefaultExchangeName
}

func (c *ComputeShippingProducer) GetProducerOtherDeclareInfo() *rabbitMQ.ExchangeBuild {

	return &rabbitMQ.ExchangeBuild{
		ExchangeType: amqp.ExchangeDirect,
		Durable:      true,
		Arguments: map[string]interface{}{
			"x-max-priority": 20,
		},
	}

}
