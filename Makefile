run: diydht
	./diydht

diydht: main.go
	go build .

test:
	go test -v

deps:
	go get .
