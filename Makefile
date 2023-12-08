start:
	docker-compose up &
	sleep 5
	go run service/main.go

stop:
	docker-compose down

test:
	go test tests/service_test.go

script:
	go run publisher/main.go