package stan_publisher

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetStanConnection() stan.Conn {
	connect, err := stan.Connect("test-cluster", "publisher")
	if err != nil {
		log.Fatal(err)
	}
	return connect
}

func PublishMessanges(conn *stan.Conn) {
	_, filename, _, _ := runtime.Caller(0)
	dirPath := filepath.Dir(filename) + "/../orders/"
	orders, _ := os.ReadDir(dirPath)
	for _, fp := range orders {
		data, err := os.ReadFile(dirPath + fp.Name())
		if err != nil {
			log.Fatal(err)
		}
		if err = (*conn).Publish("order", data); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Messages delivered...")
}
