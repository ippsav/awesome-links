run:
	go run cmd/api/main.go

generate:
	go generate ./ent
	gqlgen
	
migrate:
	go run cmd/migration/main.go