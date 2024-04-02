dev:
	~/go/bin/air --build.cmd "go build -o ./bin/mrt-go ./cmd/main.go" --build.bin ./bin/mrt-go
api-docs:
	~/go/bin/swag init -g ./cmd/main.go