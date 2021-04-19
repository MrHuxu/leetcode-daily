gen:
	go run cmd/*.go -c=gen

test:
	go test -v ./questions/...

web:
	rm -rf output || mkdir output
	go build -o output/website website/main.go
	go run cmd/*.go -c=fetch

serve:
	./output/website