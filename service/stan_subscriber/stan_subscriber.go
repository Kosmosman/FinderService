package stan_subscriber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Cache struct {
	Mutex sync.RWMutex
	Data  map[string][]byte
}

func (ch *Cache) Add(orderJson []byte) {
	orderData := json.NewDecoder(bytes.NewReader(orderJson))
	type orderId struct {
		Id string `json:"order_uid"`
	}
	var Id orderId
	err := orderData.Decode(&Id)
	if err != nil {
		log.Fatal(err)
	}
	ch.Mutex.Lock()
	defer ch.Mutex.Unlock()
	if ch.Data == nil {
		ch.Data = make(map[string][]byte)
	}
	ch.Data[Id.Id] = orderJson
}

func ListenStream(cache *Cache, wg *sync.WaitGroup) {
	defer wg.Done()

	sc, err := stan.Connect("test-cluster", "subscriber")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("order", func(msg *stan.Msg) {
		cache.Add(msg.Data)
		fmt.Println("Add new message")
	}, stan.StartWithLastReceived())
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	defer sub.Close()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Ожидаем сигнала прерывания
	<-signalCh

	//select {}
}
