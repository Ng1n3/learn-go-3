.DEFAULT_GOAL := run

.PHONY: fmt vet build run clean

fmt:
	go fmt ./main.go
vet:
	go vet ./main.go
clean:
	rm -f errors
build: clean
	go build -o errors ./main.go
run: build
	./errors