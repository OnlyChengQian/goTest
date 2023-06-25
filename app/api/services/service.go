package services

import (
	facade2 "advt/internal/facade"
	"sync"
)

var facade facade2.InterfaceFacade

func init() {
	var once sync.Once
	once.Do(func() {
		facade = facade2.NewFacade()
	})

}
