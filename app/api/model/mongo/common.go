package mongo

import (
	"advt/app/api/model"
	"advt/internal/facade"
	"advt/internal/mongodb"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type Mongo struct {
}

var mongoProvider mongodb.MongoProviderInterface

func init() {
	var once sync.Once
	once.Do(func() {
		mongoProvider = facade.NewFacade().GetMongoProvider()
	})
}

// First tableServer := &mongo.Mongo{}
// data, err := tableServer.FindOne(new(mongo.Test), bson.M{"site_code": "DE"})
func (m *Mongo) First(table model.MongoModelInterface, where bson.M) (interface{}, error) {
	client, err := m.GetMongoClient()
	defer m.ReleaseClient(client)
	if err != nil {
		return nil, errors.New("获取mongo-client异常：" + err.Error())
	}
	result := table.GetTable()
	db := client.Database(table.DataBase())
	collection := db.Collection(table.Collection())
	err = collection.FindOne(context.Background(), where).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *Mongo) Get(table model.MongoModelInterface, where bson.D) ([]interface{}, error) {
	client, err := m.GetMongoClient()
	defer m.ReleaseClient(client)
	if err != nil {
		return nil, errors.New("获取mongo-client异常：" + err.Error())
	}

	cursor, err := client.Database(table.DataBase()).Collection(table.Collection()).Find(context.Background(), where)
	if err != nil {
		return nil, errors.New("获取数据异常" + err.Error())
	}

	var result []interface{}

	for cursor.Next(context.Background()) {
		data := table.GetTable()
		err = cursor.Decode(&data)
		if err != nil {
			return nil, errors.New("获取数据转化失败" + err.Error())
		}
		result = append(result, data)
	}
	return result, nil
}

func (m *Mongo) InsertOne(table model.MongoModelInterface, data interface{}) error {
	return nil
}

func (m *Mongo) InsertMany(table model.MongoModelInterface, data []interface{}) error {
	return nil
}

func (m *Mongo) GetMongoClient() (*mongo.Client, error) {
	return mongoProvider.GetMongoClient()
}

func (m *Mongo) ReleaseClient(client *mongo.Client) {
	mongoProvider.ReleaseMongoClient(client)
}
