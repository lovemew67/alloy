default:
	go mod tidy
	- rm -f alloy.bin
	go build -o=alloy.bin main.go
	./alloy.bin -v

test:
	go test -v ./...

# https://stackoverflow.com/questions/66921582/bash-find-files-with-certain-extension-but-exclude-those-with-certain-keyword-i
# find /Users/Tim/Downloads -type f -name "*.dpx" -o '(' -name "*.exr" -a '!' -name "*BTY*.exr" ')'
# find . -type f '!' -name "*.go"
