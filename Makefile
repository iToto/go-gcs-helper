.PHONY: clean

build:
	go build -i -o bin/go-gcs-helper
linux-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/go-gcs-helper
clean:
	rm -R bin/