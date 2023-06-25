package controller

import (
	facade2 "advt/internal/facade"
	"sync"
)

var facade facade2.InterfaceFacade

var facadeOnce sync.Once

func init() {
	facadeOnce.Do(func() {
		facade = facade2.NewFacade()
	})
}
