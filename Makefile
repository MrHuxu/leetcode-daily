test:
	go test -v ./questions/...

build:
	rm -rf output || mkdir output
	go build -o output/website website/main.go

run:
	./output/website