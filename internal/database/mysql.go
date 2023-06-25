package database

import (
	"advt/internal/file"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type MySQLPoolProvider interface {
	GetDB() (*gorm.DB, error)
}

func NewMySQLPoolProvider(configReader file.ConfigReader) MySQLPoolProvider {
	return &MySQLProvider{
		ConfigReader: configReader,
	}
}

type MySQLProvider struct {
	ConfigReader file.ConfigReader
}

func (d *MySQLProvider) GetDB() (*gorm.DB, error) {
	config := d.ConfigReader.GetDatabaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Port, config.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("db open error" + err.Error())
	}
	isDbConf, err := db.DB()
	if err != nil {
		return nil, errors.New("get mysql db error" + err.Error())
	}

	pingErr := isDbConf.Ping()
	if pingErr != nil {
		return nil, errors.New("mysql ping error" + pingErr.Error())
	}

	maxIdleConn, _ := strconv.Atoi(config.MaxIdleConns)
	isDbConf.SetMaxIdleConns(maxIdleConn)
	maxOpenConn, _ := strconv.Atoi(config.MaxOpenConns)
	isDbConf.SetMaxOpenConns(maxOpenConn)
	return db, nil
}
