package queue

import "advt/internal/rabbitMQ"

type ConsumerProviderInterface interface {
	GetRoutingKey() string                        //路由键值
	GetQueueName() string                         //队列名称
	GetExchangeName() string                      //交换机名称
	GetOtherDeclareInfo() *rabbitMQ.ExchangeBuild //其他配置
	Process(data string) error                    //业务处理
	FailsAfter(data string, err error)            //出现错误后的处理工作
}

type ProducerCommonInterface interface {
	GetProducerRoutingKey() string                        //路由键值
	GetProducerQueueName() string                         //队列名称
	GetProducerExchangeName() string                      //交换机名称
	GetProducerOtherDeclareInfo() *rabbitMQ.ExchangeBuild //其他配置
}

// queuename
const (
	ComputeShippingQueueName = "compute_shipping"
	TestQueueName            = "test.number.one"
)

// routingkey
const (
	ComputeShippingRoutingKey = "compute.shipping"
	TestQueueRoutingKey       = "test.number.one"
)

// exchange
const (
	DefaultExchangeName = "compute_shipping_exchange"
)
