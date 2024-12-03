.DEFAULT_GOAL := run

.PHONY: fmt vet build run clean

fmt:
	go fmt ./main.go
vet:
	go vet ./main.go
clean:
	rm -f interfaces
build: clean
	go build -o interfaces ./main.go
run: build
	./interfaces