package rabbitMQ

import (
	"advt/internal/file"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
	"time"
)

type RabbitConnectionPoolInterface interface {
	GetConn() (*amqp.Connection, error)
	RevertConn(conn *amqp.Connection)
	PushByExchange(params PushParams) error
	LazyInitRabbitMQPool() error
	QueueBindExchange(channel *amqp.Channel, LogsQueueName string, LogsRoutingKey string, LogsExchange string) error
	PrepareExchange(channel *amqp.Channel, exchangeInfo *ExchangeBuild) error
	PrepareQueue(channel *amqp.Channel, queueInfo *QueueBuild) (queue amqp.Queue, err error)
}

type RabbitProvider struct {
	connections    chan *amqp.Connection
	ConfigProvider file.ConfigReader
	HostUrl        string
}

func (r *RabbitProvider) LazyInitRabbitMQPool() error {
	if r.connections == nil {
		poolNums, _ := strconv.Atoi(r.ConfigProvider.GetRabbitMQConfig().MaxLifeCap)
		r.connections = make(chan *amqp.Connection, poolNums) //需要先初始化
		for i := 0; i < poolNums; i++ {
			conn, err := amqp.Dial(r.HostUrl)
			if err != nil {
				return err
			}
			r.connections <- conn
		}
	}
	return nil
}

func NewRabbitMQ(configProvider file.ConfigReader) RabbitConnectionPoolInterface {
	return &RabbitProvider{
		ConfigProvider: configProvider,
		HostUrl:        mqHost(configProvider),
	}
}

func (r *RabbitProvider) GetConn() (*amqp.Connection, error) {
	select {
	case conn := <-r.connections:
		if conn.IsClosed() {
			conn.Close()
			return nil, errors.New("RabbitMQ connection is closed")
		}
		return conn, nil
	case <-time.After(time.Second * 5):
		return nil, errors.New("timeout: No available RabbitMQ connection")
	default:
		return nil, errors.New("RabbitMQ connection pool is empty or too many connection")
	}

}

func (r *RabbitProvider) RevertConn(conn *amqp.Connection) {
	select {
	case r.connections <- conn:
	default:
		conn.Close()
	}
}

func mqHost(configProvider file.ConfigReader) string {
	rabbitMQConf := configProvider.GetRabbitMQConfig()
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s", rabbitMQConf.User, rabbitMQConf.Pass, rabbitMQConf.Host, rabbitMQConf.Port, rabbitMQConf.Vhost)
}
