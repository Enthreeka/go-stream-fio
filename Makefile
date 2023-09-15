
####### Main #######

server:
	go run cmd/server/main.go


####### Docker compose #######

docker-up:
	docker compose -f docker-compose.dev.yaml up -d

docker-down:
	docker compose -f docker-compose.dev.yaml down

