.DEFAULT_GOAL := run

.PHONY: fmt vet build run clean

fmt:
	go fmt ./main.go
vet:
	go vet ./main.go
clean:
	rm -f tooling
build: clean
	go build -o tooling ./main.go
run: build
	./tooling