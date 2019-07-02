build:
	mkdir -p functions
	go get ./...
	go build -o functions/logs ./logs.go

