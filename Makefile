
####### Main #######

server:
	go run cmd/server/main.go


####### Docker compose #######

dev:
	docker compose -f docker-compose.dev.yaml up -d
