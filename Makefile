env:
	cp ./configs/local/.env_template .env

test:
	go test ./...

run: env
	rm -rf .database && docker compose up --build

swagger:
	swag init --parseInternal -d ./cmd/movies_catalog,./internal

bench:
	go run cmd/benchmark/main.go