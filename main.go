package main

import (
	"SenseHoney/app/api"
	"SenseHoney/lib/settings"
	"SenseHoney/lib/utils/pool"
)

type Service struct {
	api.Service
}

func Start() {
	s := new(Service)
	settings.Init((*settings.Service)(s))
}
func main() {

	wg, poolX := pool.New(6)
	wg.Add(6)
	poolX.Submit(func() {
		go Start()
	})
	wg.Wait()
}
