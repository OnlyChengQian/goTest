package exception

import errors "github.com/pkg/errors"

type Exception struct {
	message string `json:"message"`
	code    int    `json:"code"`
}

func (e Exception) Error() string {
	return e.message
}

const (
	ServerError = 20001

	RedisConnError   = 21001
	RedisConfigError = 21002

	MysqlConnError   = 22001
	MysqlConfigError = 22002

	RabbitMQConnError   = 23001
	RabbitMQConfigError = 23002
)

var codeMapTag = map[int]string{

	ServerError: "Internal Server Error",

	RedisConnError:   "redis链接异常",
	RedisConfigError: "redis配置异常",

	MysqlConnError:   "数据库链接异常",
	MysqlConfigError: "数据库配置异常",

	RabbitMQConnError:   "rabbit链接异常",
	RabbitMQConfigError: "rabbit配置异常",
}

func mapMessage(code int) string {
	return codeMapTag[code]
}

func (Exception) MakeError(code int) error {
	return errors.Wrap(&Exception{code: code, message: codeMapTag[code]}, "")
}
