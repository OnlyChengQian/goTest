package model

type MysqlModelInterface interface {
	GetTableName() string
	GetConnection() string
	GetTable() interface{}
}

type MongoModelInterface interface {
	DataBase() string
	Collection() string
	GetTable() interface{}
}
