
####### Main #######

server:
	go run cmd/server/main.go

producer:
	go run cmd/server/producer/producer.go

test:
	go test -v ./internal/usecase/user_test.go

####### Docker compose #######

docker-up:
	docker compose -f docker-compose.dev.yaml up -d

docker-down:
	docker compose -f docker-compose.dev.yaml down

####### Migrate #######

migrate-up:
	 migrate -path migrations -database "postgres://root:postgres@localhost:5435/fio?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://root:postgres@localhost:5435/fio?sslmode=disable" down



