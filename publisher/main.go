package main

import (
	"github.com/Kosmosman/publisher/stan_publisher"
)

func main() {
	connect := stan_publisher.GetStanConnection()
	defer connect.Close()
	stan_publisher.PublishMessanges(&connect)
}
