package consume

import (
	"advt/app/queue"
)

var queueHandle = [...]queue.ConsumerProviderInterface{
	&ComputerShippingByAccount{},
	&FullUpdatePrice{},
}
