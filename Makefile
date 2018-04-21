run: main.go
	go run $^

all: main.go
	go build $^

windows_32: main.go
    GOOS=windows GOARCH=386 go build $^

mac_32: main.go
    GOOS=darwin GOARCH=386 go build $^

.PHONY: run all
