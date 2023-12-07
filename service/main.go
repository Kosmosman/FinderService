package main

import (
	"github.com/Kosmosman/service/orderdb"
	"github.com/Kosmosman/service/stan_subscriber"
	"github.com/Kosmosman/service/types"
	"github.com/Kosmosman/service/wbserver"
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
	server := wbserver.ServerAPI{Cache: &cache}
	server.StartServer()
	wg.Wait()
}
