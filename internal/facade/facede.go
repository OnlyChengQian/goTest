package facade

import (
	"advt/internal/cache"
	"advt/internal/database"
	"advt/internal/file"
	"advt/internal/mongodb"
	"advt/internal/rabbitMQ"
	"errors"
	"github.com/gomodule/redigo/redis"
	"sync"
)

type InterfaceFacade interface {
	GetRedisProvider() *redis.Pool
	GetRabbitMQProvider() (rabbitMQ.RabbitConnectionPoolInterface, error)
	GetMySQLProvider() database.MySQLPoolProvider
	GetConfig() file.ConfigReader
	GetMongoProvider() mongodb.MongoProviderInterface
}

var configFacade *ConfigFacade

type ConfigFacade struct {
	ConfigReader file.ConfigReader
}

type Facade struct {
	redisProvider    cache.RedisPoolProvider
	rabbitMQProvider rabbitMQ.RabbitConnectionPoolInterface
	mysqlProvider    database.MySQLPoolProvider
	mongoProvider    mongodb.MongoProviderInterface
}

var mu sync.Once

func init() {
	mu.Do(func() {
		_ = NewConfigFacade("./config/.yaml")
	})
}

func NewFacade() InterfaceFacade {
	return &Facade{}
}

func NewConfigFacade(filePath string) error {
	if configFacade == nil {
		err := file.InitConfig(filePath)
		if err != nil {
			return errors.New("failed to initialize config: " + err.Error())
		}
		configReader := file.GetConfig()
		configFacade = &ConfigFacade{
			ConfigReader: configReader,
		}
	}
	return nil
}

func (f *Facade) GetConfig() file.ConfigReader {
	return configFacade.ConfigReader
}

func (f *Facade) GetRedisProvider() *redis.Pool {

	if f.redisProvider == nil {
		f.redisProvider = cache.NewRedisPoolProvider(configFacade.ConfigReader)
	}
	return f.redisProvider.GetRedisConnection()
}

func (f *Facade) GetMySQLProvider() database.MySQLPoolProvider {
	if f.mysqlProvider == nil {
		f.mysqlProvider = database.NewMySQLPoolProvider(configFacade.ConfigReader)
	}
	return f.mysqlProvider
}

func (f *Facade) GetRabbitMQProvider() (rabbitMQ.RabbitConnectionPoolInterface, error) {
	if f.rabbitMQProvider == nil {
		provider := rabbitMQ.NewRabbitMQ(configFacade.ConfigReader)

		err := provider.LazyInitRabbitMQPool()
		if err != nil {
			return nil, err
		}
		f.rabbitMQProvider = provider
	}
	return f.rabbitMQProvider, nil
}

func (f *Facade) GetMongoProvider() mongodb.MongoProviderInterface {
	if f.mongoProvider == nil {
		provider := mongodb.NewMongoProvide(configFacade.ConfigReader)
		f.mongoProvider = provider
	}

	return f.mongoProvider
}
