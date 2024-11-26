package main

import (
	"log"
	"synapsis-ecommerce/config/util"
	"synapsis-ecommerce/src"
	"sync"
)

func main() {

	env, err := util.LoadConfig()
	if err != nil {
		log.Fatal("cannot load env", err)
	}

	server := src.InitServer(env)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()
}
