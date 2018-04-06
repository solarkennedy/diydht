run: diydht
	./diydht

diydht: main.go
	go build .

test:
	go test -v

benchmark:
	go test -v -bench=.

deps:
	go get .
