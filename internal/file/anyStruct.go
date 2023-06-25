package file

type conf struct {
	DatabaseLogs DatabaseConf `yaml:"database"`
	Redis        RedisConf    `yaml:"redis"`
	RabbitMQ     RabbitMQConf `yaml:"rabbitMQ"`
	Server       ServerConf   `yaml:"server"`
	MongoDB      MongoDBConf  `yaml:"mongoDB"`
}

type DatabaseConf struct {
	Name         string `yaml:"Name"`
	Host         string `yaml:"Host"`
	Port         string `yaml:"Port"`
	User         string `yaml:"User"`
	Pass         string `yaml:"Pass"`
	MaxOpenConns string `yaml:"MaxOpenConns"`
	MaxIdleConns string `yaml:"MaxIdleConns"`
	MaxIdleTime  string `yaml:"MaxIdleTime"`
	MaxLifetime  string `yaml:"MaxLifetime"`
}

type RedisConf struct {
	Host        string `yaml:"Host"`
	User        string `yaml:"User"`
	Port        string `yaml:"Port"`
	Auth        string `yaml:"Auth"`
	MaxIdle     string `yaml:"MaxIdle"`
	MaxActive   string `yaml:"MaxActive"`
	IdleTimeout string `yaml:"IdleTimeout"`
}

type RabbitMQConf struct {
	Host        string `yaml:"Host"`
	Port        string `yaml:"Port"`
	User        string `yaml:"User"`
	Pass        string `yaml:"Pass"`
	MinConnCap  string `yaml:"MinConnCap"`
	MaxLifeCap  string `yaml:"MaxLifeCap"`
	MaxIdle     string `yaml:"MaxIdle"`
	IdleTimeout string `yaml:"IdleTimeout"`
	Vhost       string `yaml:"Vhost"`
}

type ServerConf struct {
	Port    string `yaml:"Port"`
	AppKey  string `yaml:"AppKey"`
	Debug   string `yaml:"Debug"`
	RpcPort string `yaml:"RpcPort"`
}

type MongoDBConf struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Auth     string `yaml:"Auth"`
	Database string `yaml:"Database"`
	MaxConn  string `yaml:"MaxConn"`
}
