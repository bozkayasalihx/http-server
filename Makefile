migrate-up: 
	migrate -verbose -path ./db/migration -database postgresql://maker:maker@localhost/maker?sslmode=disable up

migrate-down:
	migrate -verbose -path ./db/migration -database postgresql://maker:maker@localhost/maker?sslmode=disable down 

generate:
	sqlc generate

test: 
	go test ./... -v -cover 

.PHONY: migrate-up migrate-down generate test