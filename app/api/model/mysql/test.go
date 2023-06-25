package mysql

type Test struct {
}

type TestTable struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

func (Test) GetTableName() string {
	return "test"
}

func (Test) GetConnection() string {
	return "istore2"
}

func (Test) GetTable() interface{} {
	return TestTable{}
}
