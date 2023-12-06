package main

import (
	"fmt"
	"github.com/Kosmosman/service/stan_subscriber"
	"sync"
)

func main() {
	var cache stan_subscriber.Cache
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(cache *stan_subscriber.Cache, wg *sync.WaitGroup) {
		stan_subscriber.ListenStream(cache, wg)
	}(&cache, &wg)
	wg.Wait()
	fmt.Printf("Was readed %d messages\n", len(cache.Data))
	for _, data := range cache.Data {
		println(string(data))
	}

}
