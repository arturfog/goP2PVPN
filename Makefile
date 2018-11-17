run: gui
	./$^

all: gui run

gui: gui.go
	go build $^

windows_32: main.go
	GOOS=windows GOARCH=386 go build $^

mac_32: main.go
	GOOS=darwin GOARCH=386 go build $^

clean:
	rm gui

.PHONY: run all clean
