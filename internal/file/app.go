package file

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
)

type ConfigFacade struct {
	ConfigReader ConfigReader
}

type ConfigReader interface {
	GetDatabaseConfig() DatabaseConf
	GetRedisConfig() RedisConf
	GetRabbitMQConfig() RabbitMQConf
	GetServerConfig() ServerConf
	GetAppKey() string
	GetMongoDBConfig() MongoDBConf
}

type YAMLConfigProvider struct {
	config *conf
}

var config *YAMLConfigProvider

func InitConfig(filePath string) error {

	if config != nil {
		return errors.New("config already initialized")
	}
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	var cfg conf
	err = yaml.Unmarshal(yamlFile, &cfg)

	if err != nil {
		return err
	}
	config = &YAMLConfigProvider{config: &cfg}
	return nil
}

func GetConfig() *YAMLConfigProvider {
	return config
}

func (r *YAMLConfigProvider) GetDatabaseConfig() DatabaseConf {
	return r.config.DatabaseLogs
}

func (r *YAMLConfigProvider) GetRedisConfig() RedisConf {
	return r.config.Redis
}

func (r *YAMLConfigProvider) GetRabbitMQConfig() RabbitMQConf {
	return r.config.RabbitMQ
}

func (r *YAMLConfigProvider) GetServerConfig() ServerConf {
	return r.config.Server
}

func (r *YAMLConfigProvider) GetAppKey() string {
	return r.config.Server.AppKey
}

func (r *YAMLConfigProvider) GetMongoDBConfig() MongoDBConf {
	return r.config.MongoDB
}
