run:
	go build -o bin/app ./cmd/app
	./bin/app
migrate-up:
	go run internal/migrations/cmd/main.go up
migrate-down:
	go run internal/migrations/cmd/main.go down