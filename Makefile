test:
	go test -v ./questions/...

build:
	rm -rf output || mkdir output
	go run cmd/*.go -c=fetch
	go build -o output/website website/main.go

run:
	./output/website