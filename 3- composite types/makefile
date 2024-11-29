.DEFAULT_GOAL := run

.PHONY: fmt vet build run clean

fmt:
	go fmt ./main.go
vet:
	go vet ./main.go
clean:
	rm -f composite-types
build: clean
	go build -o composite-types ./main.go
run: build
	./composite-types