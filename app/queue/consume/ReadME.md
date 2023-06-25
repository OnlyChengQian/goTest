# 介绍

这个文件夹下面是消费队列

需要增加自定义方法来完成数据的操作


方法在ConsumerHandle里面的ListenFunction()

增加一个监听，需要ListenFunction里添加"queue" : func(data string) error
> ps 因为队列的数据是string，所以接受到的data值也是string @-@！！！

代码如下：
```
func (l *ListenConsumers) ListenFunction() map[string]func(data string) error {
	var isFunc = map[string]func(data string) error{
		"queue-a": func(data string) error {
            ...
            ...
			return nil
		},
		"queue-b" : func(data string) error {
            ...
            ...
            return nil
        }   

	}
	return isFunc
}
```
