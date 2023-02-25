air-cmd:
	go mod tidy
	go build -o ./tmp/main ./main.go
