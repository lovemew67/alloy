default:
	go mod tidy
	- rm -f alloy.bin
	go build -o=alloy.bin main.go
	./alloy.bin -v
