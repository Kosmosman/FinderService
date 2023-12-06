package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func main() {
	var dirParh string = "./publisher/orders/"
	connect, err := stan.Connect("test-cluster", "publisher")
	if err != nil {
		log.Fatal(err)
	}
	defer connect.Close()

	orders, _ := os.ReadDir(dirParh)
	for _, fp := range orders {
		data, err := os.ReadFile(dirParh + fp.Name())
		if err != nil {
			log.Fatal(err)
		}
		if err = connect.Publish("order", data); err != nil {
			log.Fatal(err)
		}
	}
	println("Messages delivered...")
}
