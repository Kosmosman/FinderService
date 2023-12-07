package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func main() {
	var dirPath string = "./publisher/orders/"
	connect, err := stan.Connect("test-cluster", "publisher")
	if err != nil {
		log.Fatal(err)
	}
	defer connect.Close()

	orders, _ := os.ReadDir(dirPath)
	for _, fp := range orders {
		data, err := os.ReadFile(dirPath + fp.Name())
		if err != nil {
			log.Fatal(err)
		}
		if err = connect.Publish("order", data); err != nil {
			log.Fatal(err)
		}
	}
	println("Messages delivered...")
}
