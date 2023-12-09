services:
	docker-compose up &

server:
	go run service/main.go

stop:
	docker-compose down

test:
	go test tests/service_test.go

script:
	go run publisher/main.go

interface:
	firefox -new-tab http://localhost:8080/orders

.PHONY: interface