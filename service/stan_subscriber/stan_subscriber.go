package stan_subscriber

import (
	"encoding/json"
	"github.com/Kosmosman/service/orderdb"
	"github.com/Kosmosman/service/types"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func add(orderJson []byte, ch *types.Cache, db *orderdb.OrderDB) {
	v := validator.New(validator.WithRequiredStructEnabled())

	var order types.Order
	if err := json.Unmarshal(orderJson, &order); err != nil {
		log.Fatal(err)
	}
	if err := v.Struct(order); err == nil {
		ch.Mutex.Lock()
		defer ch.Mutex.Unlock()
		if ch.Data == nil {
			ch.Data = make(map[string][]byte)
		}
		if _, ok := ch.Data[order.OrderUID]; !ok {
			orderStringView := string(orderJson)
			ch.Data[order.OrderUID] = orderJson
			db.Add(&order.OrderUID, &orderStringView)
		}
	} else {
		println(err.Error())
	}
}

func ListenStream(cache *types.Cache, db *orderdb.OrderDB, wg *sync.WaitGroup) {
	defer wg.Done()

	sc, err := stan.Connect("test-cluster", "subscriber")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("order", func(msg *stan.Msg) {
		add(msg.Data, cache, db)
	}, stan.StartWithLastReceived())
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	defer sub.Close()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
}
