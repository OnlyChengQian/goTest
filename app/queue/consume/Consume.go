package consume

import (
	"advt/app/queue"
	facade2 "advt/internal/facade"
	"fmt"
	"log"
	"sync"
)

var facade facade2.InterfaceFacade

func init() {
	var once sync.Once
	once.Do(func() {
		facade = facade2.NewFacade()
	})
}

func Run() {
	if len(queueHandle) == 0 {
		return
	}
	provider, err := facade.GetRabbitMQProvider()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := provider.GetConn()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()
	log.Print("==========================================================================")

	for _, qh := range queueHandle {

		_, err = channel.QueueDeclare(qh.GetQueueName(), true, false, false, false, qh.GetOtherDeclareInfo().Arguments)
		if err != nil {
			log.Fatal(err)
		}
		err = channel.ExchangeDeclare(qh.GetExchangeName(), qh.GetOtherDeclareInfo().ExchangeType, true, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}
		err = channel.QueueBind(qh.GetQueueName(), qh.GetRoutingKey(), qh.GetExchangeName(), false, nil)
		if err != nil {
			log.Fatal(err)
		}
		delivery, err := channel.Consume(qh.GetQueueName(), "", false, false, false, false, nil)
		log.Print("-----------------------" + qh.GetQueueName() + " BEGIN-----------------------------")
		go func(isQueueHandle queue.ConsumerProviderInterface) {
			for d := range delivery {
				err = isQueueHandle.Process(string(d.Body))
				if err != nil {
					isQueueHandle.FailsAfter(string(d.Body), err)
					d.Nack(false, false)
				} else {
					d.Ack(true)
				}
			}
		}(qh)
	}
	log.Print("==========================================================================")
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	select {}
}
