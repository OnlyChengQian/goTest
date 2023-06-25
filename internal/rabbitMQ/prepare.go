package rabbitMQ

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

type QueueBuild struct {
	QueueName  string
	Durable    bool //是否持久化
	AutoDelete bool //自动删除
	Exclusive  bool
	NoWait     bool
	Arguments  amqp.Table
}

type ExchangeBuild struct {
	ExchangeName string
	ExchangeType string
	Durable      bool
	AutoDelete   bool
	Internal     bool
	NoWait       bool
	Arguments    amqp.Table
}

func (r *RabbitProvider) PrepareQueue(channel *amqp.Channel, queueInfo *QueueBuild) (queue amqp.Queue, err error) {
	if queueInfo.QueueName == "" {
		return queue, errors.New("queueName error")
	}
	queue, err = channel.QueueDeclare(
		queueInfo.QueueName,
		queueInfo.Durable,
		queueInfo.AutoDelete,
		queueInfo.Exclusive,
		queueInfo.NoWait,
		queueInfo.Arguments,
	)
	return
}

func (r *RabbitProvider) PrepareExchange(channel *amqp.Channel, exchangeInfo *ExchangeBuild) error {
	if exchangeInfo.ExchangeName == "" {
		return errors.New("exchange name empty")
	}
	return channel.ExchangeDeclare(
		exchangeInfo.ExchangeName,
		exchangeInfo.ExchangeType,
		exchangeInfo.Durable,
		exchangeInfo.AutoDelete,
		exchangeInfo.Internal,
		exchangeInfo.NoWait,
		exchangeInfo.Arguments,
	)
}

func (r *RabbitProvider) QueueBindExchange(channel *amqp.Channel, QueueName string, RoutingKey string, ExchangeName string) error {
	return channel.QueueBind(QueueName, RoutingKey, ExchangeName, false, nil)
}

type PushParams struct {
	Content      string
	QueueName    string
	ExchangeName string
	ExchangeType string
	Durable      bool
	AutoDelete   bool
	Internal     bool
	NoWait       bool
	Arguments    amqp.Table
	Exclusive    bool
	RoutingKey   string
}

func (r *RabbitProvider) PushByExchange(params PushParams) error {
	conn, err := r.GetConn()
	defer r.RevertConn(conn)
	if err != nil {
		return errors.New(fmt.Sprintf("获取链接失败：%s", err.Error()))
	}
	channel, err := conn.Channel()
	if err != nil {
		return errors.New(fmt.Sprintf("创建channel错误：%s", err.Error()))
	}
	prepareExchangeErr := r.PrepareExchange(channel, &ExchangeBuild{
		ExchangeName: params.ExchangeName,
		ExchangeType: params.ExchangeType,
		Durable:      params.Durable,
		AutoDelete:   params.AutoDelete,
		Internal:     params.Internal,
		NoWait:       params.NoWait,
		Arguments:    params.Arguments,
	})
	if prepareExchangeErr != nil {
		return errors.New(fmt.Sprintf("绑定exchange错误：%s", err.Error()))
	}

	_, err = r.PrepareQueue(channel, &QueueBuild{
		QueueName:  params.QueueName,
		Durable:    params.Durable,
		AutoDelete: params.AutoDelete,
		Exclusive:  params.Exclusive,
		NoWait:     params.NoWait,
		Arguments:  params.Arguments,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("绑定queue错误：%s", err.Error()))
	}
	if bindErr := r.QueueBindExchange(channel, params.QueueName, params.RoutingKey, params.ExchangeName); bindErr != nil {
		return errors.New(fmt.Sprintf("queue绑定exchange错误：%s", bindErr.Error()))
	}
	pushErr := channel.Publish(params.ExchangeName, params.RoutingKey, true, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(params.Content),
	})
	if pushErr != nil {
		return errors.New(fmt.Sprintf("推送队列异常：%s", pushErr.Error()))
	}
	return nil
}
