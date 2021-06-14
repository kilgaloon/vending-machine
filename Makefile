test-all:
	go test -race ./handlers/user -v
	go test -race ./handlers/product -v