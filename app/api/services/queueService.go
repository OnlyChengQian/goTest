package services

type QueueService struct {
}

func (QueueService) GetQueue() interface{} {
	a, _ := facade.GetRedisProvider().Get().Do("get", "aaa")
	return a
}
