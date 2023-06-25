package command

import (
	"advt/app/queue/produce"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type ComputeShippingScript struct {
}

func (c *ComputeShippingScript) handle() error {
	var producer = &produce.Producer{}
	rand.Seed(time.Now().UnixNano())
	number := strconv.Itoa(rand.Intn(999999))
	producer.Push(new(produce.ComputeShippingProducer), number)
	fmt.Println(number)
	return nil
}
