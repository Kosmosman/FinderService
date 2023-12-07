package main

import (
	"github.com/Kosmosman/service/orderdb"
	"github.com/Kosmosman/service/stan_subscriber"
	"github.com/Kosmosman/service/types"
	"sync"
)

func main() {
	var cache types.Cache
	var db orderdb.OrderDB
	db.Connect()
	db.RestoreCache(&cache)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(cache *types.Cache, db *orderdb.OrderDB, wg *sync.WaitGroup) {
		stan_subscriber.ListenStream(cache, db, wg)
	}(&cache, &db, &wg)
	wg.Wait()
}
