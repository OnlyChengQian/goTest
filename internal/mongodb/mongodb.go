package mongodb

import (
	"advt/internal/file"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"sync"
)

type MongoProviderInterface interface {
	LazyMongoDBPool() error
	GetMongoClient() (*mongo.Client, error)
	ReleaseMongoClient(client *mongo.Client)
}

type MongoProvider struct {
	configReader file.ConfigReader
	pool         chan *mongo.Client
	mutex        sync.Mutex
}

func NewMongoProvide(reader file.ConfigReader) MongoProviderInterface {
	return &MongoProvider{
		configReader: reader,
	}
}

func (r *MongoProvider) LazyMongoDBPool() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	config := r.configReader.GetMongoDBConfig()
	maxNum, _ := strconv.Atoi(config.MaxConn)
	r.pool = make(chan *mongo.Client, maxNum)
	clientOptions := options.Client().ApplyURI("mongodb://" + config.Host + ":" + config.Port)
	for i := 0; i < cap(r.pool); i++ {
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return err
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			return err
		}
		r.pool <- client
	}
	return nil
}

func (r *MongoProvider) GetMongoClient() (*mongo.Client, error) {
	if cap(r.pool) == 0 {
		err := r.LazyMongoDBPool()
		if err != nil {
			return nil, err
		}
	}
	client, ok := <-r.pool
	if !ok {
		return nil, errors.New("MongoDB pool is empty")
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, errors.New("mongo ping error: " + err.Error())
	}
	return client, nil
}

func (r *MongoProvider) ReleaseMongoClient(client *mongo.Client) {
	r.pool <- client
}
