run: main.go
	go run $^
all: main.go
	go build $^

.PHONY: run all
