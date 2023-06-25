package mysql

import (
	"advt/app/api/model"
	"advt/internal/database"
	"advt/internal/facade"
	"errors"
	"gorm.io/gorm"
	"sync"
)

type Mysql struct {
}

var databaseProvider database.MySQLPoolProvider

func init() {
	var once sync.Once
	once.Do(func() {
		if databaseProvider == nil {
			databaseProvider = facade.NewFacade().GetMySQLProvider()
		}
	})
}

func (m *Mysql) Get(table model.MysqlModelInterface, where interface{}) ([]interface{}, error) {
	var content []interface{}
	content = append(content, table.GetTable())
	db, err := m.GetMysqlClient()
	if err != nil {
		return nil, err
	}
	db.Table(table.GetTableName()).Where(where).Find(&content)
	return content, nil

}

func (m *Mysql) First(table model.MysqlModelInterface, where interface{}) (interface{}, error) {
	return nil, nil
}

func (m *Mysql) GetMysqlClient() (*gorm.DB, error) {
	client, err := databaseProvider.GetDB()
	if err != nil {
		return nil, errors.New("获取数据库client异常" + err.Error())
	}
	return client, nil
}
